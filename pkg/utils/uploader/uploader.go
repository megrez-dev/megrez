package uploader

import "mime/multipart"

type Uploader interface {
	Upload(file *multipart.FileHeader, path string) error
	Delete(path string) error
}
