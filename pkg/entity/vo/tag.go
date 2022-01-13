package vo

import "github.com/megrez/pkg/model"

type BriefTag struct {
	Name string
	Slug string
}

type TagWithArticlesNum struct {
	Name        string
	Slug        string
	ArticlesNum int64
}

func GetBriefTagFromPO(tag model.Tag) *BriefTag {
	return &BriefTag{
		Name: tag.Name,
		Slug: tag.Slug,
	}
}

func GetTagWithArticlesNumFromPO(tag model.Tag) (*TagWithArticlesNum, error) {
	articlesNum, err := model.CountArticlesByTagID(tag.ID)
	if err != nil {
		return nil, err
	}
	tagWithArticlesNum := &TagWithArticlesNum{
		Name:        tag.Name,
		Slug:        tag.Slug,
		ArticlesNum: articlesNum,
	}
	return tagWithArticlesNum, nil
}
