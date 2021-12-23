package dao

import (
	"fmt"
	"log"

	"github.com/megrez/pkg/entity/po"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var instance *DAO

type DAO struct {
	db *gorm.DB
}

func New(dsn string) (*DAO, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("connect db failed, ", err)
		return nil, err
	}
	err = db.AutoMigrate(
		&po.Author{},
		&po.Article{},
		&po.Category{},
		&po.Comment{},
		&po.Journal{},
		&po.Link{},
		&po.Option{},
		&po.Page{},
		&po.Menu{},
		&po.Tag{})
	if err != nil {
		log.Println("migrate database failed, ", err)
		return nil, err
	}
	instance = &DAO{db}
	return instance, nil
}

func GetDAO() (*DAO, error) {
	if instance == nil || instance.db == nil {
		return nil, fmt.Errorf("DAO have not init")
	} else {
		return instance, nil
	}
}
