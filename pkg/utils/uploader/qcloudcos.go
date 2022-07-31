package uploader

import (
	"context"
	"errors"
	"github.com/megrez/pkg/entity/dto"
	"github.com/megrez/pkg/log"
	"github.com/megrez/pkg/model"
	cos "github.com/tencentyun/cos-go-sdk-v5"
	"mime/multipart"
	"net/http"
	"net/url"
	"path"
)

type QcloudCosUploader struct {
	Config *QcloudCosConfig
	Client *cos.Client
}

type QcloudCosConfig struct {
	SecretId  string
	SecretKey string
	Domain    string
	Path      string
}

var QcloudCosUploaderInstance *QcloudCosUploader = nil

// GetQcloudCosUploader returns a QcloudCosUploader instance
func GetQcloudCosUploader() (*QcloudCosUploader, error) {
	secretId, err := model.GetOptionByKey(model.OptionKeyQcloudCosSecretId)
	if err != nil {
		return nil, err
	}
	secretKey, err := model.GetOptionByKey(model.OptionKeyQcloudCosSecretKey)
	if err != nil {
		return nil, err
	}
	domain, err := model.GetOptionByKey(model.OptionKeyQcloudCosDomain)
	if err != nil {
		return nil, err
	}
	src, err := model.GetOptionByKey(model.OptionKeyQcloudCosPath)
	if err != nil {
		return nil, err
	}
	config := &QcloudCosConfig{
		SecretId:  secretId,
		SecretKey: secretKey,
		Domain:    domain,
		Path:      src,
	}
	if QcloudCosUploaderInstance == nil {
		QcloudCosUploaderInstance = &QcloudCosUploader{
			Config: &QcloudCosConfig{
				SecretId:  secretId,
				SecretKey: secretKey,
				Domain:    domain,
				Path:      src,
			},
		}
		u, _ := url.Parse(config.Domain)
		b := &cos.BaseURL{BucketURL: u}
		client := cos.NewClient(b, &http.Client{
			Transport: &cos.AuthorizationTransport{
				SecretID:  config.SecretId,
				SecretKey: config.SecretKey,
			},
		})
		QcloudCosUploaderInstance.Client = client
		return QcloudCosUploaderInstance, nil
	} else {
		if QcloudCosUploaderInstance.Config.SecretId != config.SecretId ||
			QcloudCosUploaderInstance.Config.SecretKey != config.SecretKey ||
			QcloudCosUploaderInstance.Config.Domain != config.Domain ||
			QcloudCosUploaderInstance.Config.Path != config.Path {
			u, _ := url.Parse(config.Domain)
			b := &cos.BaseURL{BucketURL: u}
			client := cos.NewClient(b, &http.Client{
				Transport: &cos.AuthorizationTransport{
					SecretID:  config.SecretId,
					SecretKey: config.SecretKey,
				},
			})
			QcloudCosUploaderInstance.Config = config
			QcloudCosUploaderInstance.Client = client
		}
		return QcloudCosUploaderInstance, nil
	}
}

func GetTempQcloudCosUploader(setting dto.QCloudCosSetting) (*QcloudCosUploader, error) {
	tempQcloudCosUploader := &QcloudCosUploader{
		Config: &QcloudCosConfig{
			SecretId:  setting.SecretID,
			SecretKey: setting.SecretKey,
			Domain:    setting.Domain,
		},
	}
	u, _ := url.Parse(setting.Domain)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  setting.SecretID,
			SecretKey: setting.SecretKey,
		},
	})
	tempQcloudCosUploader.Client = client
	return tempQcloudCosUploader, nil
}

func (t *QcloudCosUploader) Ping() error {
	exist, err := t.Client.Bucket.IsExist(context.Background())
	if err != nil {
		return errors.New("连接失败，请检查你的配置")
	}
	if !exist {
		return errors.New("Bucket 不存在或已被删除")
	}
	return nil
}

func (t *QcloudCosUploader) Upload(file *multipart.FileHeader, src string) error {
	if t.Client == nil {
		return errors.New("COS client is nil")
	}
	opened, err := file.Open()
	if err != nil {
		return err
	}
	defer opened.Close()
	log.Debug("Uploading to COS... [%s]", path.Join(t.Config.Path, src))
	_, err = t.Client.Object.Put(context.Background(), path.Join(t.Config.Path, src), opened, nil)
	return err
}

func (t *QcloudCosUploader) Delete(src string) error {
	if t.Client == nil {
		return errors.New("COS client is nil")
	}
	_, err := t.Client.Object.Delete(context.Background(), path.Join(t.Config.Path, src))
	return err
}

func (t *QcloudCosUploader) GetUrl(src string) string {
	return path.Join(t.Config.Domain, t.Config.Path, src)
}
