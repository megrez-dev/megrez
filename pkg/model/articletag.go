package model

type ArticleTag struct {
	ID        uint `gorm:"primarykey" json:"id"`
	ArticleID uint `gorm:"type:int(11)" json:"articleID"`
	TagID     uint `gorm:"type:int(11)" json:"tagID"`
}

func CreateArticleTag(aid, tid uint) error {
	if db.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	articleTag := ArticleTag{
		ArticleID: aid,
		TagID:     tid,
	}
	result := db.Create(&articleTag)
	return result.Error
}

func DeleteArticleTag(aid, tid uint) error {
	if db.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	result := db.Where("article_id = ? AND tag_id = ?", aid, tid).Delete(&ArticleTag{})
	return result.Error
}

// DeleteArticleTagsByArticleID deletes all article tags by article id
func DeleteArticleTagsByArticleID(aid uint) error {
	if db.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	result := db.Where("article_id = ?", aid).Delete(&ArticleTag{})
	return result.Error
}
