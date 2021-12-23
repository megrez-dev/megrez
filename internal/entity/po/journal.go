package po

import "gorm.io/gorm"

type Journal struct {
	OriginalContent string `gorm:"type:longtext"`
	FormatContent   string `gorm:"type:longtext"`
	Images          string `gorm:"type:varchar(4095)"`
	Private         bool
	Visits          uint `gorm:"type:int(11)"`
	Likes           uint `gorm:"type:int(11)"`
	Status          int  `gorm:"type:int(11)"`
	gorm.Model
}
