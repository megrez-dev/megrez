package admin

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/megrez/pkg/model"
	"github.com/megrez/pkg/utils/errmsg"
)

func CreateCategory(c *gin.Context) {
	var data model.Category
	err := c.ShouldBindJSON(&data)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusOK, errmsg.Fail(errmsg.ERROR_INVALID_PARAM))
		return
	}

	category, err := model.GetCategoryBySlug(data.Slug)
	if category.ID != 0 {
		c.JSON(http.StatusOK, errmsg.Fail(errmsg.ERROR_CATENAME_SLUG_EXIST))
		return
	}

	err = model.CreateCategory(&data)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}

	c.JSON(http.StatusOK, errmsg.Success(data))
}

func UpdateCategory(c *gin.Context) {
	var data model.Category
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusOK, errmsg.Error())
	}
	err = c.ShouldBindJSON(&data)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusOK, errmsg.Fail(errmsg.ERROR_INVALID_PARAM))
	}

	err = model.UpdateCategoryByID(uint(id), &data)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusOK, errmsg.Error())
	}

	c.JSON(http.StatusOK, errmsg.Success(data))
}

func DeleteCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusOK, errmsg.Error())
	}

	err = model.DeleteCategoryByID(uint(id))
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusOK, errmsg.Error())
	}

	c.JSON(http.StatusOK, errmsg.Success(nil))
}

func GetCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusOK, errmsg.Error())
	}
	category, err := model.GetCategoryByID(uint(id))
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusOK, errmsg.Error())
	}

	c.JSON(http.StatusOK, errmsg.Success(category))
}

func ListCategories(c *gin.Context) {
	pageNumStr := c.Query("pageNum")
	pageSizeStr := c.Query("pageSize")
	if pageNumStr == "" && pageSizeStr == "" {
		categories, err := model.ListAllCategories()
		if err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusOK, errmsg.Error())
			return
		}
		c.JSON(http.StatusOK, errmsg.Success(categories))
		return
	} else {
		pageNum, err := strconv.Atoi(c.Query("pageNum"))
		pageSize, err := strconv.Atoi(c.Query("pageSize"))
		if err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusOK, errmsg.Fail(errmsg.ERROR_INVALID_PARAM))
			return
		}
		categories, err := model.ListCategoriesByPage(pageNum, pageSize)
		if err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusOK, errmsg.Error())
			return
		}

		c.JSON(http.StatusOK, errmsg.Success(categories))
	}
}
