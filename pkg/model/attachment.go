package model

import "time"

type Attachment struct {
	ID         uint      `gorm:"primarykey" json:"id"`
	URL        string    `gorm:"type:varchar(255)" json:"url"`
	ThumbURL   string    `gorm:"type:varchar(255)" json:"ThumbURL"`
	FileName   string    `gorm:"type:varchar(255)" json:"fileName"`
	Type       uint      `gorm:"type:int(11)" json:"type"`
	CreateTime time.Time `gorm:"default:NULL" json:"createTime"`
}
