package admin

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/megrez/pkg/model"
	"github.com/megrez/pkg/utils/errmsg"
)

func CreateArticle(c *gin.Context) {
	var data model.Article
	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{})
	}

	err = model.CreateArticle(&data)
	if err != nil {
		c.JSON(http.StatusOK, errmsg.Error())
	}

	c.JSON(http.StatusOK, errmsg.Success(data))
}

func UpdateArticle(c *gin.Context) {
	var data model.Article
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, errmsg.Error())
	}
	err = c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{})
	}

	err = model.UpdateArticleByID(uint(id), &data)
	if err != nil {
		c.JSON(http.StatusOK, errmsg.Error())
	}

	c.JSON(http.StatusOK, errmsg.Success(nil))
}

func DeleteArticle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, errmsg.Error())
	}

	err = model.DeleteArticleByID(uint(id))
	if err != nil {
		c.JSON(http.StatusOK, errmsg.Error())
	}

	c.JSON(http.StatusOK, errmsg.Success(nil))
}

func GetArticle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, errmsg.Error())
	}
	article, err := model.GetArticleByID(uint(id))
	if err != nil {
		c.JSON(http.StatusOK, errmsg.Error())
	}

	c.JSON(http.StatusOK, errmsg.Success(article))
}

func ListArticles(c *gin.Context) {
	pageNum, err := strconv.Atoi(c.Query("pageNum"))
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
	if err != nil {
		c.JSON(http.StatusOK, errmsg.Error())
	}
	articles, err := model.ListAllArticles(pageNum, pageSize)
	if err != nil {
		c.JSON(http.StatusOK, errmsg.Error())
	}

	c.JSON(http.StatusOK, errmsg.Success(articles))
}