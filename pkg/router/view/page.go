package view

import (
	"log"
	"net/http"
	"strconv"

	"github.com/flosch/pongo2/v6"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/megrez/pkg/entity/vo"
	"github.com/megrez/pkg/model"
)

func RoutePage(g *gin.Engine) {
	g.GET("page/:slug", page)
	g.GET("page/:slug/comment-page/:pageNum", page)
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
			c.Redirect(http.StatusInternalServerError, "/error")
		}
	}
	pageSize = 10
	slug := c.Param("slug")
	page, err := model.GetPageBySlugAndType(slug, model.PageTypeCustom)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.Redirect(404, "/404")
		} else {
			c.Redirect(http.StatusInternalServerError, "/error")
		}
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
	template := page.Slug + ".html"
	c.HTML(http.StatusOK, template, pongo2.Context{
		"page": page,
		"pagination": pagination,
		"comments": comments,
		"global": globalOption,
	})
}
