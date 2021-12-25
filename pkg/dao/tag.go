package dao

import (
	"github.com/megrez/pkg/entity/po"
)

// GetTagsByArticleID return tags by articleID
func (dao *DAO) GetTagsByArticleID(aid uint) ([]po.Tag, error) {
	tags := []po.Tag{}
	result := dao.db.Where("id in (?)", dao.db.Table("article_tags").Select("tag_id").Where("article_id = ?", aid)).Find(&tags).Debug()
	return tags, result.Error
}

// ListAllTags
func (dao *DAO) ListAllTags() ([]po.Tag, error) {
	tags := []po.Tag{}
	result := dao.db.Find(&tags)
	return tags, result.Error
}

func (dao *DAO) CreateTag(tag *po.Tag) error {
	result := dao.db.Create(tag)
	return result.Error
}
