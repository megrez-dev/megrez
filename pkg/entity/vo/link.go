package vo

import "github.com/megrez/pkg/model"

type Link struct {
	Name string
	Addr string
	Logo string
}

func GetLinkFromPO(link model.Link) *Link {
	return &Link{
		Name: link.Name,
		Addr: link.Addr,
		Logo: link.Logo,
	}
}
