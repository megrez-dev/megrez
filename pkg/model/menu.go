package model

import "gorm.io/gorm"

type Menu struct {
	Name     string `gorm:"type:varchar(255)" json:"name"`
	Slug     string `gorm:"type:varchar(255)" json:"slug"`
	PageID   bool `json:"pageID"`
	Priority uint `gorm:"type:int(11)" json:"priority"`
	Status   int  `gorm:"type:int(11)" json:"status"`
	gorm.Model
}

// ListAllMenus list all menus
func ListAllMenus() ([]Menu, error) {
	var menus []Menu
	result := db.Order("priority").Find(&menus)
	return menus, result.Error
}

// CreateMenu handle create menu
func CreateMenu(menu *Menu) error {
	result := db.Create(menu)
	return result.Error
}
