package model

import (
	"gorm.io/gorm"
	"time"
)

type Comment struct {
	ID         uint      `gorm:"primarykey" json:"id"`
	ArticleID  uint      `gorm:"type:int(11)" json:"articleID"`
	PageID     uint      `gorm:"type:int(11)" json:"pageID"`
	Content    string    `gorm:"type:longtext" json:"content"`
	RootID     uint      `gorm:"type:int(11)" json:"rootID"`
	ParentID   uint      `gorm:"type:int(11)" json:"parentID"`
	Type       int       `gorm:"type:int(11)" json:"type"`
	Author     string    `gorm:"type:varchar(63)" json:"author"`
	Role       int       `gorm:"type:int(11)" json:"role"`
	Mail       string    `gorm:"type:varchar(63)" json:"mail"`
	Site       string    `gorm:"type:varchar(63)" json:"site"`
	Agent      string    `gorm:"type:varchar(1023)" json:"agent"`
	IP         string    `gorm:"type:varchar(20)" json:"ip"`
	Status     int       `gorm:"type:int(11)" json:"status"`
	CreateTime time.Time `gorm:"default:NULL" json:"createTime"`
	UpdateTime time.Time `gorm:"default:NULL" json:"updateTime"`
}

// CreateComment handle create comment
func CreateComment(tx *gorm.DB, comment *Comment) error {
	if tx == nil {
		tx = db
	}
	if tx.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	result := tx.Create(comment)
	return result.Error
}

// GetCommentByID return comment by ID
func GetCommentByID(id uint) (Comment, error) {
	if db.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	comment := Comment{}
	result := db.First(&comment, id)
	return comment, result.Error
}

// ListAllComments return all comments
func ListAllComments(pageNum, pageSize int) ([]Comment, error) {
	if db.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	var comments []Comment
	result := db.Offset(pageSize * (pageNum - 1)).Limit(pageSize).Find(&comments)
	return comments, result.Error
}

// ListCommentsByRootID return comments by rootID
func ListCommentsByRootID(rid uint) ([]Comment, error) {
	if db.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	var comments []Comment
	result := db.Find(&comments, "root_id = ?", rid)
	return comments, result.Error
}

// ListRootCommentsByArticleID return root comments by parentID
func ListRootCommentsByArticleID(aid uint, pageNum, pageSize int) ([]Comment, error) {
	if db.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	var comments []Comment
	var result *gorm.DB
	if pageNum == 0 && pageSize == 0 {
		result = db.Where("article_id = ? AND root_id = ?", aid, 0).Find(&comments)
	} else {
		result = db.Where("article_id = ? AND root_id = ?", aid, 0).
			Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&comments)
	}
	return comments, result.Error
}

// ListRootCommentsByPageID return root comments by pageID
func ListRootCommentsByPageID(pid uint, pageNum, pageSize int) ([]Comment, error) {
	if db.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	var comments []Comment
	var result *gorm.DB
	if pageNum == 0 && pageSize == 0 {
		result = db.Where("page_id = ? AND root_id = ?", pid, 0).Find(&comments)
	} else {
		result = db.Where("page_id = ? AND root_id = ?", pid, 0).
			Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&comments)
	}
	return comments, result.Error
}

// ListLatestComments return latest comments
func ListLatestComments() ([]Comment, error) {
	if db.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	var comments []Comment
	result := db.Limit(8).Find(&comments)
	return comments, result.Error
}

// CountCommentsByArticleID return count of comments by articleID
func CountCommentsByArticleID(aid uint) (int64, error) {
	if db.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	var count int64
	result := db.Model(&Comment{}).Where("article_id = ?", aid).Count(&count)
	return count, result.Error
}

// CountCommentsByPageID return count of comments by pageID
func CountCommentsByPageID(pid uint) (int64, error) {
	if db.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	var count int64
	result := db.Model(&Comment{}).Where("page_id = ?", pid).Count(&count)
	return count, result.Error
}

// CountRootCommentsByArticleID count root comments by articleID
func CountRootCommentsByArticleID(aid uint) (int64, error) {
	if db.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	var count int64
	result := db.Model(&Comment{}).Where("article_id = ? AND root_id = ?", aid, 0).Count(&count)
	return count, result.Error
}

// CountRootCommentsByPageID count root comments by pageID
func CountRootCommentsByPageID(pid uint) (int64, error) {
	if db.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	var count int64
	result := db.Model(&Comment{}).Where("page_id = ? AND root_id = ?", pid, 0).Count(&count)
	return count, result.Error
}
