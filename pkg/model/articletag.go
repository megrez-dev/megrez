package model

import "gorm.io/gorm"

type ArticleTag struct {
	ArticleID string `gorm:"type:int(11)" json:"articleID"`
	TagID     string `gorm:"type:int(11)" json:"tagID"`
	gorm.Model
}
