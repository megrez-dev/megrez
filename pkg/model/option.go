package model

import (
	"gorm.io/gorm/clause"
)

type Option struct {
	ID    uint   `gorm:"primarykey" json:"id"`
	Key   string `gorm:"type:varchar(255);uniqueIndex" json:"key"`
	Value string `gorm:"type:varchar(255)" json:"value"`
}

// GetOptionByKey return option by key
func GetOptionByKey(key string) (string, error) {
	option := Option{}
	result := db.First(&option, "`key` = ?", key)
	return option.Value, result.Error
}

// SetOption handle set option
func SetOption(key, value string) error {
	option := Option{
		Key:   key,
		Value: value,
	}
	result := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "key"}},
		DoUpdates: clause.Assignments(map[string]interface{}{"value": value}),
	}).Create(&option)
	return result.Error
}
