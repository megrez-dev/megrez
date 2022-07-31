package dto

import (
	"github.com/megrez/pkg/model"
	"gorm.io/gorm"
)

type Setting struct {
	Basic      BasicSetting      `json:"basic"`
	Attachment AttachmentSetting `json:"attachment"`
}

type BasicSetting struct {
	BlogTitle       string `json:"blogTitle"`
	BlogDescription string `json:"blogDescription"`
	BlogURL         string `json:"blogURL"`
	BlogFavicon     string `json:"blogFavicon"`
}

type AttachmentSetting struct {
	UploadType string           `json:"uploadType"`
	QCloudCos  QCloudCosSetting `json:"qCloudCos"`
}

type QCloudCosSetting struct {
	SecretID  string `json:"secretID"`
	SecretKey string `json:"secretKey"`
	Path      string `json:"path"`
	Domain    string `json:"domain"`
}

func (s *Setting) LoadFromModel() error {
	title, err := model.GetOptionByKey(model.OptionKeyBlogTitle)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			title = ""
		} else {
			return err
		}
	}
	s.Basic.BlogTitle = title
	description, err := model.GetOptionByKey(model.OptionKeyBlogDescription)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			description = ""
		} else {
			return err
		}
	}
	s.Basic.BlogDescription = description
	url, err := model.GetOptionByKey(model.OptionKeyBlogURL)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			url = ""
		} else {
			return err
		}
	}
	s.Basic.BlogURL = url
	favicon, err := model.GetOptionByKey(model.OptionKeyBlogFavicon)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			favicon = ""
		} else {
			return err
		}
	}
	s.Basic.BlogFavicon = favicon
	uploadType, err := model.GetOptionByKey(model.OptionKeyUploadType)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			uploadType = model.AttachmentTypeLocal
		} else {
			return err
		}
	}
	s.Attachment.UploadType = uploadType
	qCloudCosSecretID, err := model.GetOptionByKey(model.OptionKeyQcloudCosSecretId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			qCloudCosSecretID = ""
		} else {
			return err
		}
	}
	s.Attachment.QCloudCos.SecretID = qCloudCosSecretID
	qCloudCosSecretKey, err := model.GetOptionByKey(model.OptionKeyQcloudCosSecretKey)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			qCloudCosSecretKey = ""
		} else {
			return err
		}
	}
	s.Attachment.QCloudCos.SecretKey = qCloudCosSecretKey
	qCloudCosPath, err := model.GetOptionByKey(model.OptionKeyQcloudCosPath)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			qCloudCosPath = ""
		} else {
			return err
		}
	}
	s.Attachment.QCloudCos.Path = qCloudCosPath
	qCloudCosBucketDomain, err := model.GetOptionByKey(model.OptionKeyQcloudCosDomain)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			qCloudCosBucketDomain = ""
		} else {
			return err
		}
	}
	s.Attachment.QCloudCos.Domain = qCloudCosBucketDomain
	return nil
}

func (s *Setting) SaveToModel() error {
	tx := model.BeginTx()
	err := model.SetOption(nil, model.OptionKeyBlogTitle, s.Basic.BlogTitle)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = model.SetOption(nil, model.OptionKeyBlogDescription, s.Basic.BlogDescription)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = model.SetOption(nil, model.OptionKeyBlogURL, s.Basic.BlogURL)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = model.SetOption(nil, model.OptionKeyBlogFavicon, s.Basic.BlogFavicon)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = model.SetOption(nil, model.OptionKeyUploadType, s.Attachment.UploadType)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = model.SetOption(nil, model.OptionKeyQcloudCosSecretId, s.Attachment.QCloudCos.SecretID)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = model.SetOption(nil, model.OptionKeyQcloudCosSecretKey, s.Attachment.QCloudCos.SecretKey)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = model.SetOption(nil, model.OptionKeyQcloudCosPath, s.Attachment.QCloudCos.Path)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = model.SetOption(nil, model.OptionKeyQcloudCosDomain, s.Attachment.QCloudCos.Domain)
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
