package po

import "gorm.io/gorm"

type Page struct {
	Name     string `gorm:"type:varchar(255)"`
	Slug     string `gorm:"type:varchar(255);uniqueIndex"`
	Thumb    string `gorm:"type:varchar(255);uniqueIndex"`
	Password string `gorm:"type:varchar(255)"`
	Private  bool
	Visits   uint `gorm:"type:int(11)"`
	Likes    uint `gorm:"type:int(11)"`
	Status   int  `gorm:"type:int(11)"`
	gorm.Model
}
