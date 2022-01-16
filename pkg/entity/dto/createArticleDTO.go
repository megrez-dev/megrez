package dto

import "github.com/megrez/pkg/model"

type CreateArticleDTO struct {
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

func (dto CreateArticleDTO) Transfer2Model() model.Article {
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
		seoKeywords = seoKeywords + ";" + seoKeyword
	}
	article.SeoKeywords = seoKeywords
	return article
}
