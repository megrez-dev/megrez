package dao

import (
	"github.com/megrez/pkg/entity/po"
)

// GetArticleByID return article by id
func (dao *DAO) GetArticleByID(id uint) (po.Article, error) {
	article := po.Article{}
	result := dao.db.First(&article, id)
	return article, result.Error
}

// GetArticlesByIDs return articles by ids
func (dao *DAO) GetArticlesByIDs(ids []uint) ([]po.Article, error) {
	articles := []po.Article{}
	result := dao.db.Find(articles, ids)
	return articles, result.Error
}

// ListAllArticles return all articles
func (dao *DAO) ListAllArticles(pageNum, pageSize int) ([]po.Article, error) {
	articles := []po.Article{}
	result := dao.db.Offset(pageSize * (pageNum - 1)).Limit(pageSize).Find(&articles)
	return articles, result.Error
}

// ListLatestArticles return latest articles
func (dao *DAO) ListLatestArticles() ([]po.Article, error) {
	articles := []po.Article{}
	result := dao.db.Order("created_at desc").Limit(8).Find(&articles)
	return articles, result.Error
}

// CountAllArticles return count of all articles
func (dao *DAO) CountAllArticles() (int64, error) {
	var count int64
	result := dao.db.Model(&po.Article{}).Count(&count)
	return count, result.Error
}

// CreateArticle handle create article
func (dao *DAO) CreateArticle(article *po.Article) error {
	result := dao.db.Create(&article)
	return result.Error
}
