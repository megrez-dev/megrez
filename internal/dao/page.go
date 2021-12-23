package dao

import "github.com/megrez/internal/entity/po"

// GetPageByID return page by pageID
func (dao *DAO) GetPageByID(id int) (po.Page, error) {
	page := po.Page{}
	result := dao.db.First(&page, id)
	return page, result.Error
}
