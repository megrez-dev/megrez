package uploader

import (
	"context"
	"errors"
	"github.com/megrez/pkg/log"
	"github.com/megrez/pkg/model"
	cos "github.com/tencentyun/cos-go-sdk-v5"
	"mime/multipart"
	"net/http"
	"net/url"
	"path"
)

type TencentCosUploader struct {
	Config *TencentCosConfig
	Client *cos.Client
}

type TencentCosConfig struct {
	SecretId  string
	SecretKey string
	Domain    string
	Path      string
}

var TencentCosUploaderInstance *TencentCosUploader = nil

// GetTencentCosUploader returns a TencentCosUploader instance
func GetTencentCosUploader() (*TencentCosUploader, error) {
	secretId, err := model.GetOptionByKey(model.OptionKeyTencentCosSecretId)
	if err != nil {
		return nil, err
	}
	secretKey, err := model.GetOptionByKey(model.OptionKeyTencentCosSecretKey)
	if err != nil {
		return nil, err
	}
	domain, err := model.GetOptionByKey(model.OptionKeyTencentCosBucketDomain)
	if err != nil {
		return nil, err
	}
	src, err := model.GetOptionByKey(model.OptionKeyTencentCosBucketPath)
	if err != nil {
		return nil, err
	}
	config := &TencentCosConfig{
		SecretId:  secretId,
		SecretKey: secretKey,
		Domain:    domain,
		Path:      src,
	}
	if TencentCosUploaderInstance == nil {
		TencentCosUploaderInstance = &TencentCosUploader{
			Config: &TencentCosConfig{
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
		TencentCosUploaderInstance.Client = client
		return TencentCosUploaderInstance, nil
	} else {
		if TencentCosUploaderInstance.Config.SecretId != config.SecretId ||
			TencentCosUploaderInstance.Config.SecretKey != config.SecretKey ||
			TencentCosUploaderInstance.Config.Domain != config.Domain ||
			TencentCosUploaderInstance.Config.Path != config.Path {
			u, _ := url.Parse(config.Domain)
			b := &cos.BaseURL{BucketURL: u}
			client := cos.NewClient(b, &http.Client{
				Transport: &cos.AuthorizationTransport{
					SecretID:  config.SecretId,
					SecretKey: config.SecretKey,
				},
			})
			TencentCosUploaderInstance.Config = config
			TencentCosUploaderInstance.Client = client
		}
		return TencentCosUploaderInstance, nil
	}
}

func (t *TencentCosUploader) Upload(file *multipart.FileHeader, src string) error {
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

func (t *TencentCosUploader) Delete(src string) error {
	if t.Client == nil {
		return errors.New("COS client is nil")
	}
	_, err := t.Client.Object.Delete(context.Background(), path.Join(t.Config.Path, src))
	return err
}

func (t *TencentCosUploader) GetUrl(src string) string {
	return path.Join(t.Config.Domain, t.Config.Path, src)
}
