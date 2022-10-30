package admin

import (
	"time"

	"github.com/megrez/pkg/model"
)

type CommentListDTO struct {
	ID         uint       `json:"id"`
	Article    ArticleDTO `json:"article"`
	Page       PageDTO    `json:"page"`
	Content    string     `json:"content"`
	Author     string     `json:"author"`
	IP         string     `json:"ip"`
	URL        string     `json:"url"`
	Mail       string     `json:"mail"`
	Status     int        `json:"status"`
	Type       string     `json:"type"`
	RootID     uint       `json:"rootID"`
	ParentID   uint       `json:"parentID"`
	CreateTime time.Time  `json:"createTime"`
}

type CreateCommentForm struct {
	Content    string    `json:"content"`
	Type       string    `json:"type"`
	ArticleID  uint      `json:"articleID"`
	PageID     uint      `json:"pageID"`
	RootID     uint      `json:"rootID"`
	ParentID   uint      `json:"parentID"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
}

func (dto *CommentListDTO) LoadFromModel(comment model.Comment) error {
	dto.ID = comment.ID
	dto.Content = comment.Content
	dto.Author = comment.Author
	dto.URL = comment.URL
	dto.Mail = comment.Email
	dto.IP = comment.IP
	dto.Type = comment.Type
	dto.Status = comment.Status
	dto.RootID = comment.RootID
	dto.ParentID = comment.ParentID
	dto.CreateTime = comment.CreateTime
	if comment.Type == model.CommentTypeArticle {
		article, err := model.GetArticleByID(comment.ArticleID)
		if err != nil {
			return err
		}
		if err := dto.Article.LoadFromModel(article); err != nil {
			return err
		}
	} else {
		page, err := model.GetPageByID(comment.PageID)
		if err != nil {
			return err
		}
		if err := dto.Page.LoadFromModel(page); err != nil {
			return err
		}
	}
	return nil
}
