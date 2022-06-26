package model

import (
	"gorm.io/gorm"
)

var (
	OptionKeyIsInstalled            = "is_installed"
	OptionKeyBlogURL                = "blog_url"
	OptionKeyBlogTitle              = "blog_title"
	OptionKeyBlogDescription        = "blog_description"
	OptionKeyBlogBirth              = "blog_birth"
	OptionKeyBlogTheme              = "blog_theme"
	OptionKeyUploadType             = "upload_type"
	OptionKeyTencentCosSecretId     = "tencent_cos_secret_id"
	OptionKeyTencentCosSecretKey    = "tencent_cos_secret_key"
	OptionKeyTencentCosBucketDomain = "tencent_cos_bucket_domain"
	OptionKeyTencentCosBucketPath   = "tencent_cos_bucket_path"
	OptionKeySMTPURL                = "smtp_url"
	OptionKeySMTPProtocol           = "smtp_protocol"
	OptionKeySMTPSSLPort            = "smtp_ssl_port"
	OptionKeySMTPEmail              = "smtp_email"
	OptionKeySMTPPassword           = "smtp_password"
	OptionKeyGithub                 = "github"
	OptionKeyEmail                  = "email"
	OptionKeyCommentsPageSize       = "comments_page_size"
)

type Option struct {
	ID    uint   `gorm:"primarykey" json:"id"`
	Key   string `gorm:"type:varchar(255);uniqueIndex" json:"key"`
	Value string `gorm:"type:varchar(255)" json:"value"`
}

// GetOptionByKey return option by key
func GetOptionByKey(key string) (string, error) {
	if db.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	option := Option{}
	result := db.First(&option, "`key` = ?", key)
	return option.Value, result.Error
}

// SetOption handle set option
func SetOption(tx *gorm.DB, key, value string) error {
	if tx == nil {
		tx = db
	}
	if tx.Dialector.Name() == "sqlite3" {
		lock.Lock()
		defer lock.Unlock()
	}
	option := &Option{
		Key:   key,
		Value: value,
	}
	result := db.First(&option, "`key` = ?", key)
	if result.Error == gorm.ErrRecordNotFound {
		result = tx.Create(option)
		return result.Error
	} else if result.Error == nil {
		option.Value = value
		result = tx.Save(option)
		return result.Error
	} else {
		return result.Error
	}
}

//// SetOption handle set option
//func SetOption(tx *gorm.DB, key, value string) error {
//	if tx == nil {
//		tx = db
//	}
//	if tx.Dialector.Name() == "sqlite3" {
//		lock.Lock()
//		defer lock.Unlock()
//	}
//	option := Option{
//		Key:   key,
//		Value: value,
//	}
//	result := tx.Clauses(clause.OnConflict{
//		Columns:   []clause.Column{{Name: "key"}},
//		DoUpdates: clause.Assignments(map[string]interface{}{"value": value}),
//	}).Create(&option)
//	return result.Error
//}
