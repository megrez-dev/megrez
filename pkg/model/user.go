package model

type User struct {
	ID          uint   `gorm:"primarykey" json:"id"`
	Username    string `gorm:"type:varchar(255);unique_index" json:"username"`
	Nickname    string `gorm:"type:varchar(255)" json:"nickname"`
	Avatar      string `gorm:"type:varchar(255)" json:"avatar"`
	Description string `gorm:"type:varchar(255)" json:"description"`
	Email       string `gorm:"type:varchar(255);unique_index" json:"email"`
	Password    string `gorm:"type:varchar(255)" json:"password"`
	Status      int    `gorm:"type:int(11)" json:"status"`
}

// GetUserByUsername return user by username or password
func GetUserByUsername(username string) (User, error) {
	user := User{}
	result := db.Where("username = ?", username).Or("email = ?", username).First(&user)
	return user, result.Error
}

// CreateUser handle create user
func CreateUser(user *User) error {
	result := db.Create(user)
	return result.Error
}
