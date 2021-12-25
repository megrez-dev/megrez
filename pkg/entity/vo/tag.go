package vo

import (
	"github.com/megrez/pkg/dao"
	"github.com/megrez/pkg/entity/po"
)

type BriefTag struct {
	Name string
	Slug string
}

type TagWithArticlesNum struct {
	Name        string
	Slug        string
	ArticlesNum int64
}

func GetBriefTagFromPO(po po.Tag) *BriefTag {
	return &BriefTag{
		Name: po.Name,
		Slug: po.Slug,
	}
}

func GetTagWithArticlesNumFromPO(po po.Tag) (*TagWithArticlesNum, error) {
	dao, err := dao.GetDAO()
	if err != nil {
		return nil, err
	}
	articlesNum, err := dao.CountArticlesByTagID(po.ID)
	if err != nil {
		return nil, err
	}
	tag := &TagWithArticlesNum{
		Name:        po.Name,
		Slug:        po.Slug,
		ArticlesNum: articlesNum,
	}
	return tag, nil
}
