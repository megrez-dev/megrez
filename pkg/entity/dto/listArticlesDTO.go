package dto

import (
	"github.com/megrez/pkg/model"
	"strings"
	"time"
)

type ListArticlesDTO struct {
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
	IsTop           bool             `json:"isTop"`
	Visits          int64            `json:"visits"`
	Likes           int64            `json:"likes"`
	WordCount       int64            `json:"wordCount"`
	SeoKeywords     []string         `json:"seoKeywords"`
	SeoDescription  string           `json:"seoDescription"`
	Status          int              `json:"status"`
	PublishTime     time.Time        `json:"publishTime"`
	Total           int64            `json:"total"`
}

func LoadFromModel(article model.Article) (ListArticlesDTO, error) {
	dto := ListArticlesDTO{
		ID:              article.ID,
		Title:           article.Title,
		OriginalContent: article.OriginalContent,
		FormatContent:   article.FormatContent,
		Summary:         article.Summary,
		Slug:            article.Slug,
		Password:        article.Password,
		Cover:           article.Cover,
		Private:         article.Private,
		AllowedComment:  article.AllowedComment,
		Visits:          article.Visits,
		Likes:           article.Likes,
		WordCount:       article.WordCount,
		SeoKeywords:     strings.Split(article.SeoKeywords, ","),
		SeoDescription:  article.SeoDescription,
		// TODO: 0:正常 1:草稿 2:回收站 ...
		Status:      article.Status,
		PublishTime: article.PublishTime,
	}
	// isTop
	dto.IsTop = article.TopPriority == 0
	// categories
	categories, err := model.ListCategoriesByArticleID(article.ID)
	if err != nil {
		return dto, err
	}
	dto.Categories = categories
	// tags
	tags, err := model.ListTagsByArticleID(article.ID)
	if err != nil {
		return dto, err
	}
	dto.Tags = tags
	total, err := model.CountAllArticles()
	if err != nil {
		return dto, err
	}
	dto.Total = total
	return dto, nil
}
