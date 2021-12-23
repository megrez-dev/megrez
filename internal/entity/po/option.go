package po

import "gorm.io/gorm"

type Option struct {
	Key   string `gorm:"type:varchar(255);uniqueIndex"`
	Value string `gorm:"type:varchar(255)"`
	gorm.Model
}
