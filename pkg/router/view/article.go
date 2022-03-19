package view

import (
	"log"
	"net/http"
	"strconv"

	"github.com/flosch/pongo2/v4"
	"github.com/gin-gonic/gin"
	"github.com/megrez/pkg/entity/vo"
	"github.com/megrez/pkg/model"
)

func RouteArticle(g *gin.Engine) {
	g.GET("/", index)
	g.GET("/index/:pageNum", index)
	g.GET("/article/:id", articleDetail)
	g.GET("/article/:id/comment-page/:pageNum", articleDetail)
}

func index(c *gin.Context) {
	var pageNum, pageSize int
	var err error
	if c.Param("pageNum") == "" {
		pageNum = 1
	} else {
		pageNum, err = strconv.Atoi(c.Param("pageNum"))
		if err != nil {
			log.Println("incorrect param pageNum, err:", err)
			c.Redirect(http.StatusInternalServerError, "/error")
		}
	}
	if c.Param("pageSize") == "" {
		pageSize = 10
	} else {
		pageSize, err = strconv.Atoi(c.Param("pageSize"))
		if err != nil {
			log.Println("incorrect param pageSize, err:", err)
			c.Redirect(http.StatusInternalServerError, "/error")
		}
	}
	articlePOs, err := model.ListAllArticles(pageNum, pageSize)
	if err != nil {
		log.Println("get articles from db failed, err:", err)
		c.Redirect(http.StatusInternalServerError, "/error")
	}
	var articleVOs []vo.IndexArticle
	for _, articlePO := range articlePOs {
		articleVO, err := vo.GetIndexArticleFromPO(&articlePO)
		if err != nil {
			c.Redirect(http.StatusInternalServerError, "/error")
		}
		articleVOs = append(articleVOs, articleVO)
	}
	globalOption, err := vo.GetGlobalOption()
	if err != nil {
		c.Redirect(http.StatusInternalServerError, "/error")
	}
	// pageInfo
	count, err := model.CountAllArticles()
	if err != nil {
		c.Redirect(http.StatusInternalServerError, "/error")
	}
	page := vo.CalculatePagination(pageNum, pageSize, int(count))
	c.HTML(http.StatusOK, "index.html", pongo2.Context{"articles": articleVOs, "global": globalOption, "page": page})
}

func articleDetail(c *gin.Context) {
	var pageNum, pageSize int
	var err error
	if c.Param("pageNum") == "" {
		pageNum = 1
	} else {
		pageNum, err = strconv.Atoi(c.Param("pageNum"))
		if err != nil {
			log.Println("incorrect param pageNum, err:", err)
			c.Redirect(http.StatusInternalServerError, "/error")
		}
	}

	if c.Param("pageSize") == "" {
		pageSize = 10
	} else {
		pageSize, err = strconv.Atoi(c.Param("pageSize"))
		if err != nil {
			log.Println("incorrect param pageSize, err:", err)
			c.Redirect(http.StatusInternalServerError, "/error")
		}
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println("incorrect param id, err:", err)
		c.Redirect(http.StatusInternalServerError, "/error")
	}
	articlePO, err := model.GetArticleByID(uint(id))
	if err != nil {
		log.Println("query article from db failed, err: ", err)
	}
	articleDetial, err := vo.GetArticleDetailFromPO(articlePO, pageNum, pageSize)
	if err != nil {
		log.Println("get article detail failed, err:", err)
		c.Redirect(http.StatusInternalServerError, "/error")
	}
	globalOption, err := vo.GetGlobalOption()
	if err != nil {
		c.Redirect(http.StatusInternalServerError, "/error")
	}
	c.HTML(http.StatusOK, "article.html", pongo2.Context{"article": articleDetial, "global": globalOption})
}
