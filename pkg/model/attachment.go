package model

import "time"

type Attachment struct {
	ID         uint      `gorm:"primarykey" json:"id"`
	URL        string    `gorm:"type:varchar(255)" json:"url"`
	ThumbURL   string    `gorm:"type:varchar(255)" json:"ThumbURL"`
	FileName   string    `gorm:"type:varchar(255)" json:"fileName"`
	Ext        string    `gorm:"type:varchar(255)" json:"ext"`
	Size       int64     `gorm:"type:bigint" json:"size"`
	Width      int       `gorm:"type:int(11)" json:"width"`
	Height     int       `gorm:"type:int(11)" json:"height"`
	Type       int       `gorm:"type:int(11)" json:"type"`
	UploadTime time.Time `gorm:"default:NULL" json:"uploadTime"`
}

func CreateAttachment(attachment *Attachment) error {
	return db.Create(attachment).Error
}

func ListAttachments(pageNum, pageSize int) ([]Attachment, error) {
	var attachments []Attachment
	result := db.Offset(pageSize * (pageNum - 1)).Limit(pageSize).Find(&attachments)
	return attachments, result.Error
}

func CountAllAttachments() (int64, error) {
	var count int64
	result := db.Model(&Attachment{}).Count(&count)
	return count, result.Error
}
