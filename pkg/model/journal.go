package model

import (
	"gorm.io/gorm"
	"time"
)

type Journal struct {
	ID         uint      `gorm:"primarykey" json:"id"`
	Content    string    `gorm:"type:longtext" json:"content"`
	Images     string    `gorm:"type:varchar(4095)" json:"images"`
	Private    bool      `json:"private"`
	Visits     int64     `gorm:"type:int(11)" json:"visits"`
	Likes      int64     `gorm:"type:int(11)" json:"likes"`
	Status     int       `gorm:"type:int(11)" json:"status"`
	CreateTime time.Time `gorm:"default:NULL" json:"createTime"`
	UpdateTime time.Time `gorm:"default:NULL" json:"updateTime"`
}

// ListAllJournals return all journals
func ListAllJournals(pageNum, pageSize int) ([]Journal, error) {
	if db.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	var journals []Journal
	result := db.Order("create_time DESC").Offset(pageSize * (pageNum - 1)).Limit(pageSize).Find(&journals)
	return journals, result.Error
}

// CountAllJournals count all journal
func CountAllJournals() (int64, error) {
	if db.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	var count int64
	result := db.Model(&Journal{}).Count(&count)
	return count, result.Error
}

// CreateJournal handle create link
func CreateJournal(tx *gorm.DB, journal *Journal) error {
	if tx == nil {
		tx = db
	}
	if tx.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	result := tx.Create(journal)
	return result.Error
}
