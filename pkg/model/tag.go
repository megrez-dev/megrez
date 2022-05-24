package model

import "gorm.io/gorm"

type Tag struct {
	ID          uint   `gorm:"primarykey" json:"id"`
	Name        string `gorm:"type:varchar(255)" json:"name"`
	Slug        string `gorm:"type:varchar(255);uniqueIndex" json:"slug"`
	Description string `gorm:"type:varchar(255)" json:"description"`
	Status      int    `gorm:"type:int(11)" json:"status"`
}

func CreateTag(tx *gorm.DB, tag *Tag) error {
	if tx == nil {
		tx = db
	}
	if tx.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	result := tx.Create(tag)
	return result.Error
}

func GetTagByID(tid uint) (Tag, error) {
	if db.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	tag := Tag{}
	result := db.First(&tag, tid)
	return tag, result.Error
}

// GetTagByName return tags by articleID
func GetTagByName(name string) (Tag, error) {
	if db.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	tag := Tag{}
	result := db.First(&tag, "`name` = ?", name)
	return tag, result.Error
}

// ListAllTags return all tags
func ListAllTags() ([]Tag, error) {
	if db.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	var tags []Tag
	result := db.Find(&tags)
	return tags, result.Error
}

// ListTagsByPage return tags by page
func ListTagsByPage(pageNum int, pageSize int) ([]Tag, error) {
	if db.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	var tags []Tag
	result := db.Offset(pageSize * (pageNum - 1)).Limit(pageSize).Find(&tags)
	return tags, result.Error
}

// ListTagsByArticleID return categories by articleID
func ListTagsByArticleID(aid uint) ([]Tag, error) {
	if db.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	var tags []Tag
	result := db.Where("id in (?)", db.Table("article_tags").Select("tag_id").Where("article_id = ?", aid)).Find(&tags)
	return tags, result.Error
}
