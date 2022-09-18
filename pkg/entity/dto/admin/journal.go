package admin

import (
	"github.com/megrez/pkg/model"
	"strings"
	"time"
)

type CreateJournalForm struct {
	Content string   `json:"content"`
	Images  []string `json:"images"`
	Private bool     `json:"private"`
	Status  int      `json:"status"`
}

type JournalDTO struct {
	ID         uint      `json:"id"`
	Content    string    `json:"content"`
	Images     []string  `json:"images"`
	Private    bool      `json:"private"`
	Visits     int64     `json:"visits"`
	Likes      int64     `json:"likes"`
	Status     int       `json:"status"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
}

func (dto *JournalDTO) LoadFromModel(journal model.Journal) {
	dto.ID = journal.ID
	dto.Content = journal.Content
	if journal.Images != "" {
		dto.Images = strings.Split(journal.Images, ";")
	}
	dto.Private = journal.Private
	dto.Visits = journal.Visits
	dto.Likes = journal.Likes
	dto.Status = journal.Status
	dto.CreateTime = journal.CreateTime
	dto.UpdateTime = journal.UpdateTime
}
