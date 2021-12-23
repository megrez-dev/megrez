package po

import (
	"gorm.io/gorm"
)

type Article struct {
	Title           string `gorm:"type:varchar(255)"`
	OriginalContent string `gorm:"type:longtext"`
	FormatContent   string `gorm:"type:longtext"`
	Summary         string `gorm:"type:longtext"`
	Slug            string `gorm:"type:varchar(255);uniqueIndex"`
	Password        string `gorm:"type:varchar(255)"`
	Thumb           string `gorm:"type:varchar(255)"`
	Private         bool
	CategoryID      uint `gorm:"type:int(11)"`
	TopPriority     uint `gorm:"type:int(11)"`
	Visits          uint `gorm:"type:int(11)"`
	Likes           uint `gorm:"type:int(11)"`
	WordCount       uint `gorm:"type:int(11)"`
	Status          int  `gorm:"type:int(11)"`
	gorm.Model
}
