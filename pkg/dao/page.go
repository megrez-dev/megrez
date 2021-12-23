package dao

import "github.com/megrez/pkg/entity/po"

// GetPageByID return page by pageID
func (dao *DAO) GetPageByID(id uint) (po.Page, error) {
	page := po.Page{}
	result := dao.db.First(&page, id)
	return page, result.Error
}
