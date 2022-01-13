package view

import (
	"log"
	"strconv"

	"github.com/flosch/pongo2/v4"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/megrez/pkg/entity/vo"
	"github.com/megrez/pkg/model"
)

func RoutePage(g *gin.Engine) {
	g.GET("/:slug", page)
	g.GET("/:slug/comment-page/:pageNum", page)
}

func page(c *gin.Context) {
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
	slug := c.Param("slug")
	page, err := model.GetPageBySlug(slug)
	if err == gorm.ErrRecordNotFound {
		c.Redirect(404, "/404")
	} else {
		c.Redirect(500, "/error")
	}
	pageVO := vo.GetPageFromPO(page)

	commentPOs, err := model.ListRootCommentsByPageID(page.ID, pageNum, pageSize)
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
	commentsNum, err := model.CountRootCommentsByPageID(page.ID)
	if err != nil {
		c.Redirect(500, "/error")
	}
	pageVO.CommentsNum = commentsNum
	pagination := vo.CaculatePagination(pageNum, pageSize, int(commentsNum))
	template := page.Slug + ".html"
	c.HTML(200, template, pongo2.Context{"page": page, "pagination": pagination, "comments": comments, "global": globalOption})
}
