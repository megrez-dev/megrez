package model

import "gorm.io/gorm"

type Category struct {
	Name        string `gorm:"type:varchar(255)"`
	Slug        string `gorm:"type:varchar(255);uniqueIndex"`
	Description string `gorm:"type:varchar(255)"`
	Status      int `gorm:"type:int(11)"`
	gorm.Model
}

// GetCategoryByID return category by id
func GetCategoryByID(id uint) (Category, error) {
	category := Category{}
	result := db.First(&category, id)
	return category, result.Error
}

// GetCategoryBySlug return category by slug
func GetCategoryBySlug(slug string) (Category, error) {
	category := Category{}
	result := db.First(&category, "`slug` = ?", slug)
	return category, result.Error
}

// ListAllCategories return all categories
func ListAllCategories() ([]Category, error) {
	categories := []Category{}
	result := db.Find(&categories)
	return categories, result.Error
}

// CreateCategory handle create category
func CreateCategory(category *Category) error {
	result := db.Create(category)
	return result.Error
}
