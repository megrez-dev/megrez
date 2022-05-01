package admin

import (
	"github.com/google/uuid"
	"github.com/megrez/pkg/entity/dto"
	"github.com/megrez/pkg/entity/vo"
	"github.com/megrez/pkg/model"
	dirUtils "github.com/megrez/pkg/utils/dir"
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

func UploadAttachment(c *gin.Context) {
	file, err := c.FormFile("file")
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
	var url string
	// TODO: make thumbnail
	uploadType, err := model.GetOptionByKey(vo.OptionKeyUploadType)
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
	case "local":
		uploadHome, err := dirUtils.GetOrCreateUploadHome()
		dayDir := path.Join(strconv.Itoa(time.Now().Year()),
			strconv.Itoa(int(time.Now().Month())),
			strconv.Itoa(time.Now().Day()))
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
		if err := c.SaveUploadedFile(file, path.Join(uploadHome, dayDir, newName)); err != nil {
			log.Error("upload file error: ", err)
			c.JSON(http.StatusOK, errmsg.Error())
			return
		}
		log.Infof("upload file success, type: [%s], name: %s", "local", newName)
	case "qiniu":
		// TODO
	case "ali_oss":
		// TODO
	case "qcloud_cos":
		// TODO
	case "huawei_obs":
		// TODO
	case "youpai":
		// TODO
	default:
		uploadDir, err := dirUtils.GetOrCreateUploadHome()
		if err != nil {
			log.Error("get upload dir error: ", err)
			c.JSON(http.StatusOK, errmsg.Error())
			return
		}
		url = "/upload/" + newName
		if err := c.SaveUploadedFile(file, path.Join(uploadDir, newName)); err != nil {
			log.Error("upload file error: ", err)
			c.JSON(http.StatusOK, errmsg.Error())
			return
		}
		log.Infof("upload file success, type: [%s], name: %s", "local", newName)
	}
	attahcment := &model.Attachment{
		FileName: fileName,
		URL:      url,
		ThumbURL: url,
		Size:     size,
		Ext:      ext,
		Height:   height,
		Width:    width,
		// TODO: 0: local, 1: qiniu, 2: ali_oss, 3: qcloud_cos, 4: huawei_obs, 5: youpai
		Type:       0,
		UploadTime: time.Now(),
	}
	err = model.CreateAttachment(attahcment)
	if err != nil {
		log.Error("create attachment error: ", err)
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	c.JSON(http.StatusOK, errmsg.Success(attahcment))
}

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
