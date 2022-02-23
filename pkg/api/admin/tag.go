package admin

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/megrez/pkg/model"
	"github.com/megrez/pkg/utils/errmsg"
)

func CreateTag(c *gin.Context) {
	var data model.Tag
	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusOK, errmsg.Fail(errmsg.ErrorInvalidParam))
		return
	}

	tag, err := model.GetTagByName(data.Name)
	if tag.ID != 0 {
		c.JSON(http.StatusOK, errmsg.Success(tag))
		return
	}

	err = model.CreateTag(&data)
	if err != nil {
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	c.JSON(http.StatusOK, errmsg.Success(data))
}

func ListTags(c *gin.Context) {
	pageNumStr := c.Query("pageNum")
	pageSizeStr := c.Query("pageSize")
	if pageNumStr == "" && pageSizeStr == "" {
		tags, err := model.ListAllTags()
		if err != nil {
			c.JSON(http.StatusOK, errmsg.Error())
			return
		}
		c.JSON(http.StatusOK, errmsg.Success(tags))
		return
	} else {
		pageNum, err := strconv.Atoi(c.Query("pageNum"))
		pageSize, err := strconv.Atoi(c.Query("pageSize"))
		if err != nil {
			c.JSON(http.StatusOK, errmsg.Fail(errmsg.ErrorInvalidParam))
			return
		}
		tags, err := model.ListTagsByPage(pageNum, pageSize)
		if err != nil {
			c.JSON(http.StatusOK, errmsg.Error())
			return
		}

		c.JSON(http.StatusOK, errmsg.Success(tags))
	}
}
