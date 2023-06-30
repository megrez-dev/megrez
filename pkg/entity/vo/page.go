package vo

import (
	"time"

	"github.com/megrez/pkg/model"
)

type Page struct {
	ID          uint
	Name        string
	Slug        string
	Cover       string
	Password    string
	Private     bool
	Visits      int64
	Likes       int64
	CommentsNum int64
	CreateTime  time.Time
	UpdateTime  time.Time
	Status      int
}

func GetPageFromPO(page model.Page) *Page {
	return &Page{
		ID:         page.ID,
		Name:       page.Name,
		Slug:       page.Slug,
		Cover:      page.Cover,
		Password:   page.Password,
		Private:    page.Private,
		Visits:     page.Visits,
		Likes:      page.Likes,
		CreateTime: page.CreateTime,
		UpdateTime: page.UpdateTime,
		Status:     page.Status,
	}
}
