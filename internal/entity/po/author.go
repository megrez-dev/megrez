package po

import "gorm.io/gorm"

type Author struct {
	Name   string `gorm:"type:varchar(255)"`
	Mail   string `gorm:"type:varchar(255)"`
	Site   string `gorm:"type:varchar(255)"`
	Role   int    `gorm:"type:int(11)"`
	Avatar string `gorm:"type:varchar(255)"`
	IP     string `gorm:"type:varchar(255)"`
	Status int    `gorm:"type:int(11)"`
	gorm.Model
}
