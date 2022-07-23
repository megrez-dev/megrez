package admin

import (
	"github.com/megrez/pkg/entity/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/megrez/pkg/log"
	"github.com/megrez/pkg/model"
	"github.com/megrez/pkg/utils/errmsg"
)

// CreateCategory godoc
// @Summary create category
// @Schemes http https
// @Description create category
// @Accept application/json
// @Param Authorization header string false "Authorization"
// @Param req body model.Category true "body"
// @Success 200 {object} errmsg.Response{data=model.Category}
// @Router /api/admin/category [post]
func CreateCategory(c *gin.Context) {
	var data model.Category
	err := c.ShouldBindJSON(&data)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusOK, errmsg.Fail(errmsg.ErrorInvalidParam))
		return
	}

	category, err := model.GetCategoryBySlug(data.Slug)
	if category.ID != 0 {
		c.JSON(http.StatusOK, errmsg.Fail(errmsg.ErrorCategorySlugExist))
		return
	}

	tx := model.BeginTx()
	err = model.CreateCategory(tx, &data)
	if err != nil {
		log.Error(err.Error())
		tx.Rollback()
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}

	tx.Commit()
	c.JSON(http.StatusOK, errmsg.Success(data))
}

// UpdateCategory godoc
// @Summary update category
// @Schemes http https
// @Description update category
// @Accept application/json
// @Param Authorization header string false "Authorization"
// @Param id path int true "category id"
// @Param req body model.Category true "category"
// @Success 200 {object} errmsg.Response{data=model.Category}
// @Router /api/admin/category/{id} [put]
func UpdateCategory(c *gin.Context) {
	var data model.Category
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	err = c.ShouldBindJSON(&data)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusOK, errmsg.Fail(errmsg.ErrorInvalidParam))
		return
	}

	tx := model.BeginTx()
	err = model.UpdateCategoryByID(tx, uint(id), &data)
	if err != nil {
		log.Error(err.Error())
		tx.Rollback()
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}

	tx.Commit()
	c.JSON(http.StatusOK, errmsg.Success(data))
}

// DeleteCategory godoc
// @Summary delete category by category id
// @Schemes http https
// @Description delete category by category id
// @Accept application/json
// @Param Authorization header string false "Authorization"
// @Param id path int true "category id"
// @Success 200 {object} errmsg.Response{}
// @Router /api/admin/category/{id} [delete]
func DeleteCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}

	tx := model.BeginTx()
	err = model.DeleteCategoryByID(tx, uint(id))
	if err != nil {
		log.Error(err.Error())
		tx.Rollback()
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}

	tx.Commit()
	c.JSON(http.StatusOK, errmsg.Success(nil))
}

// GetCategory godoc
// @Summary get category by category id
// @Schemes http https
// @Description get category by category id
// @Accept application/json
// @Param Authorization header string false "Authorization"
// @Param id path int false "category id"
// @Success 200 {object} errmsg.Response{data=model.Category}
// @Router /api/admin/category/{id} [get]
func GetCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	category, err := model.GetCategoryByID(uint(id))
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}

	c.JSON(http.StatusOK, errmsg.Success(category))
}

// ListCategories godoc
// @Summary list categories
// @Schemes http https
// @Description list categories
// @Accept application/json
// @Param Authorization header string false "Authorization"
// @Param pageNum query int false "page num"
// @Param pageSize query int false "page size"
// @Success 200 {object} errmsg.Response{data=dto.Pagination{list=[]model.Category}}
// @Router /api/admin/categories [get]
func ListCategories(c *gin.Context) {
	pageNumStr := c.Query("pageNum")
	pageSizeStr := c.Query("pageSize")
	if pageNumStr == "" && pageSizeStr == "" {
		categories, err := model.ListAllCategories()
		if err != nil {
			log.Error(err.Error())
			c.JSON(http.StatusOK, errmsg.Error())
			return
		}
		c.JSON(http.StatusOK, errmsg.Success(categories))
		return
	} else {
		pageNum, err := strconv.Atoi(c.Query("pageNum"))
		pageSize, err := strconv.Atoi(c.Query("pageSize"))
		if err != nil {
			log.Error(err.Error())
			c.JSON(http.StatusOK, errmsg.Fail(errmsg.ErrorInvalidParam))
			return
		}
		categories, err := model.ListCategoriesByPage(pageNum, pageSize)
		if err != nil {
			log.Error(err.Error())
			c.JSON(http.StatusOK, errmsg.Error())
			return
		}
		total, err := model.CountCategories()
		if err != nil {
			log.Error(err.Error())
			c.JSON(http.StatusOK, errmsg.Error())
			return
		}
		pagination := dto.Pagination{
			Current:  pageNum,
			PageSize: pageSize,
			Total:    total,
			List:     categories,
		}
		c.JSON(http.StatusOK, errmsg.Success(pagination))
	}
}
