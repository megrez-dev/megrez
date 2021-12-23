package po

import "gorm.io/gorm"

type Comment struct {
	ArticleID uint   `gorm:"type:int(11)"`
	PageID    uint   `gorm:"type:int(11)"`
	Content   string `gorm:"type:longtext"`
	RootID    uint   `gorm:"type:int(11)"`
	ParentID  uint   `gorm:"type:int(11)"`
	AuthorID  uint   `gorm:"type:int(11)"`
	Type      int    `gorm:"type:int(11)"`
	Status    int    `gorm:"type:int(11)"`
	gorm.Model
}
