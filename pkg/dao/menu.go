package dao

import (
	"github.com/megrez/pkg/entity/po"
)

// GetAuthorByID return author by id
func (dao *DAO) ListAllMenus() ([]po.Menu, error) {
	menus := []po.Menu{}
	result := dao.db.Order("priority").Find(&menus)
	return menus, result.Error
}

// CreateMenu handle create menu
func (dao *DAO) CreateMenu(menu *po.Menu) error {
	result := dao.db.Create(&menu)
	return result.Error
}
