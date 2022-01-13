package model

import "gorm.io/gorm"

type Comment struct {
	ArticleID  uint   `gorm:"type:int(11)"`
	PageID     uint   `gorm:"type:int(11)"`
	Content    string `gorm:"type:longtext"`
	RootID     uint   `gorm:"type:int(11)"`
	ParentID   uint   `gorm:"type:int(11)"`
	Type       int    `gorm:"type:int(11)"`
	Author string `gorm:"type:varchar(63)"`
	Role       int    `gorm:"type:int(11)"`
	Mail       string `gorm:"type:varchar(63)"`
	Site       string `gorm:"type:varchar(63)"`
	Agent      string `gorm:"type:varchar(1023)"`
	IP         string `gorm:"type:varchar(20)"`
	Status     int    `gorm:"type:int(11)"`
	gorm.Model
}

// GetCommentByID return comment by ID
func GetCommentByID(id uint) (Comment, error) {
	comment := Comment{}
	result := db.First(&comment, id)
	return comment, result.Error
}

// ListCommentsByRootID return comments by rootID
func ListCommentsByRootID(rid uint) ([]Comment, error) {
	comments := []Comment{}
	result := db.Find(&comments, "root_id = ?", rid)
	return comments, result.Error
}

// ListCommentsByArticleID return root comments by parentID
func ListRootCommentsByArticleID(aid uint, pageNum, pageSize int) ([]Comment, error) {
	comments := []Comment{}
	var result *gorm.DB
	if pageNum == 0 && pageSize == 0 {
		result = db.Order("created_at desc").Where("article_id = ? AND root_id = ?", aid, 0).Find(&comments)
	} else {
		result = db.Order("created_at desc").Where("article_id = ? AND root_id = ?", aid, 0).
			Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&comments)
	}
	return comments, result.Error
}

// ListRootCommentsByPageID return root comments by pageID
func ListRootCommentsByPageID(pid uint, pageNum, pageSize int) ([]Comment, error) {
	comments := []Comment{}
	var result *gorm.DB
	if pageNum == 0 && pageSize == 0 {
		result = db.Order("created_at desc").Where("page_id = ? AND root_id = ?", pid, 0).Find(&comments)
	} else {
		result = db.Order("created_at desc").Where("page_id = ? AND root_id = ?", pid, 0).
			Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&comments)
	}
	return comments, result.Error
}

// ListLatestComments return latest comments
func ListLatestComments() ([]Comment, error) {
	comments := []Comment{}
	result := db.Order("created_at desc").Limit(8).Find(&comments)
	return comments, result.Error
}

// CountCommentsByArticleID return count of comments by articleID
func CountCommentsByArticleID(aid uint) (int64, error) {
	var count int64
	result := db.Model(&Comment{}).Where("article_id = ?", aid).Count(&count)
	return count, result.Error
}

// CountCommentsByPageID return count of comments by pageID
func CountCommentsByPageID(pid uint) (int64, error) {
	var count int64
	result := db.Model(&Comment{}).Where("page_id = ?", pid).Count(&count)
	return count, result.Error
}

// CountRootCommentsByArticleID count root comments by articleID
func CountRootCommentsByArticleID(aid uint) (int64, error) {
	var count int64
	result := db.Model(&Comment{}).Where("article_id = ? AND root_id = ?", aid, 0).Count(&count)
	return count, result.Error
}

// CountRootCommentsByPageID count root comments by pageID
func CountRootCommentsByPageID(pid uint) (int64, error) {
	var count int64
	result := db.Model(&Comment{}).Where("page_id = ? AND root_id = ?", pid, 0).Count(&count)
	return count, result.Error
}

// CreateComment handle create comment
func CreateComment(comment *Comment) error {
	result := db.Create(comment)
	return result.Error
}
