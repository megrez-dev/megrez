package model

import (
	"gorm.io/gorm"
	"time"
)

type Article struct {
	ID              uint      `gorm:"primarykey" json:"id"`
	Title           string    `gorm:"type:varchar(255)" json:"title"`
	OriginalContent string    `gorm:"type:longtext" json:"originalContent"`
	FormatContent   string    `gorm:"type:longtext" json:"formatContent"`
	Summary         string    `gorm:"type:longtext" json:"summary"`
	Slug            string    `gorm:"type:varchar(255);uniqueIndex" json:"slug"`
	Password        string    `gorm:"type:varchar(255)" json:"password"`
	Cover           string    `gorm:"type:varchar(255)" json:"cover"`
	Private         bool      `json:"private"`
	AllowedComment  bool      `json:"allowedComment"`
	TopPriority     int       `gorm:"type:int(11)" json:"topPriority"`
	Visits          int64     `gorm:"type:int(11)" json:"visits"`
	Likes           int64     `gorm:"type:int(11)" json:"likes"`
	WordCount       int64     `gorm:"type:int(11)" json:"wordCount"`
	SeoKeywords     string    `gorm:"type:varchar(255)" json:"seoKeywords"`
	SeoDescription  string    `gorm:"type:varchar(1023)" json:"seoDescription"`
	Status          int       `gorm:"type:int(11)" json:"status"`
	PublishTime     time.Time `gorm:"default:NULL" json:"publishTime"`
	EditTime        time.Time `gorm:"default:NULL" json:"editTime"`
}

// CreateArticle handle create article
func CreateArticle(tx *gorm.DB, article *Article) error {
	if tx == nil {
		tx = db
	}
	if tx.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	result := tx.Create(article)
	return result.Error
}

// UpdateArticleByID update article by id and data
func UpdateArticleByID(tx *gorm.DB, article *Article) error {
	if tx == nil {
		tx = db
	}
	if tx.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	result := tx.Model(&article).Select("*").Omit("publish_time").Updates(article)
	return result.Error
}

// DeleteArticleByID delete article by id
func DeleteArticleByID(tx *gorm.DB, id uint) error {
	if tx == nil {
		tx = db
	}
	if tx.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	result := tx.Delete(&Article{}, id)
	return result.Error
}

// GetArticleByID return article by id
func GetArticleByID(id uint) (Article, error) {
	if db.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	article := Article{}
	result := db.First(&article, id)
	return article, result.Error
}

// GetArticleBySlug return article by slug
func GetArticleBySlug(slug string) (Article, error) {
	if db.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	article := Article{}
	result := db.First(&article, "`slug` = ?", slug)
	return article, result.Error
}

// ListArticlesByIDs return articles by ids
func ListArticlesByIDs(ids []uint) ([]Article, error) {
	if db.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	var articles []Article
	result := db.Order("publish_time DESC").Find(&articles, ids)
	return articles, result.Error
}

// ListAllArticles return all articles
func ListAllArticles(pageNum, pageSize int) ([]Article, error) {
	if db.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	var articles []Article
	result := db.Order("publish_time DESC").Offset(pageSize * (pageNum - 1)).Limit(pageSize).Find(&articles)
	return articles, result.Error
}

// ListLatestArticles return latest articles
func ListLatestArticles() ([]Article, error) {
	if db.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	var articles []Article
	result := db.Order("publish_time DESC").Limit(8).Find(&articles)
	return articles, result.Error
}

// ListArticlesByCategoryID return articles by categoryID
func ListArticlesByCategoryID(cid uint, pageNum, pageSize int) ([]Article, error) {
	if db.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	var articles []Article
	result := db.Where("id in (?)", db.Table("article_categories").Select("article_id").Where("category_id = ?", cid)).Order("publish_time DESC").Find(&articles)
	return articles, result.Error
}

// CountAllArticles return count of all articles
func CountAllArticles() (int64, error) {
	if db.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	var count int64
	result := db.Model(&Article{}).Count(&count)
	return count, result.Error
}

// CountArticlesByCategoryID return count for articles by categoryID
func CountArticlesByCategoryID(cid uint) (int64, error) {
	if db.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	var count int64
	result := db.Model(&ArticleCategory{}).Where("category_id = ?", cid).Count(&count)
	return count, result.Error
}

// CountArticlesByTagID return count for articles by tagID
func CountArticlesByTagID(tid uint) (int64, error) {
	if db.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	var count int64
	result := db.Model(&ArticleTag{}).Where("tag_id = ?", tid).Count(&count)
	return count, result.Error
}
