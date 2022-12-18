package view

import (
	"log"
	"net/http"
	"strconv"

	"github.com/flosch/pongo2/v6"
	"github.com/gin-gonic/gin"
	"github.com/megrez/pkg/entity/vo"
	"github.com/megrez/pkg/model"
)

func RouteAbout(g *gin.Engine) {
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
	page, err := model.GetPageBySlugAndType("about", model.PageTypeBuildIn)
	if err != nil {
		c.Redirect(http.StatusInternalServerError, "/error")
	}
	pageVO := vo.GetPageFromPO(page)
	commentPOs, err := model.ListRootCommentsByPageID(page.ID, pageNum, pageSize)
	if err != nil {
		c.Redirect(http.StatusInternalServerError, "/error")
	}
	var comments []*vo.Comment
	for _, commentPO := range commentPOs {
		comment, err := vo.GetCommentFromPO(commentPO)
		if err != nil {
			c.Redirect(http.StatusInternalServerError, "/error")
		}
		comments = append(comments, comment)
	}
	globalOption, err := vo.GetGlobalOption()
	if err != nil {
		c.Redirect(http.StatusInternalServerError, "/error")
	}
	commentsNum, err := model.CountRootCommentsByPageID(page.ID)
	if err != nil {
		c.Redirect(http.StatusInternalServerError, "/error")
	}
	pageVO.CommentsNum = commentsNum
	pagination := vo.CalculatePagination(pageNum, pageSize, int(commentsNum))
	c.HTML(http.StatusOK, "about.html", pongo2.Context{
		"page":       pageVO,
		"comments":   comments,
		"pagination": pagination,
		"global":     globalOption,
	})
}
