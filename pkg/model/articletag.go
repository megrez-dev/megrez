package model

import "gorm.io/gorm"

type ArticleTag struct {
	ID        uint `gorm:"primarykey" json:"id"`
	ArticleID uint `gorm:"type:int(11)" json:"articleID"`
	TagID     uint `gorm:"type:int(11)" json:"tagID"`
}

func CreateArticleTag(tx *gorm.DB, aid, tid uint) error {
	if tx == nil {
		tx = db
	}
	if tx.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	articleTag := ArticleTag{
		ArticleID: aid,
		TagID:     tid,
	}
	result := tx.Create(&articleTag)
	return result.Error
}

func DeleteArticleTag(tx *gorm.DB, aid, tid uint) error {
	if tx == nil {
		tx = db
	}
	if tx.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	result := tx.Where("article_id = ? AND tag_id = ?", aid, tid).Delete(&ArticleTag{})
	return result.Error
}

// DeleteArticleTagsByArticleID deletes all article tags by article id
func DeleteArticleTagsByArticleID(tx *gorm.DB, aid uint) error {
	if tx == nil {
		tx = db
	}
	if tx.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	result := tx.Where("article_id = ?", aid).Delete(&ArticleTag{})
	return result.Error
}
