package router

import (
	"log"
	"strconv"

	"github.com/flosch/pongo2/v4"
	"github.com/gin-gonic/gin"
	"github.com/megrez/pkg/dao"
	"github.com/megrez/pkg/entity/vo"
	"gorm.io/gorm"
)

func routePage(g *gin.Engine, dao *dao.DAO) {
	DAO = dao
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
	page, err := DAO.GetPageBySlug(slug)
	if err == gorm.ErrRecordNotFound {
		c.Redirect(404, "/404")
	}else {
		c.Redirect(500, "/error")
	}


	commentPOs, err := DAO.ListRootCommentsByPageID(page.ID, pageNum, pageSize)
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
	commentsNum, err := DAO.CountRootCommentsByPageID(page.ID)
	if err != nil {
		c.Redirect(500, "/error")
	}
	page.CommentsNum = commentsNum
	pagination := vo.CaculatePagination(pageNum, pageSize, int(commentsNum))
	template := page.slug + ".html"
	c.HTML(200, template, pongo2.Context{"page": page, "pagination": pagination, "comments": comments, "global": globalOption})
}
