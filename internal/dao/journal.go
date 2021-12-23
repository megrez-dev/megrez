package dao

import "github.com/megrez/internal/entity/po"

// ListAllJournals return all journals
func (dao *DAO) ListAllJournals(pageNum, pageSize int) ([]po.Journal, error) {
	journals := []po.Journal{}
	result := dao.db.Offset(pageSize * (pageNum - 1)).Limit(pageSize).Find(&journals)
	return journals, result.Error
}

// CreateJournal handle create link
func (dao *DAO) CreateJournal(journal *po.Journal) error {
	result := dao.db.Create(&journal)
	return result.Error
}
