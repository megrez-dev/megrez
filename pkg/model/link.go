package model

import "gorm.io/gorm"

type Link struct {
	Name     string `gorm:"type:varchar(255)"`
	Addr     string `gorm:"type:varchar(255)"`
	Logo     string `gorm:"type:varchar(255)"`
	Priority uint   `gorm:"type:int(11)"`
	Status   int    `gorm:"type:int(11)"`
	gorm.Model
}

// ListAllLinks return all links
func ListAllLinks() ([]Link, error) {
	links := []Link{}
	result := db.Find(&links)
	return links, result.Error
}

// CreateLink handle create link
func CreateLink(link *Link) error {
	result := db.Create(link)
	return result.Error
}
