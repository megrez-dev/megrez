package model

import "gorm.io/gorm"

type Page struct {
	Name     string `gorm:"type:varchar(255)"`
	Slug     string `gorm:"type:varchar(255);uniqueIndex"`
	Thumb    string `gorm:"type:varchar(1023)"`
	Password string `gorm:"type:varchar(255)"`
	Private  bool
	Visits   int64 `gorm:"type:int(11)"`
	Likes    int64 `gorm:"type:int(11)"`
	Status   int  `gorm:"type:int(11)"`
	gorm.Model
}

// GetPageByID return page by pageID
func GetPageByID(id uint) (Page, error) {
	page := Page{}
	result := db.First(&page, id)
	return page, result.Error
}

// GetPageBySlug return page by slug
func GetPageBySlug(slug string) (Page, error) {
	page := Page{}
	result := db.First(&page, "`slug` = ?", slug)
	return page, result.Error
}
