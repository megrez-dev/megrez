package view

import (
	"fmt"
	"github.com/megrez/pkg/log"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
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
	agent := c.Request.UserAgent()
	log.Debug("agent:", agent)
	comment := &model.Comment{
		ArticleID:  uint(articleID),
		Content:    c.PostForm("text"),
		RootID:     uint(rootID),
		ParentID:   uint(parentID),
		Author:     c.PostForm("author"),
		Email:      c.PostForm("email"),
		Site:       c.PostForm("url"),
		Agent:      c.Request.UserAgent(),
		IP:         c.ClientIP(),
		Type:       1,
		Status:     0,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	tx := model.BeginTx()
	err = model.CreateComment(tx, comment)
	if err != nil {
		log.Error("create comment failed, err: ", err)
		tx.Rollback()
		c.Redirect(http.StatusInternalServerError, "/error")
	}
	// calculate pagination
	commentsPageSizeStr, err := model.GetOptionByKey(model.OptionKeyCommentsPageSize)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			commentsPageSizeStr = "10"
		} else {
			tx.Rollback()
			c.Redirect(http.StatusInternalServerError, "/error")
		}
	}
	commentPageSize, err := strconv.Atoi(commentsPageSizeStr)
	if err != nil {
		tx.Rollback()
		c.Redirect(http.StatusInternalServerError, "/error")
	}
	rootComments, err := model.ListRootCommentsByArticleID(comment.ArticleID, 0, 0)
	if err != nil {
		tx.Rollback()
		c.Redirect(http.StatusInternalServerError, "/error")
	}
	var index int
	for i, rootComment := range rootComments {
		if rootComment.ID == comment.ID || rootComment.ID == comment.RootID {
			index = i + 1
			break
		}
	}
	pagination := (index + commentPageSize - 1) / commentPageSize
	url := fmt.Sprintf("/article/%d/comment-page/%d#comment-%d", comment.ArticleID, pagination, comment.ID)
	tx.Commit()
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
		Email:    c.PostForm("email"),
		Site:     c.PostForm("url"),
		Agent:    c.Request.UserAgent(),
		IP:       c.ClientIP(),
		Type:     2,
		Status:   0,
	}
	page, err := model.GetPageByID(uint(pageID))
	if err != nil {
		log.Error("get page failed, err: ", err.Error())
		c.Redirect(http.StatusInternalServerError, "/error")
	}
	tx := model.BeginTx()
	err = model.CreateComment(tx, comment)
	if err != nil {
		log.Error("create comment failed, err: ", err.Error())
		tx.Rollback()
		c.Redirect(http.StatusInternalServerError, "/error")
	}
	// calculate pagination
	commentsPageSizeStr, err := model.GetOptionByKey(model.OptionKeyCommentsPageSize)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			commentsPageSizeStr = "10"
		} else {
			tx.Rollback()
			c.Redirect(http.StatusInternalServerError, "/error")
		}
	}
	commentsPageSize, err := strconv.Atoi(commentsPageSizeStr)
	if err != nil {
		log.Error(err.Error())
		tx.Rollback()
		c.Redirect(http.StatusInternalServerError, "/error")
	}
	rootComments, err := model.ListRootCommentsByPageID(comment.PageID, 0, 0)
	if err != nil {
		tx.Rollback()
		c.Redirect(http.StatusInternalServerError, "/error")
	}
	var index int
	for i, rootComment := range rootComments {
		if rootComment.ID == comment.ID || rootComment.ID == comment.RootID {
			index = i + 1
			break
		}
	}
	pagination := (index + commentsPageSize - 1) / commentsPageSize
	url := fmt.Sprintf("/%s/comment-page/%d#comment-%d", page.Slug, pagination, comment.ID)
	tx.Commit()
	c.Redirect(http.StatusFound, url)
}
