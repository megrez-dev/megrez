package model

import (
	"gorm.io/gorm"
	"time"
)

type Page struct {
	ID              uint      `gorm:"primarykey" json:"id"`
	Name            string    `gorm:"type:varchar(255)" json:"name"`
	Slug            string    `gorm:"type:varchar(255);uniqueIndex" json:"slug"`
	Cover           string    `gorm:"type:varchar(1023)" json:"cover"`
	Password        string    `gorm:"type:varchar(255)" json:"password"`
	Private         bool      `json:"private"`
	Visits          int64     `gorm:"type:int(11)" json:"visits"`
	Likes           int64     `gorm:"type:int(11)" json:"likes"`
	Status          int       `gorm:"type:int(11)" json:"status"`
	OriginalContent string    `gorm:"type:longtext" json:"originalContent"`
	FormatContent   string    `gorm:"type:longtext" json:"formatContent"`
	CreateTime      time.Time `gorm:"default:NULL" json:"createTime"`
	UpdateTime      time.Time `gorm:"default:NULL" json:"updateTime"`
}

// GetPageByID return page by pageID
func GetPageByID(id uint) (Page, error) {
	if db.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	page := Page{}
	result := db.First(&page, id)
	return page, result.Error
}

// GetPageBySlug return page by slug
func GetPageBySlug(slug string) (Page, error) {
	if db.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	page := Page{}
	result := db.First(&page, "`slug` = ?", slug)
	return page, result.Error
}

// CreatePage create a new page
func CreatePage(tx *gorm.DB, page *Page) error {
	if tx == nil {
		tx = db
	}
	if tx.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	result := tx.Create(page)
	return result.Error
}
