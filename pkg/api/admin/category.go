package admin

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/megrez/pkg/log"
	"github.com/megrez/pkg/model"
	"github.com/megrez/pkg/utils/errmsg"
)

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

		c.JSON(http.StatusOK, errmsg.Success(categories))
	}
}
