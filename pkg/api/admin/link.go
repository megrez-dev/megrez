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

// CreateLink godoc
// @Summary create link
// @Schemes http https
// @Description create link
// @Accept application/json
// @Param Authorization header string false "Authorization"
// @Param req body model.Link true "body"
// @Success 200 {object} errmsg.Response{data=model.Link}
// @Router /api/admin/link [post]
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

// UpdateLink godoc
// @Summary update link
// @Schemes http https
// @Description update link
// @Accept application/json
// @Param Authorization header string false "Authorization"
// @Param id path int true "link id"
// @Param req body model.Link true "link"
// @Success 200 {object} errmsg.Response{data=model.Link}
// @Router /api/admin/link/{id} [put]
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

// DeleteLink godoc
// @Summary delete link by link id
// @Schemes http https
// @Description delete link by link id
// @Accept application/json
// @Param Authorization header string false "Authorization"
// @Param id path int true "link id"
// @Success 200 {object} errmsg.Response{}
// @Router /api/admin/link/{id} [delete]
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

// ListLinks godoc
// @Summary list links
// @Schemes http https
// @Description list links
// @Accept application/json
// @Param Authorization header string false "Authorization"
// @Param pageNum query int false "page num"
// @Param pageSize query int false "page size"
// @Success 200 {object} errmsg.Response{data=dto.Pagination{list=[]model.Link}}
// @Router /api/admin/links [get]
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
