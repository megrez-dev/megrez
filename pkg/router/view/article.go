package view

import (
	"log"
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
			// TODO: 应该是 4XX?
			c.Redirect(500, "/error")
		}
	}
	if c.Param("pageSize") == "" {
		pageSize = 8
	} else {
		pageSize, err = strconv.Atoi(c.Param("pageSize"))
		if err != nil {
			log.Println("incorrect param pageSize, err:", err)
			c.Redirect(500, "/error")
		}
	}
	articlePOs, err := model.ListAllArticles(pageNum, pageSize)
	if err != nil {
		log.Println("get articles from db failed, err:", err)
		c.Redirect(500, "/error")
	}
	articleVOs := []vo.IndexArticle{}
	for _, articlePO := range articlePOs {
		articleVO, err := vo.GetIndexArticleFromPO(&articlePO)
		if err != nil {
			c.Redirect(500, "/error")
		}
		articleVOs = append(articleVOs, articleVO)
	}
	globalOption, err := vo.GetGlobalOption()
	if err != nil {
		c.Redirect(500, "/error")
	}
	// pageInfo
	count, err := model.CountAllArticles()
	if err != nil {
		c.Redirect(500, "/error")
	}
	page := vo.CaculatePagination(pageNum, pageSize, int(count))
	// TODO: 过滤器格式化时间
	c.HTML(200, "index.html", pongo2.Context{"articles": articleVOs, "global": globalOption, "page": page})
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
			// TODO: 应该是 4XX?
			c.Redirect(500, "/error")
		}
	}

	if c.Param("pageSize") == "" {
		pageSize = 8
	} else {
		pageSize, err = strconv.Atoi(c.Param("pageSize"))
		if err != nil {
			log.Println("incorrect param pageSize, err:", err)
			c.Redirect(500, "/error")
		}
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println("incorrect param id, err:", err)
		c.Redirect(500, "/error")
	}
	articlePO, err := model.GetArticleByID(uint(id))
	if err != nil {
		log.Println("query article from db failed, err: ", err)
	}
	articleDetial, err := vo.GetArticleDetailFromPO(articlePO, pageNum, pageSize)
	if err != nil {
		log.Println("get article detail failed, err:", err)
		c.Redirect(500, "/error")
	}
	globalOption, err := vo.GetGlobalOption()
	if err != nil {
		c.Redirect(500, "/error")
	}
	c.HTML(200, "article.html", pongo2.Context{"article": articleDetial, "global": globalOption})
}
