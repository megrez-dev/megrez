package vo

import (
	"time"

	"github.com/megrez/pkg/model"
)

type IndexArticle struct {
	ID          uint
	Title       string
	Summary     string
	Cover       string
	Private     bool
	Categories  []*BriefCategory
	CommentsNum int64
	TopPriority int
	Visits      int64
	Likes       int64
	WordCount   int64
	Status      int
	PublishTime time.Time
	EditTime    time.Time
}

type CommonArticle struct {
	ID          uint
	Title       string
	Slug        string
	Status      int
	Private     bool
	PublishTime time.Time
}

type ArticleDetail struct {
	ID            uint
	Title         string
	FormatContent string
	Cover         string
	Private       bool
	Categories    []*BriefCategory
	Tags          []*BriefTag
	Comments      []*Comment
	CommentsNum   int64
	Visits        int64
	Likes         int64
	WordCount     int64
	PublishTime   time.Time
	EditTime      time.Time
	Pre           *NextPreArticle
	Next          *NextPreArticle
	Page          *Pagination
	Status        int
}

type NextPreArticle struct {
	ID    uint
	Title string
	Cover string
}

func GetIndexArticleFromPO(article *model.Article) (IndexArticle, error) {
	vo := IndexArticle{}
	vo.ID = article.ID
	vo.Title = article.Title
	vo.Summary = article.Summary
	vo.Cover = article.Cover
	vo.Private = article.Private
	vo.TopPriority = article.TopPriority
	vo.Visits = article.Visits
	vo.Status = article.Status
	vo.PublishTime = article.PublishTime
	vo.EditTime = article.EditTime

	categoryPOs, err := model.ListCategoriesByArticleID(article.ID)
	if err != nil {
		return vo, err
	}
	var categories []*BriefCategory
	for _, categoryPO := range categoryPOs {
		category := GetBriefCategoryFromPO(categoryPO)
		categories = append(categories, category)
	}
	vo.Categories = categories

	commentsNum, err := model.CountCommentsByArticleID(article.ID)
	if err != nil {
		return vo, err
	}
	vo.CommentsNum = commentsNum
	return vo, nil
}

func GetCommonArticleFromPO(article model.Article) *CommonArticle {
	commonArticle := &CommonArticle{
		ID:          article.ID,
		Title:       article.Title,
		Slug:        article.Slug,
		Status:      article.Status,
		Private:     article.Private,
		PublishTime: article.PublishTime,
	}
	return commonArticle
}

func GetArticleDetailFromPO(article model.Article, pageNum, pageSize int) (*ArticleDetail, error) {
	vo := &ArticleDetail{}
	vo.ID = article.ID
	vo.Title = article.Title
	vo.FormatContent = article.FormatContent
	vo.Cover = article.Cover
	vo.Private = article.Private
	vo.Visits = article.Visits
	vo.Likes = article.Likes
	vo.WordCount = article.WordCount
	vo.PublishTime = article.PublishTime
	vo.EditTime = article.EditTime
	vo.Status = article.Status

	categoryPOs, err := model.ListCategoriesByArticleID(article.ID)
	if err != nil {
		return vo, err
	}
	var categories []*BriefCategory
	for _, categoryPO := range categoryPOs {
		category := GetBriefCategoryFromPO(categoryPO)
		categories = append(categories, category)
	}
	vo.Categories = categories
	tagPOs, err := model.ListTagsByArticleID(article.ID)
	if err != nil {
		return vo, err
	}
	var tags []*BriefTag
	for _, tagPO := range tagPOs {
		tag := GetBriefTagFromPO(tagPO)
		tags = append(tags, tag)
	}
	vo.Tags = tags
	commentPOs, err := model.ListRootCommentsByArticleID(article.ID, pageNum, pageSize)
	if err != nil {
		return vo, err
	}
	var comments []*Comment
	for _, commentPO := range commentPOs {
		comment, err := GetCommentFromPO(commentPO)
		if err != nil {
			return vo, err
		}
		comments = append(comments, comment)
	}
	vo.Comments = comments
	commentsNum, err := model.CountCommentsByArticleID(article.ID)
	if err != nil {
		return vo, err
	}
	vo.CommentsNum = commentsNum

	pre, err := model.GetPreArticleByID(vo.ID)
	if err == nil {
		vo.Pre = GetNextPreArticleFromPO(pre)
	}
	next, err := model.GetNextArticleByID(vo.ID)
	if err == nil {
		vo.Next = GetNextPreArticleFromPO(next)
	}
	rootCount, err := model.CountRootCommentsByArticleID(article.ID)
	if err == nil {
		page := CalculatePagination(pageNum, pageSize, int(rootCount))
		vo.Page = page
	}
	return vo, nil
}

func GetNextPreArticleFromPO(article model.Article) *NextPreArticle {
	vo := &NextPreArticle{}
	vo.ID = article.ID
	vo.Title = article.Title
	vo.Cover = article.Cover
	return vo
}
