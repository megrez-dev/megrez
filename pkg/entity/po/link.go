package po

import "gorm.io/gorm"

type Link struct {
	Name     string `gorm:"type:varchar(255)"`
	Addr     string `gorm:"type:varchar(255)"`
	Logo     string `gorm:"type:varchar(255)"`
	Priority uint   `gorm:"type:int(11)"`
	Status   int    `gorm:"type:int(11)"`
	gorm.Model
}
