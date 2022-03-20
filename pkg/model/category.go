package model

type Category struct {
	ID          uint   `gorm:"primarykey" json:"id"`
	Name        string `gorm:"type:varchar(255)" json:"name"`
	Slug        string `gorm:"type:varchar(255);uniqueIndex" json:"slug"`
	Description string `gorm:"type:varchar(255)" json:"description"`
	Status      int    `gorm:"type:int(11)" json:"status"`
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
	var categories []Category
	result := db.Find(&categories)
	return categories, result.Error
}

// ListCategoriesByPage return categories by page
func ListCategoriesByPage(pageNum int, pageSize int) ([]Category, error) {
	var categories []Category
	result := db.Offset(pageSize * (pageNum - 1)).Limit(pageSize).Find(&categories)
	return categories, result.Error
}

// ListCategoriesByArticleID return categories by articleID
func ListCategoriesByArticleID(aid uint) ([]Category, error) {
	var categories []Category
	result := db.Where("id in (?)", db.Table("article_categories").Select("category_id").Where("article_id = ?", aid)).Find(&categories)
	return categories, result.Error
}

// CreateCategory handle create category
func CreateCategory(category *Category) error {
	result := db.Create(category)
	return result.Error
}

// UpdateCategoryByID update article by id and data
func UpdateCategoryByID(id uint, category *Category) error {
	result := db.Model(&category).Where("id= ？", id).Updates(&category)
	return result.Error
}

// DeleteCategoryByID delete article by id
func DeleteCategoryByID(id uint) error {
	result := db.Delete(&Category{}, id)
	return result.Error
}

// DeleteArticleCategoriesByArticleID delete article categories by articleID
func DeleteArticleCategoriesByArticleID(aid uint) error {
	result := db.Delete(&ArticleCategory{}, "article_id = ?", aid)
	return result.Error
}
