package dao

import (
	"github.com/megrez/pkg/entity/po"
)

// GetCategoryByID return category by id
func (dao *DAO) GetCategoryByID(id uint) (po.Category, error) {
	category := po.Category{}
	result := dao.db.First(&category, id)
	return category, result.Error
}

// GetCategoryBySlug return category by slug
func (dao *DAO) GetCategoryBySlug(slug string) (po.Category, error) {
	category := po.Category{}
	result := dao.db.First(&category, "`slug` = ?", slug)
	return category, result.Error
}

// ListAllCategories return all categories
func (dao *DAO) ListAllCategories() ([]po.Category, error) {
	categories := []po.Category{}
	result := dao.db.Find(&categories)
	return categories, result.Error
}

// CreateCategory handle create category
func (dao *DAO) CreateCategory(category *po.Category) error {
	result := dao.db.Create(category)
	return result.Error
}
