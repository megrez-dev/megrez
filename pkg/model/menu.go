package model

import "gorm.io/gorm"

type Menu struct {
	Name     string `gorm:"type:varchar(255)"`
	Slug     string `gorm:"type:varchar(255)"`
	PageID   bool
	Priority uint `gorm:"type:int(11)"`
	Status   int  `gorm:"type:int(11)"`
	gorm.Model
}

// GetAuthorByID return author by id
func ListAllMenus() ([]Menu, error) {
	menus := []Menu{}
	result := db.Order("priority").Find(&menus)
	return menus, result.Error
}

// CreateMenu handle create menu
func CreateMenu(menu *Menu) error {
	result := db.Create(menu)
	return result.Error
}
