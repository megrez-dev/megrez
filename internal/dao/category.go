package dao

import (
	"github.com/megrez/internal/entity/po"
)

// GetCategoryByID return category by id
func (dao *DAO) GetCategoryByID(id uint) (po.Category, error) {
	category := po.Category{}
	result := dao.db.First(&category, id)
	return category, result.Error
}

// CreateCategory handle create category
func (dao *DAO) CreateCategory(category *po.Category) error {
	result := dao.db.Create(&category)
	return result.Error
}
