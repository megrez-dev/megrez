package router

import (
	"log"
	"strconv"

	"github.com/flosch/pongo2/v4"
	"github.com/gin-gonic/gin"
	"github.com/megrez/pkg/dao"
	"github.com/megrez/pkg/entity/vo"
)

func routeAbout(g *gin.Engine, dao *dao.DAO) {
	DAO = dao
	g.GET("/about", about)
}

func about(c *gin.Context) {
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

	commentPOs, err := DAO.ListRootCommentsByPageID(1, pageNum, pageSize)
	if err != nil {
		c.Redirect(500, "/error")
	}
	var comments []*vo.Comment
	for _, commentPO := range commentPOs {
		comment, err := vo.GetCommentFromPO(commentPO)
		if err != nil {
			c.Redirect(500, "/error")
		}
		comments = append(comments, comment)
	}

	globalOption, err := vo.GetGlobalOption()
	if err != nil {
		c.Redirect(500, "/error")
	}
	page := struct {
		ID   uint
		Slug string
	}{
		ID:   1,
		Slug: "about",
	}
	pagination
	c.HTML(200, "about.html", pongo2.Context{"page": page, "comments": comments, "global": globalOption})
}
