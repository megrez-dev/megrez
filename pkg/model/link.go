package model

import "time"

type Link struct {
	ID         uint      `gorm:"primarykey" json:"id"`
	Name       string    `gorm:"type:varchar(255)" json:"name"`
	Addr       string    `gorm:"type:varchar(255)" json:"addr"`
	Logo       string    `gorm:"type:varchar(255)" json:"logo"`
	Priority   uint      `gorm:"type:int(11)" json:"priority"`
	Status     int       `gorm:"type:int(11)" json:"status"`
	CreateTime time.Time `gorm:"default:NULL" json:"createTime"`
	UpdateTime time.Time `gorm:"default:NULL" json:"updateTime"`
}

// ListAllLinks return all links
func ListAllLinks() ([]Link, error) {
	var links []Link
	result := db.Find(&links)
	return links, result.Error
}

// CreateLink handle create link
func CreateLink(link *Link) error {
	result := db.Create(link)
	return result.Error
}
