package view

import (
	"log"
	"strconv"

	"github.com/flosch/pongo2/v4"
	"github.com/gin-gonic/gin"
	"github.com/megrez/pkg/entity/vo"
	"github.com/megrez/pkg/model"
)

func RouteCategory(g *gin.Engine) {
	g.GET("/category/:slug", listArticlesByCategory)
	g.GET("/category/:slug/:pageNum", listArticlesByCategory)
	g.POST("/admin/category", createCategory)
}

func listArticlesByCategory(c *gin.Context) {
	var pageNum, pageSize int
	var err error
	if c.Param("pageNum") == "" {
		pageNum = 1
	} else {
		pageNum, err = strconv.Atoi(c.Param("pageNum"))
		if err != nil {
			log.Println("incorrect param pageNum, err:", err)
			// TODO: 应该是 4XX?
			c.Redirect(500, "/error")
		}
	}
	if c.Param("pageSize") == "" {
		pageSize = 20
	} else {
		pageSize, err = strconv.Atoi(c.Param("pageSize"))
		if err != nil {
			log.Println("incorrect param pageSize, err:", err)
			c.Redirect(500, "/error")
		}
	}
	categorySlug := c.Param("slug")
	if categorySlug == "" {
		c.Redirect(404, "/404")
	}
	category, err := model.GetCategoryBySlug(categorySlug)
	if err != nil {
		c.Redirect(500, "/error")
	}
	articlePOs, err := model.ListArticlesByCategoryID(category.ID, pageNum, pageSize)
	if err != nil {
		c.Redirect(500, "/error")
	}
	articles := []*vo.CommonArticle{}
	for _, articlePO := range articlePOs {
		article := vo.GetCommonArticleFromPO(articlePO)
		articles = append(articles, article)
	}
	globalOption, err := vo.GetGlobalOption()
	if err != nil {
		c.Redirect(500, "/error")
	}
	articlesNum, err := model.CountArticlesByCategoryID(category.ID)
	if err != nil {
		c.Redirect(500, "/error")
	}
	pagination := vo.CaculatePagination(pageNum, pageSize, int(articlesNum))
	c.HTML(200, "category.html", pongo2.Context{"category": category, "pagination": pagination, "articles": articles, "global": globalOption})
}

func createCategory(c *gin.Context) {
	category := &model.Category{
		Name:   "默认分类",
		Slug:   "default",
		Status: 0,
	}
	err := model.CreateCategory(category)
	if err != nil {
		log.Println("create category failed, err: ", err)
		c.JSON(500, "failed")
	}
	c.JSON(200, "success")
}
