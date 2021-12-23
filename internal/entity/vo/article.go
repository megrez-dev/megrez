package vo

import (
	"time"

	"github.com/megrez/internal/dao"
	"github.com/megrez/internal/entity/po"
)

type IndexArticle struct {
	ID          uint
	Title       string
	Summary     string
	Thumb       string
	Private     bool
	Category    *BriefCategory
	CommentsNum int64
	TopPriority uint
	Visits      uint
	Likes       uint
	WordCount   uint
	Status      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ArticleDetial struct {
	ID            uint
	Title         string
	FormatContent string
	Thumb         string
	Private       bool
	Category      *BriefCategory
	Tags          []*BriefTag
	Comments      []*Comment
	CommentsNum   int64
	Visits        uint
	Likes         uint
	WordCount     uint
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Pre           *NextPreArticle
	Next          *NextPreArticle
	Page          *Page
	Status        int
}

type NextPreArticle struct {
	ID    uint
	Title string
	Thumb string
}

func GetIndexArticleFromPO(po *po.Article) (IndexArticle, error) {
	vo := IndexArticle{}
	vo.ID = po.ID
	vo.Title = po.Title
	vo.Summary = po.Summary
	vo.Thumb = po.Thumb
	vo.Private = po.Private
	vo.TopPriority = po.TopPriority
	vo.Visits = po.Visits
	vo.Status = po.Status
	vo.CreatedAt = po.CreatedAt
	vo.UpdatedAt = po.UpdatedAt

	dao, err := dao.GetDAO()
	if err != nil {
		return vo, err
	}
	// TODO: 默认 CategoryID = 1
	category, err := dao.GetCategoryByID(po.CategoryID)
	if err != nil {
		return vo, err
	}
	commentsNum, err := dao.CountCommentsByArticleID(po.ID)
	if err != nil {
		return vo, err
	}
	vo.CommentsNum = commentsNum
	categoryVO := GetBriefCategoryFromPO(&category)
	vo.Category = categoryVO
	return vo, nil
}

func GetArticleDetailFromPO(po *po.Article, pageNum, pageSize int) (ArticleDetial, error) {
	vo := ArticleDetial{}
	vo.ID = po.ID
	vo.Title = po.Title
	vo.FormatContent = po.FormatContent
	vo.Thumb = po.Thumb
	vo.Private = po.Private
	vo.Visits = po.Visits
	vo.Likes = po.Likes
	vo.WordCount = po.WordCount
	vo.CreatedAt = po.CreatedAt
	vo.UpdatedAt = po.UpdatedAt
	vo.Status = po.Status

	dao, err := dao.GetDAO()
	if err != nil {
		return vo, err
	}
	// TODO: 默认 CategoryID = 1
	categoryPO, err := dao.GetCategoryByID(po.CategoryID)
	if err != nil {
		return vo, err
	}
	category := GetBriefCategoryFromPO(&categoryPO)
	vo.Category = category
	// TODO: Tags
	// vo.Tags = ...
	commentPOs, err := dao.ListRootCommentsByArticleID(po.ID, pageNum, pageSize)
	if err != nil {
		return vo, err
	}
	comments := []*Comment{}
	for _, commmentPO := range commentPOs {
		comment, err := GetCommentFromPO(commmentPO)
		if err != nil {
			return vo, err
		}
		comments = append(comments, comment)
	}
	vo.Comments = comments
	comentsNum, err := dao.CountCommentsByArticleID(po.ID)
	if err != nil {
		return vo, err
	}
	vo.CommentsNum = comentsNum

	pre, err := dao.GetArticleByID(vo.ID - 1)
	if err == nil {
		vo.Pre = GetNextPreArticleFromPO(pre)
	}
	next, err := dao.GetArticleByID(vo.ID + 1)
	if err == nil {
		vo.Next = GetNextPreArticleFromPO(next)
	}
	rootCount, err := dao.CountRootCommentsByArticleID(po.ID)
	if err == nil {
		page := CaculatePage(pageNum, pageSize, int(rootCount))
		vo.Page = page
	}
	return vo, nil
}

func GetNextPreArticleFromPO(po po.Article) *NextPreArticle {
	vo := &NextPreArticle{}
	vo.ID = po.ID
	vo.Title = po.Title
	vo.Thumb = po.Thumb
	return vo
}
