package model

import (
	"github.com/glebarez/sqlite"
	"github.com/megrez/pkg/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func NewMySQL(dsn string) (*gorm.DB, error) {
	return NewDB(mysql.Open(dsn))
}

func NewSQLite(path string) (*gorm.DB, error) {
	return NewDB(sqlite.Open(path))
}

func NewDB(dial gorm.Dialector) (*gorm.DB, error) {
	db, err = gorm.Open(dial, &gorm.Config{})
	if err != nil {
		log.Error("connect db failed, ", err)
		return nil, err
	}
	if db.Dialector.Name() == "sqlite3" {
		sqlDB, e := db.DB()
		if e != nil {
			log.Error("get underlying sql.DB failed, ", e)
			return nil, e
		}
		sqlDB.SetMaxOpenConns(1)
		if _, e = sqlDB.Exec("PRAGMA journal_mode=WAL"); e != nil {
			log.Error("set journal_mode=WAL failed, ", e)
		}
		if _, e = sqlDB.Exec("PRAGMA busy_timeout=5000"); e != nil {
			log.Error("set busy_timeout failed, ", e)
		}
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
		&ThemeOption{},
	)
	if err != nil {
		log.Error("migrate database failed, ", err)
		return nil, err
	}
	return db, nil
}

func BeginTx() *gorm.DB {
	return db.Begin()
}
