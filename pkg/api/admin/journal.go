package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/megrez/pkg/entity/dto"
	"github.com/megrez/pkg/log"
	"github.com/megrez/pkg/model"
	"github.com/megrez/pkg/utils/errmsg"
	"net/http"
	"strconv"
	"time"
)

// CreateJournal godoc
// @Summary create journal
// @Schemes http https
// @Description create journal
// @Accept application/json
// @Param Authorization header string false "Authorization"
// @Param req body dto.CreateJournalForm true "body"
// @Success 200 {object} errmsg.Response{}
// @Router /api/admin/journal [post]
func CreateJournal(c *gin.Context) {
	var data dto.CreateJournalForm
	err := c.ShouldBindJSON(&data)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusOK, errmsg.ErrorInvalidParam)
		return
	}
	journal := &model.Journal{
		Content: data.Content,
		Private: data.Private,
		Status:  data.Status,
	}
	images := ""
	for _, image := range data.Images {
		if images == "" {
			images = image
		} else {
			images += ";" + image
		}
	}
	journal.Images = images
	journal.CreateTime = time.Now()
	journal.UpdateTime = time.Now()
	tx := model.BeginTx()
	err = model.CreateJournal(tx, journal)
	if err != nil {
		log.Error(err.Error())
		tx.Rollback()
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	tx.Commit()
	c.JSON(http.StatusOK, errmsg.Success(nil))
}

// ListJournals godoc
// @Summary list journals
// @Schemes http https
// @Description list journals
// @Accept application/json
// @Param Authorization header string false "Authorization"
// @Param pageNum query int false "page num"
// @Param pageSize query int false "page size"
// @Success 200 {object} errmsg.Response{data=dto.Pagination{list=[]dto.JournalDTO}}
// @Router /api/admin/journals [get]
func ListJournals(c *gin.Context) {
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
	journals, err := model.ListAllJournals(pageNum, pageSize)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	var journalDTOs []dto.JournalDTO
	for _, journal := range journals {
		journalDTO := dto.JournalDTO{}
		journalDTO.LoadFromModel(journal)
		if err != nil {
			log.Error(err)
			c.JSON(http.StatusOK, errmsg.Error())
			return
		}
		journalDTOs = append(journalDTOs, journalDTO)
	}
	total, err := model.CountAllJournals()
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	pagination := dto.Pagination{
		List:     journalDTOs,
		Current:  pageNum,
		PageSize: pageSize,
		Total:    total,
	}
	c.JSON(http.StatusOK, errmsg.Success(pagination))
}
