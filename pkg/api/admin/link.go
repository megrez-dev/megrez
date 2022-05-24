package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/megrez/pkg/entity/dto"
	"github.com/megrez/pkg/log"
	"github.com/megrez/pkg/model"
	"github.com/megrez/pkg/utils/errmsg"
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

	tx := model.BeginTx()
	err = model.CreateLink(tx, &data)
	if err != nil {
		log.Error("create link error: %v", err)
		tx.Rollback()
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	tx.Commit()
	c.JSON(http.StatusOK, errmsg.Success(data))
}

func UpdateLink(c *gin.Context) {
	var data model.Link
	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusOK, errmsg.Fail(errmsg.ErrorInvalidParam))
		return
	}

	tx := model.BeginTx()
	err = model.UpdateLink(tx, &data)
	if err != nil {
		log.Error("update link error: %v", err)
		tx.Rollback()
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	tx.Commit()
	c.JSON(http.StatusOK, errmsg.Success(data))
}

func DeleteLink(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	tx := model.BeginTx()
	err = model.DeleteLinkByID(tx, uint(id))
	if err != nil {
		log.Error(err.Error())
		tx.Rollback()
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	tx.Commit()
	c.JSON(http.StatusOK, errmsg.Success(nil))
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
