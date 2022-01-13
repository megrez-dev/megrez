package model

import "gorm.io/gorm"

type Journal struct {
	OriginalContent string `gorm:"type:longtext"`
	FormatContent   string `gorm:"type:longtext"`
	Images          string `gorm:"type:varchar(4095)"`
	Private         bool
	Visits          int64 `gorm:"type:int(11)"`
	Likes           int64 `gorm:"type:int(11)"`
	Status          int  `gorm:"type:int(11)"`
	gorm.Model
}

// ListAllJournals return all journals
func ListAllJournals(pageNum, pageSize int) ([]Journal, error) {
	journals := []Journal{}
	result := db.Offset(pageSize * (pageNum - 1)).Limit(pageSize).Find(&journals)
	return journals, result.Error
}

// CountAllJournals count all journal
func CountAllJournals() (int64, error) {
	var count int64
	result := db.Model(&Journal{}).Count(&count)
	return count, result.Error
}

// CreateJournal handle create link
func CreateJournal(journal *Journal) error {
	result := db.Create(journal)
	return result.Error
}
