package model

import (
	"log"

	"gorm.io/gorm"
)

type Article struct {
	Title           string `gorm:"type:varchar(255)"`
	OriginalContent string `gorm:"type:longtext"`
	FormatContent   string `gorm:"type:longtext"`
	Summary         string `gorm:"type:longtext"`
	Slug            string `gorm:"type:varchar(255);uniqueIndex"`
	Password        string `gorm:"type:varchar(255)"`
	Thumb           string `gorm:"type:varchar(255)"`
	Private         bool
	CategoryID      uint  `gorm:"type:int(11)"`
	TopPriority     uint  `gorm:"type:int(11)"`
	Visits          int64 `gorm:"type:int(11)"`
	Likes           int64 `gorm:"type:int(11)"`
	WordCount       int64 `gorm:"type:int(11)"`
	Status          int   `gorm:"type:int(11)"`
	gorm.Model
}

// GetArticleByID return article by id
func GetArticleByID(id uint) (Article, error) {
	article := Article{}
	result := db.First(&article, id)
	return article, result.Error
}

// ListArticlesByIDs return articles by ids
func ListArticlesByIDs(ids []uint) ([]Article, error) {
	articles := []Article{}
	result := db.Find(&articles, ids)
	return articles, result.Error
}

// ListAllArticles return all articles
func ListAllArticles(pageNum, pageSize int) ([]Article, error) {
	articles := []Article{}
	log.Println("数据库链接 db:", db)
	result := db.Offset(pageSize * (pageNum - 1)).Limit(pageSize).Find(&articles)
	return articles, result.Error
}

// ListLatestArticles return latest articles
func ListLatestArticles() ([]Article, error) {
	articles := []Article{}
	result := db.Order("created_at desc").Limit(8).Find(&articles)
	return articles, result.Error
}

// ListArticlesByCategoryID return articles by categoryID
func ListArticlesByCategoryID(cid uint, pageNum, pageSize int) ([]Article, error) {
	articles := []Article{}
	result := db.Order("created_at desc").Where("category_id = ?", cid).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articles)
	return articles, result.Error
}

// CountAllArticles return count of all articles
func CountAllArticles() (int64, error) {
	var count int64
	result := db.Model(&Article{}).Count(&count)
	return count, result.Error
}

// CountArticlesByCategoryID
func CountArticlesByCategoryID(cid uint) (int64, error) {
	var count int64
	result := db.Model(&Article{}).Where("category_id = ?", cid).Count(&count)
	return count, result.Error
}

// CountArticlesByTagID
func CountArticlesByTagID(tid uint) (int64, error) {
	var count int64
	result := db.Model(&ArticleTag{}).Where("tag_id = ?", tid).Count(&count)
	return count, result.Error
}

// CreateArticle handle create article
func CreateArticle(article *Article) error {
	result := db.Create(article)
	return result.Error
}
