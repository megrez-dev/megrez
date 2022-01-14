package model

import "gorm.io/gorm"

type Journal struct {
	OriginalContent string `gorm:"type:longtext" json:"originalContent"`
	FormatContent   string `gorm:"type:longtext" json:"formatContent"`
	Images          string `gorm:"type:varchar(4095)" json:"images"`
	Private         bool `json:"private"`
	Visits          int64 `gorm:"type:int(11)" json:"visits"`
	Likes           int64 `gorm:"type:int(11)" json:"likes"`
	Status          int  `gorm:"type:int(11)" json:"status"`
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
