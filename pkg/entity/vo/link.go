package vo

import "github.com/megrez/pkg/model"

type Link struct {
	Name string
	URL  string
	Logo string
}

func GetLinkFromPO(link model.Link) *Link {
	return &Link{
		Name: link.Name,
		URL:  link.URL,
		Logo: link.Logo,
	}
}
