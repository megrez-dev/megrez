package po

import "gorm.io/gorm"

type Tag struct {
	Name        string `gorm:"type:varchar(255)"`
	Slug        string `gorm:"type:varchar(255);uniqueIndex"`
	Description string `gorm:"type:varchar(255)"`
	Status      int    `gorm:"type:int(11)"`
	gorm.Model
}
