package vo

import "github.com/megrez/pkg/model"

type BriefCategory struct {
	Name string
	Slug string
}

func GetBriefCategoryFromPO(category model.Category) *BriefCategory {
	return &BriefCategory{
		Name: category.Name,
		Slug: category.Slug,
	}
}
