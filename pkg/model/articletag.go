package model

type ArticleTag struct {
	ID         uint   `gorm:"primarykey" json:"id"`
	ArticleID  uint `gorm:"type:int(11)" json:"articleID"`
	TagID      uint `gorm:"type:int(11)" json:"tagID"`
}

func CreateArticleTag(aid, tid uint) error {
	articleTag := ArticleTag{
		ArticleID:  aid,
		TagID: tid,
	}
	result := db.Create(&articleTag)
	return result.Error
}

func DeleteArticleTag(aid, tid uint) error {
	result := db.Where("article_id = ? AND tag_id = ?", aid, tid).Delete(&ArticleTag{})
	return result.Error
}