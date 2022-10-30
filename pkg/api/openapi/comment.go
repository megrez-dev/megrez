package openapi

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/megrez/pkg/entity/dto"
	openapidto "github.com/megrez/pkg/entity/dto/openapi"
	"github.com/megrez/pkg/log"
	"github.com/megrez/pkg/model"
	"github.com/megrez/pkg/utils/errmsg"
)

// CreateComment godoc
// @Summary create comment
// @Schemes http https
// @Description create comment
// @Accept application/json
// @Param Authorization header string false "Authorization"
// @Param req body openapidto.CreateCommentForm true "body"
// @Success 200 {object} errmsg.Response{}
// @Router /api/comment [post]
func CreateComment(c *gin.Context) {
	var data openapidto.CreateCommentForm
	err := c.ShouldBindJSON(&data)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusOK, errmsg.ErrorInvalidParam)
		return
	}
	comment := model.Comment{
		Author:    data.Author,
		Email:     data.Mail,
		URL:       data.URL,
		Content:   data.Content,
		ArticleID: data.ArticleID,
		PageID:    data.PageID,
		RootID:    data.RootID,
		ParentID:  data.ParentID,
		Type:      data.Type,
		Role:      "guest",
	}
	// todo: 查找设置项评论间隔，go-cache 查询间隔时间内 ip 是否发布过评论
	// todo: 检查如果有 jwt token,则解析,如果解析出来的token为有效,则从数据库中查用户信息并且设置
	comment.Agent = c.Request.UserAgent()
	comment.IP = c.ClientIP()
	comment.Status = 0
	comment.CreateTime = time.Now()
	comment.UpdateTime = time.Now()
	tx := model.BeginTx()
	err = model.CreateComment(tx, &comment)
	if err != nil {
		log.Error("create comment failed:", err.Error())
		tx.Rollback()
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	tx.Commit()
	c.JSON(http.StatusOK, errmsg.Success(nil))
}

// ListComments godoc
// @Summary list comments
// @Schemes http https
// @Description list comments
// @Accept application/json
// @Param Authorization header string false "Authorization"
// @Param pageNum query int false "page num"
// @Param pageSize query int false "page size"
// @Success 200 {object} errmsg.Response{data=dto.Pagination{list=[]openapidto.CommentListDTO}}
// @Router /api/{type}/{id}/comments [get]
func ListComments(c *gin.Context) {
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
	var parentComments []model.Comment
	tp := c.Param("type")
	var id uint
	pid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, errmsg.Fail(errmsg.ErrorInvalidParam))
		return
	}
	id = uint(pid)
	switch tp {
	case model.CommentTypeArticle:
		if parentComments, err = model.ListRootCommentsByArticleID(id, pageNum, pageSize); err != nil {
			log.Error("list root comments for article %d failed, err: %s")
		}
	}
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	var commentDTOs []openapidto.CommentDTO
	for _, comment := range parentComments {
		commentDTO := openapidto.CommentDTO{}
		err := commentDTO.LoadFromModel(comment)
		if err != nil {
			log.Error(err)
			c.JSON(http.StatusOK, errmsg.Error())
			return
		}
		commentDTOs = append(commentDTOs, commentDTO)
	}
	total, err := model.CountAllComments()
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	pagination := dto.Pagination{
		List:     commentDTOs,
		Current:  pageNum,
		PageSize: pageSize,
		Total:    total,
	}
	c.JSON(http.StatusOK, errmsg.Success(pagination))
}

// DeleteComment godoc
// @Summary delete comment by comment id
// @Schemes http https
// @Description delete comment by comment id
// @Accept application/json
// @Param Authorization header string false "Authorization"
// @Param id path int true "comment id"
// @Success 200 {object} errmsg.Response{}
// @Router /api/admin/comment/{id} [delete]
func DeleteComment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, errmsg.Fail(errmsg.ErrorInvalidParam))
		return
	}
	tx := model.BeginTx()
	err = model.DeleteCommentByID(tx, uint(id))
	if err != nil {
		log.Error("delete comment error:", err.Error())
		tx.Rollback()
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	err = model.DeleteCommentsByParentID(tx, uint(id))
	if err != nil {
		log.Error("delete comment error:", err.Error())
		tx.Rollback()
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	tx.Commit()
	c.JSON(http.StatusOK, errmsg.Success(nil))
}
