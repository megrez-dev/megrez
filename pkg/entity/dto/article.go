package dto

import (
	"github.com/megrez/pkg/model"
	"log"
	"strings"
	"time"
)

type ArticleDTO struct {
	ID              uint     `json:"id"`
	Title           string   `json:"title"`
	OriginalContent string   `json:"originalContent"`
	FormatContent   string   `json:"formatContent"`
	Summary         string   `json:"summary"`
	Slug            string   `json:"slug"`
	Password        string   `json:"password"`
	Cover           string   `json:"cover"`
	Private         bool     `json:"private"`
	AllowedComment  bool     `json:"allowedComment"`
	Categories      []uint   `json:"categories"`
	Tags            []uint   `json:"tags"`
	IsTop           bool     `json:"isTop"`
	Visits          int64    `json:"visits"`
	Likes           int64    `json:"likes"`
	WordCount       int64    `json:"wordCount"`
	SeoKeywords     []string `json:"seoKeywords"`
	SeoDescription  string   `json:"seoDescription"`
	Status          int      `json:"status"`
}

func (dto *ArticleDTO) Transfer2Model() model.Article {
	article := model.Article{
		ID:              dto.ID,
		Title:           dto.Title,
		OriginalContent: dto.OriginalContent,
		FormatContent:   dto.FormatContent,
		Summary:         dto.Summary,
		Slug:            dto.Slug,
		Password:        dto.Password,
		Cover:           dto.Cover,
		Private:         dto.Private,
		AllowedComment:  dto.AllowedComment,
		WordCount:       dto.WordCount,
		SeoDescription:  dto.SeoDescription,
		// TODO: 0:正常 1:草稿 2:回收站 ...
		Status: dto.Status,
	}
	// top priority
	if dto.IsTop {
		article.TopPriority = 1
	} else {
		article.TopPriority = 0
	}
	// seo tags
	seoKeywords := ""
	for _, seoKeyword := range dto.SeoKeywords {
		if seoKeywords == "" {
			seoKeywords = seoKeyword
		} else {
			seoKeywords = seoKeywords + ";" + seoKeyword
		}
	}
	article.SeoKeywords = seoKeywords
	return article
}

func (dto *ArticleDTO) LoadFromModel(article model.Article) error {
	dto.ID = article.ID
	dto.Title = article.Title
	dto.OriginalContent = article.OriginalContent
	dto.FormatContent = article.FormatContent
	dto.Summary = article.Summary
	dto.Slug = article.Slug
	dto.Password = article.Password
	dto.Cover = article.Cover
	dto.Private = article.Private
	dto.AllowedComment = article.AllowedComment
	dto.WordCount = article.WordCount
	dto.SeoDescription = article.SeoDescription
	dto.Status = article.Status
	// top priority
	dto.IsTop = article.TopPriority != 0
	dto.SeoKeywords = strings.Split(article.SeoKeywords, ";")
	dto.SeoDescription = article.SeoDescription
	dto.Categories = []uint{}
	categories, err := model.ListCategoriesByArticleID(article.ID)
	if err != nil {
		log.Println(err)
		return err
	}
	for _, category := range categories {
		dto.Categories = append(dto.Categories, category.ID)
	}
	dto.Tags = []uint{}
	tags, err := model.ListTagsByArticleID(article.ID)
	if err != nil {
		log.Println(err)
		return err
	}
	for _, tag := range tags {
		dto.Tags = append(dto.Tags, tag.ID)
	}
	return nil
}

type ArticlesListDTO struct {
	ID              uint             `json:"id"`
	Title           string           `json:"title"`
	OriginalContent string           `json:"originalContent"`
	FormatContent   string           `json:"formatContent"`
	Summary         string           `json:"summary"`
	Slug            string           `json:"slug"`
	Password        string           `json:"password"`
	Cover           string           `json:"cover"`
	Private         bool             `json:"private"`
	AllowedComment  bool             `json:"allowedComment"`
	Categories      []model.Category `json:"categories"`
	Tags            []model.Tag      `json:"tags"`
	CommentsNum     int64            `json:"commentsNum"`
	IsTop           bool             `json:"isTop"`
	Visits          int64            `json:"visits"`
	Likes           int64            `json:"likes"`
	WordCount       int64            `json:"wordCount"`
	Status          int              `json:"status"`
	PublishTime     time.Time        `json:"publishTime"`
	EditTime        time.Time        `json:"editTime"`
}

func (dto *ArticlesListDTO) LoadFromModel(article model.Article) error {
	dto.ID = article.ID
	dto.Title = article.Title
	dto.OriginalContent = article.OriginalContent
	dto.FormatContent = article.FormatContent
	dto.Summary = article.Summary
	dto.Slug = article.Slug
	dto.Password = article.Password
	dto.Cover = article.Cover
	dto.Private = article.Private
	dto.AllowedComment = article.AllowedComment
	dto.Visits = article.Visits
	dto.Likes = article.Likes
	dto.WordCount = article.WordCount
	dto.Status = article.Status
	dto.PublishTime = article.PublishTime
	dto.EditTime = article.EditTime
	dto.IsTop = article.TopPriority != 0
	// categories
	categories, err := model.ListCategoriesByArticleID(article.ID)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	dto.Categories = categories
	// tags
	tags, err := model.ListTagsByArticleID(article.ID)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	dto.Tags = tags
	// commentsNum
	commentsNum, err := model.CountCommentsByArticleID(article.ID)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	dto.CommentsNum = commentsNum
	return nil
}
