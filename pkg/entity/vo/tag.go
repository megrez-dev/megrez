package vo

import (
	"github.com/megrez/pkg/entity/po"
)

type BriefTag struct {
	Name string
	Slug string
}

func GetBriefTagFromPO(po po.Tag) *BriefTag {
	return &BriefTag{
		Name: po.Name,
		Slug: po.Slug,
	}
}
