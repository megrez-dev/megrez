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

// ListArticlesByIDs return articles by ids
func (dao *DAO) ListArticlesByIDs(ids []uint) ([]po.Article, error) {
	articles := []po.Article{}
	result := dao.db.Find(&articles, ids)
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

// ListArticlesByCategoryID return articles by categoryID
func (dao *DAO) ListArticlesByCategoryID(cid uint, pageNum, pageSize int) ([]po.Article, error) {
	articles := []po.Article{}
	result := dao.db.Order("created_at desc").Where("category_id = ?", cid).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articles)
	return articles, result.Error
}

// CountAllArticles return count of all articles
func (dao *DAO) CountAllArticles() (int64, error) {
	var count int64
	result := dao.db.Model(&po.Article{}).Count(&count)
	return count, result.Error
}

// CountArticlesByCategoryID
func (dao *DAO) CountArticlesByCategoryID(cid uint) (int64, error) {
	var count int64
	result := dao.db.Model(&po.Article{}).Where("category_id = ?", cid).Count(&count)
	return count, result.Error
}

// CountArticlesByTagID
func (dao *DAO) CountArticlesByTagID(tid uint) (int64, error) {
	var count int64
	result := dao.db.Model(&po.ArticleTag{}).Where("tag_id = ?", tid).Count(&count)
	return count, result.Error
}

// CreateArticle handle create article
func (dao *DAO) CreateArticle(article *po.Article) error {
	result := dao.db.Create(article)
	return result.Error
}
