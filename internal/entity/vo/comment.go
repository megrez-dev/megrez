package vo

import (
	"time"

	"github.com/megrez/internal/dao"
	"github.com/megrez/internal/entity/po"
)

type Comment struct {
	ID          uint
	ParentID    uint
	ArticleID   uint
	PageID      uint
	Content     string
	Status      int
	Author      *Author
	HasChild    bool
	SubComments []*SubComment
	CreatedAt   time.Time
}

type SubComment struct {
	ID          uint
	ParentID    uint
	ArticleID   uint
	PageID      uint
	Content     string
	Status      int
	Author      *Author
	ReplyAuthor *Author
	CreatedAt   time.Time
}

func GetCommentFromPO(po po.Comment) (*Comment, error) {
	comment := &Comment{}
	comment.ID = po.ID
	comment.ParentID = po.ParentID
	comment.ArticleID = po.ArticleID
	comment.Content = po.Content
	comment.Status = po.Status
	comment.CreatedAt = po.CreatedAt
	dao, err := dao.GetDAO()
	if err != nil {
		return comment, err
	}
	author, err := dao.GetAuthorByID(po.AuthorID)
	if err != nil {
		return comment, err
	}
	comment.Author = GetAuthorFromPO(author)
	subCommentPOs, err := dao.ListCommentsByRootID(comment.ID)
	if err != nil {
		return comment, err
	}
	subComments := []*SubComment{}
	for _, subCommentPO := range subCommentPOs {
		subComment, err := GetSubCommentFromPO(subCommentPO)
		if err != nil {
			return comment, err
		}
		subComments = append(subComments, subComment)
	}
	comment.SubComments = subComments
	if len(subComments) == 0 {
		comment.HasChild = false
	} else {
		comment.HasChild = true
	}
	return comment, nil
}

func GetSubCommentFromPO(po po.Comment) (*SubComment, error) {
	subComment := &SubComment{}
	subComment.ID = po.ID
	subComment.ParentID = po.ParentID
	subComment.ArticleID = po.ArticleID
	subComment.Content = po.Content
	subComment.Status = po.Status
	subComment.CreatedAt = po.CreatedAt
	dao, err := dao.GetDAO()
	if err != nil {
		return subComment, err
	}
	author, err := dao.GetAuthorByID(po.AuthorID)
	if err != nil {
		return subComment, err
	}
	subComment.Author = GetAuthorFromPO(author)
	// set reply author
	parent, err := dao.GetCommentByID(po.ParentID)
	if err != nil {
		return subComment, err
	}
	replyAuthorPO, err := dao.GetAuthorByID(parent.AuthorID)
	if err != nil {
		return subComment, err
	}
	replyAuthor := GetAuthorFromPO(replyAuthorPO)
	subComment.ReplyAuthor = replyAuthor
	return subComment, nil
}
