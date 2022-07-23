package bolt

import (
	"errors"
	bolt "go.etcd.io/bbolt"
)

const BucketName = "Megrez"
const BoltKeyDatabaseType = "db_type"
const BoltKeyMySQLHost = "mysql_host"
const BoltKeyMySQLPort = "mysql_port"
const BoltKeyMySQLDatabase = "mysql_database"
const BoltKeyMySQLUser = "mysql_user"
const BoltKeyMySQLPassword = "mysql_password"
const BoltKeySQLitePath = "sqlite_path"

var db *bolt.DB

func NewBolt(path string) (*bolt.DB, error) {
	var err error
	db, err = bolt.Open(path, 0666, nil)
	if err != nil {
		return nil, err
	}
	err = db.Update(func(tx *bolt.Tx) error {
		var cErr error
		_, cErr = tx.CreateBucketIfNotExists([]byte(BucketName))
		return cErr
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func Set(key, value string) error {
	if db == nil {
		return errors.New("bolt bd not init")
	}
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BucketName))
		if b == nil {
			return errors.New("bucket not exists")
		}
		err := b.Put([]byte(key), []byte(value))
		return err
	})
}

func Get(key string) (value string, err error) {
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BucketName))
		if b == nil {
			return errors.New("bucket not exists")
		}
		value = string(b.Get([]byte(key)))
		return nil
	})
	return value, err
}
