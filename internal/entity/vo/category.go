package vo

import (
	"github.com/megrez/internal/entity/po"
)

type BriefCategory struct {
	Name string
	Slug string
}

func GetBriefCategoryFromPO(po *po.Category) *BriefCategory {
	return &BriefCategory{
		Name: po.Name,
		Slug: po.Slug,
	}
}
