package dto

import (
	"github.com/megrez/pkg/model"
	"time"
)

type CommentListDTO struct {
	ID         uint       `json:"id"`
	Article    ArticleDTO `json:"article"`
	Page       PageDTO    `json:"page"`
	Content    string     `json:"content"`
	Author     string     `json:"author"`
	Site       string     `json:"site"`
	Mail       string     `json:"mail"`
	Status     int        `json:"status"`
	CreateTime time.Time  `json:"createTime"`
}

func (dto *CommentListDTO) LoadFromModel(comment model.Comment) error {
	dto.ID = comment.ID
	dto.Content = comment.Content
	dto.Author = comment.Author
	dto.Site = comment.Site
	dto.Mail = comment.Mail
	if comment.Type == 1 {
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
