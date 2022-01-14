package model

import "gorm.io/gorm"

type Link struct {
	Name     string `gorm:"type:varchar(255)" json:"name"`
	Addr     string `gorm:"type:varchar(255)" json:"addr"`
	Logo     string `gorm:"type:varchar(255)" json:"logo"`
	Priority uint   `gorm:"type:int(11)" json:"priority"`
	Status   int    `gorm:"type:int(11)" json:"status"`
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
