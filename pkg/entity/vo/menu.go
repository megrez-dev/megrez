package vo

import "github.com/megrez/pkg/model"

type Menu struct {
	Name string
	Slug string
}

func GetMenuFromPO(menu model.Menu) *Menu {
	return &Menu{
		Name: menu.Name,
		Slug: menu.Slug,
	}
}
