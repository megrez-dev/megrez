package model

import "gorm.io/gorm"

type ArticleTag struct {
	ArticleID string `gorm:"type:int(11)"`
	TagID     string `gorm:"type:int(11)"`
	gorm.Model
}
