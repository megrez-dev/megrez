package model

import (
	"gorm.io/gorm"
	"time"
)

type ThemeOption struct {
	ID         uint      `gorm:"primarykey" json:"id"`
	ThemeID    string    `gorm:"type:varchar(255);uniqueIndex:theme_id_key" json:"themeID"`
	Type       string    `gorm:"type:varchar(255)" json:"type"`
	Key        string    `gorm:"type:varchar(255);uniqueIndex:theme_id_key" json:"key"`
	Value      string    `gorm:"type:varchar(255)" json:"value"`
	CreateTime time.Time `gorm:"default:NULL" json:"createTime"`
	UpdateTime time.Time `gorm:"default:NULL" json:"updateTime"`
}

func CreateThemeOption(tx *gorm.DB, themeOption *ThemeOption) error {
	if tx == nil {
		tx = db
	}
	if tx.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	themeOption.CreateTime = time.Now()
	result := tx.Create(themeOption)
	return result.Error
}

func UpdateThemeOption(tx *gorm.DB, themeID, key, value string) error {
	if tx == nil {
		tx = db
	}
	if tx.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	result := tx.Model(&ThemeOption{}).Where("theme_id = ? AND key = ?", themeID, key).Update("value", value)
	return result.Error
}

func DeleteThemeOptionsByID(tx *gorm.DB, id string) error {
	if tx == nil {
		tx = db
	}
	if tx.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	result := tx.Delete(&ThemeOption{}, "theme_id = ?", id)
	return result.Error
}

func ListThemeOptionsByThemeID(themeID string) ([]ThemeOption, error) {
	if db.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	var themeOptions []ThemeOption
	result := db.Where("theme_id = ?", themeID).Find(&themeOptions)
	return themeOptions, result.Error
}

func GetThemeOptionByThemeIDAndKey(themeID string, key string) (string, error) {
	if db.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	var themeOption ThemeOption
	result := db.Where("theme_id = ? AND key = ?", themeID, key).First(&themeOption)
	return themeOption.Value, result.Error
}
