package admin

import (
	"github.com/megrez/pkg/log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/megrez/pkg/model"
	"github.com/megrez/pkg/utils/errmsg"
)

// CreateTag godoc
// @Summary create tag
// @Schemes http https
// @Description create tag
// @Accept application/json
// @Param Authorization header string false "Authorization"
// @Param req body model.Tag true "body"
// @Success 200 {object} errmsg.Response{data=model.Tag}
// @Router /api/admin/tag [post]
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
	tx := model.BeginTx()
	err = model.CreateTag(tx, &data)
	if err != nil {
		log.Error("create tag error: %v", err)
		tx.Rollback()
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	tx.Commit()
	c.JSON(http.StatusOK, errmsg.Success(data))
}

// ListTags godoc
// @Summary list tags
// @Schemes http https
// @Description list tags
// @Accept application/json
// @Param Authorization header string false "Authorization"
// @Param pageNum query int false "page num"
// @Param pageSize query int false "page size"
// @Success 200 {object} errmsg.Response{data=dto.Pagination{list=[]model.Tag}}
// @Router /api/admin/tags [get]
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
