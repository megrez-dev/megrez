package dao

import "github.com/megrez/pkg/entity/po"

// ListAllJournals return all journals
func (dao *DAO) ListAllJournals(pageNum, pageSize int) ([]po.Journal, error) {
	journals := []po.Journal{}
	result := dao.db.Offset(pageSize * (pageNum - 1)).Limit(pageSize).Find(&journals)
	return journals, result.Error
}

// CountAllJournals count all journal
func (dao *DAO) CountAllJournals() (int64, error) {
	var count int64
	result := dao.db.Model(&po.Journal{}).Count(&count)
	return count, result.Error
}

// CreateJournal handle create link
func (dao *DAO) CreateJournal(journal *po.Journal) error {
	result := dao.db.Create(&journal)
	return result.Error
}
