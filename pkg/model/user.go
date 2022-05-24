package model

import "gorm.io/gorm"

type User struct {
	ID          uint   `gorm:"primarykey" json:"id"`
	Username    string `gorm:"type:varchar(255);uniqueIndex" json:"username"`
	Nickname    string `gorm:"type:varchar(255)" json:"nickname"`
	Avatar      string `gorm:"type:varchar(255)" json:"avatar"`
	Description string `gorm:"type:varchar(255)" json:"description"`
	Email       string `gorm:"type:varchar(255);uniqueIndex" json:"email"`
	Password    string `gorm:"type:varchar(255)" json:"password"`
	Status      int    `gorm:"type:int(11)" json:"status"`
}

// GetUserByUsername return user by username or password
func GetUserByUsername(username string) (User, error) {
	if db.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	user := User{}
	result := db.Where("username = ?", username).Or("email = ?", username).First(&user)
	return user, result.Error
}

// CreateUser handle create user
func CreateUser(tx *gorm.DB, user *User) error {
	if tx == nil {
		tx = db
	}
	if tx.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	result := tx.Create(user)
	return result.Error
}
