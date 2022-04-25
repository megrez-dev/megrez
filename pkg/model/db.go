package model

import (
	"github.com/megrez/pkg/log"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func NewMySQL(dsn string) (*gorm.DB, error) {
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error("connect db failed, ", err)
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
		&Attachment{},
	)
	if err != nil {
		log.Error("migrate database failed, ", err)
		return nil, err
	}
	return db, nil
}

func NewSQLite(path string) (*gorm.DB, error) {
	db, err = gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		log.Error("connect db failed, ", err)
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
		&Attachment{},
	)
	if err != nil {
		log.Error("migrate database failed, ", err)
		return nil, err
	}
	return db, nil
}
