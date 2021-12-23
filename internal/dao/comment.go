package dao

import (
	"github.com/megrez/internal/entity/po"
	"gorm.io/gorm"
)

// GetCommentByID return comment by ID
func (dao *DAO) GetCommentByID(id uint) (po.Comment, error) {
	comment := po.Comment{}
	result := dao.db.First(&comment, id)
	return comment, result.Error
}

// ListCommentsByRootID return comments by rootID
func (dao *DAO) ListCommentsByRootID(rid uint) ([]po.Comment, error) {
	comments := []po.Comment{}
	result := dao.db.Find(&comments, "root_id = ?", rid)
	return comments, result.Error
}

// ListCommentsByArticleID return root comments by parentID
func (dao *DAO) ListRootCommentsByArticleID(aid uint, pageNum, pageSize int) ([]po.Comment, error) {
	comments := []po.Comment{}
	var result *gorm.DB
	if pageNum == 0 && pageSize == 0 {
		result = dao.db.Order("created_at desc").Where("article_id = ? AND root_id = ?", aid, 0).Find(&comments)
	} else {
		result = dao.db.Order("created_at desc").Where("article_id = ? AND root_id = ?", aid, 0).
			Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&comments)
	}
	return comments, result.Error
}

// ListRootCommentsByPageID return root comments by pageID
func (dao *DAO) ListRootCommentsByPageID(pid uint, pageNum, pageSize int) ([]po.Comment, error) {
	comments := []po.Comment{}
	var result *gorm.DB
	if pageNum == 0 && pageSize == 0 {
		result = dao.db.Order("created_at desc").Where("page_id = ? AND root_id = ?", pid, 0).Find(&comments)
	} else {
		result = dao.db.Order("created_at desc").Where("page_id = ? AND root_id = ?", pid, 0).
			Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&comments)
	}
	return comments, result.Error
}

// ListLatestComments return latest comments
func (dao *DAO) ListLatestComments() ([]po.Comment, error) {
	comments := []po.Comment{}
	result := dao.db.Order("created_at desc").Limit(8).Find(&comments)
	return comments, result.Error
}

// CountCommentsByArticleID return count of comments by articleID
func (dao *DAO) CountCommentsByArticleID(aid uint) (int64, error) {
	var count int64
	result := dao.db.Model(&po.Comment{}).Where("article_id = ?", aid).Count(&count)
	return count, result.Error
}

// CountRootCommentsByArticleID count root comments by articleID
func (dao *DAO) CountRootCommentsByArticleID(aid uint) (int64, error) {
	var count int64
	result := dao.db.Model(&po.Comment{}).Where("article_id = ? AND root_id = ?", aid, 0).Count(&count)
	return count, result.Error
}

// CreateComment handle create comment
func (dao *DAO) CreateComment(comment *po.Comment) error {
	result := dao.db.Create(&comment)
	return result.Error
}
