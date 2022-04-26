package model

import "time"

type Link struct {
	ID         uint      `gorm:"primarykey" json:"id"`
	Name       string    `gorm:"type:varchar(255)" json:"name"`
	URL        string    `gorm:"type:varchar(255)" json:"url"`
	Logo       string    `gorm:"type:varchar(255)" json:"logo"`
	Priority   uint      `gorm:"type:int(11)" json:"priority"`
	Status     int       `gorm:"type:int(11)" json:"status"`
	CreateTime time.Time `gorm:"default:NULL" json:"createTime"`
	UpdateTime time.Time `gorm:"default:NULL" json:"updateTime"`
}

// ListAllLinks return all links
func ListAllLinks() ([]Link, error) {
	if db.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	var links []Link
	result := db.Order("priority").Find(&links)
	return links, result.Error
}

// ListLinksByPage return links by page
func ListLinksByPage(page, pageSize int) ([]Link, error) {
	if db.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	var links []Link
	result := db.Order("priority").Offset((page - 1) * pageSize).Limit(pageSize).Find(&links)
	return links, result.Error
}

func CountLinks() (int64, error) {
	if db.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	var count int64
	result := db.Model(&Link{}).Count(&count)
	return count, result.Error
}

// CreateLink handle create link
func CreateLink(link *Link) error {
	if db.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	result := db.Create(link)
	return result.Error
}

func UpdateLink(link *Link) error {
	if db.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	result := db.Model(&Link{}).Where("id = ?", link.ID).Updates(link)
	return result.Error
}

func DeleteLinkByID(id uint) error {
	if db.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	result := db.Delete(&Link{}, "id = ?", id)
	return result.Error
}
