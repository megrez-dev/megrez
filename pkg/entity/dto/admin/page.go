package admin

import (
	"github.com/megrez/pkg/model"
	"time"
)

type PageDTO struct {
	ID              uint      `json:"id"`
	Name            string    `json:"name"`
	Slug            string    `json:"slug"`
	Cover           string    `json:"cover"`
	Password        string    `json:"password"`
	Private         bool      `json:"private"`
	Visits          int64     `json:"visits"`
	Likes           int64     `json:"likes"`
	Status          int       `json:"status"`
	OriginalContent string    `json:"originalContent"`
	FormatContent   string    `json:"formatContent"`
	CreateTime      time.Time `json:"createTime"`
	UpdateTime      time.Time `json:"updateTime"`
}

func (dto *PageDTO) LoadFromModel(page model.Page) error {
	dto.ID = page.ID
	dto.Name = page.Name
	dto.Slug = page.Slug
	dto.Cover = page.Cover
	dto.Password = page.Password
	dto.Private = page.Private
	dto.Visits = page.Visits
	dto.Likes = page.Likes
	dto.Status = page.Status
	dto.OriginalContent = page.OriginalContent
	dto.FormatContent = page.FormatContent
	dto.CreateTime = page.CreateTime
	dto.UpdateTime = page.UpdateTime
	return nil
}
