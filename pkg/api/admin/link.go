package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/megrez/pkg/entity/dto"
	"github.com/megrez/pkg/model"
	"github.com/megrez/pkg/utils/errmsg"
	"log"
	"net/http"
	"strconv"
)

func CreateLink(c *gin.Context) {
	var data model.Link
	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusOK, errmsg.Fail(errmsg.ErrorInvalidParam))
		return
	}

	err = model.CreateLink(&data)
	if err != nil {
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	c.JSON(http.StatusOK, errmsg.Success(data))
}

func UpdateLink(c *gin.Context) {
	var data model.Link
	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusOK, errmsg.Fail(errmsg.ErrorInvalidParam))
		return
	}

	err = model.UpdateLink(&data)
	if err != nil {
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	c.JSON(http.StatusOK, errmsg.Success(data))
}

func ListLinks(c *gin.Context) {
	pageNumStr := c.Query("pageNum")
	pageSizeStr := c.Query("pageSize")
	if pageNumStr == "" {
		pageNumStr = "1"
	}
	if pageSizeStr == "" {
		pageSizeStr = "10"
	}
	pageNum, err := strconv.Atoi(pageNumStr)
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		c.JSON(http.StatusOK, errmsg.Fail(errmsg.ErrorInvalidParam))
		return
	}
	links, err := model.ListLinksByPage(pageNum, pageSize)
	count, err := model.CountLinks()
	if err != nil {
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	if err != nil {
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	p := dto.Pagination{
		Current:  pageNum,
		PageSize: pageSize,
		Total:    count,
		List:     links,
	}
	c.JSON(http.StatusOK, errmsg.Success(p))
}

func DeleteLink(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	err = model.DeleteLinkByID(uint(id))
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	c.JSON(http.StatusOK, errmsg.Success(nil))
}
