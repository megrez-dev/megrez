package model

import "gorm.io/gorm"

type Tag struct {
	Name        string `gorm:"type:varchar(255)"`
	Slug        string `gorm:"type:varchar(255);uniqueIndex"`
	Description string `gorm:"type:varchar(255)"`
	Status      int    `gorm:"type:int(11)"`
	gorm.Model
}

// GetTagsByArticleID return tags by articleID
func GetTagsByArticleID(aid uint) ([]Tag, error) {
	tags := []Tag{}
	result := db.Where("id in (?)", db.Table("article_tags").Select("tag_id").Where("article_id = ?", aid)).Find(&tags)
	return tags, result.Error
}

// ListAllTags
func ListAllTags() ([]Tag, error) {
	tags := []Tag{}
	result := db.Find(&tags)
	return tags, result.Error
}

func CreateTag(tag *Tag) error {
	result := db.Create(tag)
	return result.Error
}
