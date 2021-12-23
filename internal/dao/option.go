package dao

import (
	"github.com/megrez/internal/entity/po"
	"gorm.io/gorm/clause"
)

// GetOptionByKey return option by key
func (dao *DAO) GetOptionByKey(key string) (string, error) {
	option := po.Option{}
	result := dao.db.First(&option, "`key` = ?", key)
	return option.Value, result.Error
}

// SetOption handle set option
func (dao *DAO) SetOption(key, value string) error {
	option := po.Option{
		Key:   key,
		Value: value,
	}
	result := dao.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "key"}},
		DoUpdates: clause.Assignments(map[string]interface{}{"value": value}),
	}).Create(&option)
	return result.Error
}
