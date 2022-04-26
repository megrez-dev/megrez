package model

type Menu struct {
	ID       uint   `gorm:"primarykey" json:"id"`
	Name     string `gorm:"type:varchar(255)" json:"name"`
	Slug     string `gorm:"type:varchar(255)" json:"slug"`
	PageID   uint   `gorm:"type:int(11)" json:"pageID"`
	Priority uint   `gorm:"type:int(11)" json:"priority"`
	Status   int    `gorm:"type:int(11)" json:"status"`
}

// ListAllMenus list all menus
func ListAllMenus() ([]Menu, error) {
	if db.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	var menus []Menu
	result := db.Order("priority").Find(&menus)
	return menus, result.Error
}

// CreateMenu handle create menu
func CreateMenu(menu *Menu) error {
	if db.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	result := db.Create(menu)
	return result.Error
}
