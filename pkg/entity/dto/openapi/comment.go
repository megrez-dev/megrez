package openapi

import (
	"github.com/megrez/pkg/log"
	"github.com/megrez/pkg/model"
)

type CommentDTO struct {
	ID          uint            `json:"id"`
	Content     string          `json:"content"`
	Author      string          `json:"author"`
	Avatar      string          `json:"avatar"`
	Mail        string          `json:"mail"`
	URL         string          `json:"url"`
	SubComments []SubCommentDTO `json:"subComments"`
	IP          string          `json:"ip"`
	RootID      uint            `json:"rootID"`
	ParentID    uint            `json:"parentID"`
	Type        string          `json:"type"`
	Status      int             `json:"status"`
	CreateTime  string          `json:"createTime"`
}

type SubCommentDTO struct {
	ID         uint   `json:"id"`
	Content    string `json:"content"`
	Author     string `json:"author"`
	Avatar     string `json:"avatar"`
	Mail       string `json:"mail"`
	URL        string `json:"url"`
	IP         string `json:"ip"`
	Type       string `json:"type"`
	RootID     uint   `json:"rootID"`
	ParentID   uint   `json:"parentID"`
	Status     int    `json:"status"`
	CreateTime string `json:"createTime"`
}

type CreateCommentForm struct {
	Content   string `json:"content"`
	Type      string `json:"type"`
	ArticleID uint   `json:"articleID"`
	PageID    uint   `json:"pageID"`
	Author    string `json:"author"`
	Avatar    string `json:"avatar"`
	Mail      string `json:"mail"`
	URL       string `json:"url"`
	RootID    uint   `json:"rootID"`
	ParentID  uint   `json:"parentID"`
}

func (dto *CommentDTO) LoadFromModel(comment model.Comment) error {
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
	dto.CreateTime = comment.CreateTime.Format("2006-01-02 15:04:05")
	subComments, err := model.ListCommentsByRootID(comment.ID)
	if err != nil {
		log.Error("list subcomments by root id %d failed, err: %s", comment.ID, err.Error())
		return err
	}
	for _, subComment := range subComments {
		subCommentDTO := &SubCommentDTO{}
		subCommentDTO.LoadFromModel(subComment)
		dto.SubComments = append(dto.SubComments, *subCommentDTO)
	}
	return nil
}

func (dto *SubCommentDTO) LoadFromModel(comment model.Comment) {
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
	dto.CreateTime = comment.CreateTime.Format("2006-01-02 15:04:05")
}
