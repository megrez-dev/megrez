package model

import "gorm.io/gorm"

type ArticleCategory struct {
	ID         uint `gorm:"primarykey" json:"id"`
	ArticleID  uint `gorm:"type:int(11)" json:"articleID"`
	CategoryID uint `gorm:"type:int(11)" json:"categoryID"`
}

func CreateArticleCategory(tx *gorm.DB, aid, cid uint) error {
	if tx == nil {
		tx = db
	}
	if tx.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	articleCategory := ArticleCategory{
		ArticleID:  aid,
		CategoryID: cid,
	}
	result := tx.Create(&articleCategory)
	return result.Error
}

func DeleteArticleCategory(tx *gorm.DB, aid, cid uint) error {
	if tx == nil {
		tx = db
	}
	if tx.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	result := tx.Where("article_id = ? AND category_id = ?", aid, cid).Delete(&ArticleCategory{})
	return result.Error
}
