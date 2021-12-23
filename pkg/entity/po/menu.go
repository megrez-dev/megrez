package po

import "gorm.io/gorm"

type Menu struct {
	Name     string `gorm:"type:varchar(255)"`
	Slug     string `gorm:"type:varchar(255)"`
	PageID   bool
	Priority uint `gorm:"type:int(11)"`
	Status   int  `gorm:"type:int(11)"`
	gorm.Model
}
