package admin

import (
	"github.com/google/uuid"
	"github.com/megrez/pkg/entity/dto"
	"github.com/megrez/pkg/model"
	dirUtils "github.com/megrez/pkg/utils/dir"
	"github.com/megrez/pkg/utils/uploader"
	"gorm.io/gorm"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"net/http"
	"path"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/megrez/pkg/log"
	"github.com/megrez/pkg/utils/errmsg"
)

// UploadAttachment godoc
// @Summary upload attachment
// @Schemes http https
// @Description upload attachment
// @Accept application/json, text/plain, */*
// @Param Authorization header string false "Authorization"
// @Param  file formData file true "file"
// @Success 200 {object} errmsg.Response{data=model.Attachment}
// @Router /api/admin/upload [post]
func UploadAttachment(c *gin.Context) {
	file, err := c.FormFile("file")
	log.Debug("file:", file.Filename)
	if err != nil {
		log.Error("get file from request failed: ", err)
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	fileName := file.Filename
	ext := path.Ext(fileName)
	newName := uuid.NewString() + ext
	open, err := file.Open()
	if err != nil {
		log.Error("open file failed: ", err)
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	defer open.Close()
	config, _, err := image.DecodeConfig(open)
	if err != nil {
		log.Error("decode image error: ", err)
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	width := config.Width
	height := config.Height
	size := file.Size
	dayDir := path.Join(strconv.Itoa(time.Now().Year()),
		strconv.Itoa(int(time.Now().Month())),
		strconv.Itoa(time.Now().Day()))
	if err != nil {
		log.Error("get upload dir error: ", err)
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	var url string
	var tp string
	// TODO: make thumbnail
	uploadType, err := model.GetOptionByKey(model.OptionKeyUploadType)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			uploadType = "local"
		} else {
			log.Error("get upload type error: ", err)
			c.JSON(http.StatusOK, errmsg.Error())
			return
		}
	}
	switch uploadType {
	case model.AttachmentTypeLocal:
		uploadHome, err := dirUtils.GetOrCreateUploadHome()
		if err != nil {
			log.Error("get upload dir error: ", err)
			c.JSON(http.StatusOK, errmsg.Error())
			return
		}
		err = dirUtils.CreateDir(uploadHome, dayDir)
		if err != nil {
			log.Error("create upload dir error: ", err)
			c.JSON(http.StatusOK, errmsg.Error())
			return
		}
		url = "/upload/" + path.Join(dayDir, newName)
		tp = model.AttachmentTypeLocal
		if err := c.SaveUploadedFile(file, path.Join(uploadHome, dayDir, newName)); err != nil {
			log.Error("upload file error: ", err)
			c.JSON(http.StatusOK, errmsg.Error())
			return
		}
		log.Infof("upload file success, type: [%s], name: %s", "local", newName)
	case model.AttachmentTypeQcloudCos:
		cos, err := uploader.GetQcloudCosUploader()
		if err != nil {
			log.Error("get cos uploader error: ", err)
			c.JSON(http.StatusOK, errmsg.Error())
			return
		}
		err = cos.Upload(file, path.Join(dayDir, newName))
		if err != nil {
			log.Error("upload file error: ", err)
			c.JSON(http.StatusOK, errmsg.Error())
			return
		}
		url = cos.GetUrl(path.Join(dayDir, newName))
		tp = model.AttachmentTypeQcloudCos
	case model.AttachmentTypeAliyunOss:
		// TODO
	case model.AttachmentTypeHuaweiObs:
		// TODO
	case model.AttachmentTypeQiniuyun:
		// TODO
	case model.AttachmentTypeYoupaiyun:
		// TODO
	default:
		uploadHome, err := dirUtils.GetOrCreateUploadHome()
		if err != nil {
			log.Error("get upload dir error: ", err)
			c.JSON(http.StatusOK, errmsg.Error())
			return
		}
		err = dirUtils.CreateDir(uploadHome, dayDir)
		if err != nil {
			log.Error("create upload dir error: ", err)
			c.JSON(http.StatusOK, errmsg.Error())
			return
		}
		url = "/upload/" + path.Join(dayDir, newName)
		tp = model.AttachmentTypeLocal
		if err := c.SaveUploadedFile(file, path.Join(uploadHome, dayDir, newName)); err != nil {
			log.Error("upload file error: ", err)
			c.JSON(http.StatusOK, errmsg.Error())
			return
		}
		log.Infof("upload file success, type: [%s], name: %s", "local", newName)
	}
	attachment := &model.Attachment{
		FileName:   fileName,
		URL:        url,
		ThumbURL:   url,
		Size:       size,
		Ext:        ext,
		Height:     height,
		Width:      width,
		Type:       tp,
		UploadTime: time.Now(),
	}
	tx := model.BeginTx()
	err = model.CreateAttachment(tx, attachment)
	if err != nil {
		log.Error("create attachment error: ", err)
		tx.Rollback()
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	tx.Commit()
	c.JSON(http.StatusOK, errmsg.Success(attachment))
}

// ListAttachments godoc
// @Summary list attachments
// @Schemes http https
// @Description list attachments
// @Accept application/json
// @Param Authorization header string false "Authorization"
// @Param pageNum query int false "page num"
// @Param pageSize query int false "page size"
// @Success 200 {object} errmsg.Response{data=dto.Pagination{list=[]dto.AttachmentDTO}}
// @Router /api/admin/attachments [get]
func ListAttachments(c *gin.Context) {
	var pageNum, pageSize int
	var err error
	if c.Query("pageNum") == "" {
		pageNum = 1
	} else {
		pageNum, err = strconv.Atoi(c.Query("pageNum"))
		if err != nil {
			log.Error(err)
			c.JSON(http.StatusOK, errmsg.Fail(errmsg.ErrorInvalidParam))
			return
		}
	}
	if c.Query("pageSize") == "" {
		pageSize = 10
	} else {
		pageSize, err = strconv.Atoi(c.Query("pageSize"))
		if err != nil {
			log.Error(err)
			c.JSON(http.StatusOK, errmsg.Fail(errmsg.ErrorInvalidParam))
			return
		}
	}
	attachments, err := model.ListAttachments(pageNum, pageSize)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	var attachmentDTOs []dto.AttachmentDTO
	for _, attachment := range attachments {
		attachmentDTO := dto.AttachmentDTO{}
		err := attachmentDTO.LoadFromModel(attachment)
		if err != nil {
			log.Error(err)
			c.JSON(http.StatusOK, errmsg.Error())
			return
		}
		attachmentDTOs = append(attachmentDTOs, attachmentDTO)
	}
	total, err := model.CountAllAttachments()
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	pagination := dto.Pagination{
		List:     attachmentDTOs,
		Current:  pageNum,
		PageSize: pageSize,
		Total:    total,
	}
	c.JSON(http.StatusOK, errmsg.Success(pagination))
}

func PingQcloudCos(c *gin.Context) {
	settings := &dto.QCloudCosSetting{}
	err := c.ShouldBindJSON(settings)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, errmsg.Fail(errmsg.ErrorInvalidParam))
		return
	}
	cos, err := uploader.GetTempQcloudCosUploader(*settings)
	if err != nil {
		log.Error("get cos uploader error: ", err)
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	err = cos.Ping()
	if err != nil {
		log.Error("ping cos error: ", err)
		c.JSON(http.StatusOK, errmsg.FailMsg(err.Error()))
		return
	}
	c.JSON(http.StatusOK, errmsg.Success("连接成功"))
}
