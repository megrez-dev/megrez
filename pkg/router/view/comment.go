package view

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/megrez/pkg/entity/vo"
	"github.com/megrez/pkg/model"
)

func RouteComment(g *gin.Engine) {
	g.POST("/article/:articleID/comment", createCommentForArticle)
	g.POST("/page/:pageID/comment", createCommentForPage)
}

func createCommentForArticle(c *gin.Context) {
	articleID, err := strconv.Atoi(c.Param("articleID"))
	if err != nil {
		c.Redirect(http.StatusInternalServerError, "/error")
	}
	var parentID, rootID int
	if c.PostForm("parent") != "" {
		parentID, err = strconv.Atoi(c.PostForm("parent"))
		if err != nil {
			c.Redirect(http.StatusInternalServerError, "/error")
		}
	}
	if c.PostForm("root") != "" {
		rootID, err = strconv.Atoi(c.PostForm("root"))
		if err != nil {
			c.Redirect(http.StatusInternalServerError, "/error")
		}
	}
	// TODO: 要不要存储头像到 DB
	agent := c.Request.UserAgent()
	log.Println("agent:", agent)
	comment := &model.Comment{
		ArticleID: uint(articleID),
		Content:   c.PostForm("text"),
		RootID:    uint(rootID),
		ParentID:  uint(parentID),
		Author:    c.PostForm("author"),
		Mail:      c.PostForm("mail"),
		Site:      c.PostForm("url"),
		Agent:     c.Request.UserAgent(),
		IP:        c.ClientIP(),
		Type:      1,
		Status:    0,
	}
	err = model.CreateComment(comment)
	if err != nil {
		log.Println("create comment failed, err: ", err)
		c.Redirect(http.StatusInternalServerError, "/error")
	}
	// calculate pagination
	pageSizeStr, err := model.GetOptionByKey(vo.OptionComentsPageSize)
	if err != nil {
		c.Redirect(http.StatusInternalServerError, "/error")
	}
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		c.Redirect(http.StatusInternalServerError, "/error")
	}
	rootComments, err := model.ListRootCommentsByArticleID(comment.ArticleID, 0, 0)
	if err != nil {
		c.Redirect(http.StatusInternalServerError, "/error")
	}
	var index int
	for i, rootComment := range rootComments {
		if rootComment.ID == comment.ID || rootComment.ID == comment.RootID {
			index = i + 1
			break
		}
	}
	pagination := (index + pageSize - 1) / pageSize
	url := fmt.Sprintf("/article/%d/comment-page/%d#comment-%d", comment.ArticleID, pagination, comment.ID)
	c.Redirect(http.StatusFound, url)
}

func createCommentForPage(c *gin.Context) {
	pageID, err := strconv.Atoi(c.Param("pageID"))
	if err != nil {
		c.Redirect(http.StatusInternalServerError, "/error")
	}
	var parentID, rootID int
	if c.PostForm("parent") != "" {
		parentID, err = strconv.Atoi(c.PostForm("parent"))
		if err != nil {
			c.Redirect(http.StatusInternalServerError, "/error")
		}
	}
	if c.PostForm("root") != "" {
		rootID, err = strconv.Atoi(c.PostForm("root"))
		if err != nil {
			c.Redirect(http.StatusInternalServerError, "/error")
		}
	}
	text := c.PostForm("text")
	comment := &model.Comment{
		PageID:   uint(pageID),
		Content:  text,
		RootID:   uint(rootID),
		ParentID: uint(parentID),
		Author:   c.PostForm("author"),
		Mail:     c.PostForm("mail"),
		Site:     c.PostForm("url"),
		Agent:    c.Request.UserAgent(),
		IP:       c.ClientIP(),
		Type:     2,
		Status:   0,
	}
	page, err := model.GetPageByID(uint(pageID))
	if err != nil {
		log.Println("get page failed, err: ", err.Error())
		c.Redirect(http.StatusInternalServerError, "/error")
	}
	err = model.CreateComment(comment)
	if err != nil {
		log.Println("create comment failed, err: ", err.Error())
		c.Redirect(http.StatusInternalServerError, "/error")
	}
	// calculate pagination
	pageSizeStr, err := model.GetOptionByKey(vo.OptionComentsPageSize)
	if err != nil {
		log.Println("get option pageSize failed, err: ", err.Error())
		c.Redirect(http.StatusInternalServerError, "/error")
	}
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		log.Println(err.Error())
		c.Redirect(http.StatusInternalServerError, "/error")
	}
	rootComments, err := model.ListRootCommentsByPageID(comment.PageID, 0, 0)
	if err != nil {
		c.Redirect(http.StatusInternalServerError, "/error")
	}
	var index int
	for i, rootComment := range rootComments {
		if rootComment.ID == comment.ID || rootComment.ID == comment.RootID {
			index = i + 1
			break
		}
	}
	pagination := (index + pageSize - 1) / pageSize
	url := fmt.Sprintf("/%s/comment-page/%d#comment-%d", page.Slug, pagination, comment.ID)
	c.Redirect(http.StatusFound, url)
}
