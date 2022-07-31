package model

import (
	"gorm.io/gorm"
	"time"
)

const (
	AttachmentTypeLocal     = "local"
	AttachmentTypeQcloudCos = "qcloud_cos"
	AttachmentTypeAliyunOss = "aliyun_oss"
	AttachmentTypeHuaweiObs = "huawei_obs"
	AttachmentTypeQiniuyun  = "qiniuyun"
	AttachmentTypeYoupaiyun = "youpaiyun"
)

type Attachment struct {
	ID         uint      `gorm:"primarykey" json:"id"`
	URL        string    `gorm:"type:varchar(255)" json:"url"`
	ThumbURL   string    `gorm:"type:varchar(255)" json:"ThumbURL"`
	FileName   string    `gorm:"type:varchar(255)" json:"fileName"`
	Ext        string    `gorm:"type:varchar(255)" json:"ext"`
	Size       int64     `gorm:"type:bigint" json:"size"`
	Width      int       `gorm:"type:int(11)" json:"width"`
	Height     int       `gorm:"type:int(11)" json:"height"`
	Type       string    `gorm:"type:varchar(255)" json:"type"`
	UploadTime time.Time `gorm:"default:NULL" json:"uploadTime"`
}

func CreateAttachment(tx *gorm.DB, attachment *Attachment) error {
	if tx == nil {
		tx = db
	}
	if tx.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	return tx.Create(attachment).Error
}

func ListAttachments(pageNum, pageSize int) ([]Attachment, error) {
	if db.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	var attachments []Attachment
	result := db.Order("upload_time DESC").Offset(pageSize * (pageNum - 1)).Limit(pageSize).Find(&attachments)
	return attachments, result.Error
}

func CountAllAttachments() (int64, error) {
	if db.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	var count int64
	result := db.Model(&Attachment{}).Count(&count)
	return count, result.Error
}
