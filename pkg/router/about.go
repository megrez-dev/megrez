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
	g.GET("/about/comment-page/:pageNum", about)
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
	pageSize = 8

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
		ID          uint
		Name        string
		Slug        string
		Visits      int64
		CommentsNum int64
	}{
		ID:   1,
		Name: "关于",
		Slug: "about",
		Visits: 10086,
	}
	commentsNum, err := DAO.CountRootCommentsByPageID(page.ID)
	if err != nil {
		c.Redirect(500, "/error")
	}
	page.CommentsNum = commentsNum
	pagination := vo.CaculatePagination(pageNum, pageSize, int(commentsNum))
	c.HTML(200, "about.html", pongo2.Context{"page": page, "pagination": pagination, "comments": comments, "global": globalOption})
}
