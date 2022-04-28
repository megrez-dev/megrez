package model

import "time"

type ThemeOption struct {
	ID         uint      `gorm:"primarykey" json:"id"`
	ThemeID    string    `gorm:"type:varchar(255)" json:"themeID"`
	Key        string    `gorm:"type:varchar(255)" json:"key"`
	Value      string    `gorm:"type:varchar(255)" json:"value"`
	CreateTime time.Time `gorm:"default:NULL" json:"createTime"`
	UpdateTime time.Time `gorm:"default:NULL" json:"updateTime"`
}

func ListThemeOptionsByThemeID(themeID string) ([]ThemeOption, error) {
	var themeOptions []ThemeOption
	result := db.Where("theme_id = ?", themeID).Find(&themeOptions)
	return themeOptions, result.Error
}
