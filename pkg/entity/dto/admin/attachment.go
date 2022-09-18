package admin

import (
	"github.com/megrez/pkg/model"
	fileutils "github.com/megrez/pkg/utils/file"
)

type AttachmentDTO struct {
	ID         uint   `json:"id"`
	FileName   string `json:"fileName"`
	URL        string `json:"url"`
	ThumbURL   string `json:"thumbURL"`
	Ext        string `json:"ext"`
	Size       string `json:"size"`
	Width      int    `json:"width"`
	Height     int    `json:"height"`
	Type       string `json:"type"`
	UploadTime string `json:"uploadTime"`
}

func (dto *AttachmentDTO) LoadFromModel(attachment model.Attachment) error {
	dto.ID = attachment.ID
	dto.FileName = attachment.FileName
	dto.URL = attachment.URL
	dto.ThumbURL = attachment.ThumbURL
	dto.Ext = attachment.Ext
	dto.Width = attachment.Width
	dto.Height = attachment.Height
	dto.Type = attachment.Type
	dto.UploadTime = attachment.UploadTime.Format("2006-01-02 15:04:05")
	dto.Size = fileutils.FormatSize(attachment.Size)
	return nil
}
