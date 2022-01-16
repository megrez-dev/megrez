package model

import "time"

type Page struct {
	ID         uint      `gorm:"primarykey" json:"id"`
	Name       string    `gorm:"type:varchar(255)" json:"name"`
	Slug       string    `gorm:"type:varchar(255);uniqueIndex" json:"slug"`
	Cover      string    `gorm:"type:varchar(1023)" json:"cover"`
	Password   string    `gorm:"type:varchar(255)" json:"password"`
	Private    bool      `json:"private"`
	Visits     int64     `gorm:"type:int(11)" json:"visits"`
	Likes      int64     `gorm:"type:int(11)" json:"likes"`
	Status     int       `gorm:"type:int(11)" json:"status"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
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
