package model

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func New(dsn string) (*gorm.DB, error) {
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("connect db failed, ", err)
		return nil, err
	}
	err = db.AutoMigrate(
		&User{},
		&Article{},
		&Category{},
		&Comment{},
		&Journal{},
		&Link{},
		&Option{},
		&Page{},
		&Menu{},
		&Tag{},
		&ArticleTag{},
		&ArticleCategory{},
	)
	if err != nil {
		log.Println("migrate database failed, ", err)
		return nil, err
	}
	return db, nil
}
