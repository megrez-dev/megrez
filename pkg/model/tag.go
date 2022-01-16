package model

type Tag struct {
	ID          uint   `gorm:"primarykey" json:"id"`
	Name        string `gorm:"type:varchar(255)" json:"name"`
	Slug        string `gorm:"type:varchar(255);uniqueIndex" json:"slug"`
	Description string `gorm:"type:varchar(255)" json:"description"`
	Status      int    `gorm:"type:int(11)" json:"status"`
}

func GetTagByID(tid uint) (Tag, error) {
	tag := Tag{}
	result := db.First(&tag, tid)
	return tag, result.Error
}

// GetTagsByArticleID return tags by articleID
func GetTagsByArticleID(aid uint) ([]Tag, error) {
	var tags []Tag
	result := db.Where("id in (?)", db.Table("article_tags").Select("tag_id").Where("article_id = ?", aid)).Find(&tags)
	return tags, result.Error
}

// GetTagByName return tags by articleID
func GetTagByName(name string) (Tag, error) {
	tag := Tag{}
	result := db.First(&tag, "`name` = ?", name)
	return tag, result.Error
}

// ListAllTags return all tags
func ListAllTags() ([]Tag, error) {
	var tags []Tag
	result := db.Find(&tags)
	return tags, result.Error
}

// ListTagsByPage return tags by page
func ListTagsByPage(pageNum int, pageSize int) ([]Tag, error) {
	var tags []Tag
	result := db.Offset(pageSize * (pageNum - 1)).Limit(pageSize).Find(&tags)
	return tags, result.Error
}

// ListTagsByArticleID return categories by articleID
func ListTagsByArticleID(aid uint) ([]Tag, error) {
	var tags []Tag
	result := db.Where("id in (?)", db.Table("article_tags").Select("tag_id").Where("article_id = ?", aid)).Find(&tags)
	return tags, result.Error
}

func CreateTag(tag *Tag) error {
	result := db.Create(tag)
	return result.Error
}
