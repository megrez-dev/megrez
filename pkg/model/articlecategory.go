package model

type ArticleCategory struct {
	ID         uint `gorm:"primarykey" json:"id"`
	ArticleID  uint `gorm:"type:int(11)" json:"articleID"`
	CategoryID uint `gorm:"type:int(11)" json:"categoryID"`
}

func CreateArticleCategory(aid, cid uint) error {
	if db.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	articleCategory := ArticleCategory{
		ArticleID:  aid,
		CategoryID: cid,
	}
	result := db.Create(&articleCategory)
	return result.Error
}

func DeleteArticleCategory(aid, cid uint) error {
	if db.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	result := db.Where("article_id = ? AND category_id = ?", aid, cid).Delete(&ArticleCategory{})
	return result.Error
}
