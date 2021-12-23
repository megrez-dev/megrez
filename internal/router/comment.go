package router

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/megrez/internal/dao"
	"github.com/megrez/internal/entity/po"
)

func routeComment(g *gin.Engine, dao *dao.DAO) {
	DAO = dao
	g.POST("/admin/comment", createCommentForAdmin)
	g.POST("/article/:articleID/comment", createCommentForArticle)
	g.POST("/page/:pageID/comment", createCommentForPage)
}

func createCommentForArticle(c *gin.Context) {
	articleID, err := strconv.Atoi(c.Param("articleID"))
	if err != nil {
		c.Redirect(500, "/error")
	}
	var parentID, rootID int
	if c.PostForm("parent") != "" {
		parentID, err = strconv.Atoi(c.PostForm("parent"))
		if err != nil {
			c.Redirect(500, "/error")
		}
	}
	if c.PostForm("root") != "" {
		rootID, err = strconv.Atoi(c.PostForm("root"))
		if err != nil {
			c.Redirect(500, "/error")
		}
	}
	text := c.PostForm("text")
	// TODO: create author
	// authorName := c.PostForm("author")
	// authorMail := c.PostForm("mail")
	// authorURL := c.PostForm("url")
	comment := &po.Comment{
		ArticleID: uint(articleID),
		Content:   text,
		RootID:    uint(rootID),
		ParentID:  uint(parentID),
		AuthorID:  1,
		Type:      1,
		Status:    0,
	}
	err = DAO.CreateComment(comment)
	if err != nil {
		log.Println("create comment failed, err: ", err)
		c.JSON(500, "failed")
	}
	url := fmt.Sprintf("/article/%d#comment-%d", articleID, comment.ID)
	c.Redirect(302, url)
}

func createCommentForPage(c *gin.Context) {
	pageID, err := strconv.Atoi(c.Param("pageID"))
	if err != nil {
		c.Redirect(500, "/error")
	}
	var parentID, rootID int
	if c.PostForm("parent") != "" {
		parentID, err = strconv.Atoi(c.PostForm("parent"))
		if err != nil {
			c.Redirect(500, "/error")
		}
	}
	if c.PostForm("root") != "" {
		rootID, err = strconv.Atoi(c.PostForm("root"))
		if err != nil {
			c.Redirect(500, "/error")
		}
	}
	text := c.PostForm("text")
	// TODO: create author
	// authorName := c.PostForm("author")
	// authorMail := c.PostForm("mail")
	// authorURL := c.PostForm("url")
	comment := &po.Comment{
		PageID:   uint(pageID),
		Content:  text,
		RootID:   uint(rootID),
		ParentID: uint(parentID),
		AuthorID: 1,
		Type:     2,
		Status:   0,
	}
	page, err := DAO.GetPageByID(pageID)
	if err != nil {
		log.Println("get page failed, err: ", err)
		c.JSON(500, "failed")
	}
	err = DAO.CreateComment(comment)
	if err != nil {
		log.Println("create comment failed, err: ", err)
		c.JSON(500, "failed")
	}
	url := fmt.Sprintf("/%s#comment-%d", page.Slug, comment.ID)
	c.Redirect(302, url)
}

func createCommentForAdmin(c *gin.Context) {
	tp, err := strconv.Atoi(c.PostForm("type"))
	if err != nil {
		c.Redirect(500, "/error")
	}
	var articleID, pageID int
	if tp == 1 {
		articleID, err = strconv.Atoi(c.Param("articleID"))
		if err != nil {
			c.Redirect(500, "/error")
		}
	} else if tp == 2 {
		pageID, err = strconv.Atoi(c.Param("pageID"))
		if err != nil {
			c.Redirect(500, "/error")
		}
	}
	if err != nil {
		c.Redirect(500, "/error")
	}
	parentID, err := strconv.Atoi(c.PostForm("parent"))
	if err != nil {
		c.Redirect(500, "/error")
	}
	rootID, err := strconv.Atoi(c.PostForm("root"))
	if err != nil {
		c.Redirect(500, "/error")
	}
	text := c.PostForm("text")
	// TODO: create author
	// authorName := c.PostForm("author")
	// authorMail := c.PostForm("mail")
	// authorURL := c.PostForm("url")
	comment := &po.Comment{
		ArticleID: uint(articleID),
		PageID:    uint(pageID),
		Content:   text,
		RootID:    uint(rootID),
		ParentID:  uint(parentID),
		AuthorID:  1,
		Type:      1,
		Status:    0,
	}
	err = DAO.CreateComment(comment)
	if err != nil {
		log.Println("create comment failed, err: ", err)
		c.JSON(500, "failed")
	}
	c.JSON(200, "success")
}
