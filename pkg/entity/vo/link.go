package vo

import "github.com/megrez/pkg/entity/po"

type Link struct {
	Name string
	Addr string
	Logo string
}

func GetLinkFromPO(po po.Link) *Link {
	link := &Link{
		Name: po.Name,
		Addr: po.Addr,
		Logo: po.Logo,
	}
	return link
}
