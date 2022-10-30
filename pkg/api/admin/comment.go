package admin

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/megrez/pkg/entity/dto"
	admindto "github.com/megrez/pkg/entity/dto/admin"
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
// @Param req body admindto.CreateCommentForm true "body"
// @Success 200 {object} errmsg.Response{}
// @Router /api/admin/comment [post]
func CreateComment(c *gin.Context) {
	var data admindto.CreateCommentForm
	err := c.ShouldBindJSON(&data)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusOK, errmsg.ErrorInvalidParam)
		return
	}
	comment := model.Comment{
		ArticleID: data.ArticleID,
		PageID:    data.PageID,
		Content:   data.Content,
		RootID:    data.RootID,
		ParentID:  data.ParentID,
		Type:      data.Type,
	}
	uid := c.GetUint("uid")
	if uid == 0 {
		log.Error("get uid from gin context is 0")
		c.JSON(http.StatusOK, errmsg.Fail(errmsg.ErrorTokenInvalid))
		return
	}
	user, err := model.GetUserByID(uid)
	if err != nil {
		log.Error("get user failed:", err.Error())
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	comment.Author = user.Nickname
	comment.Role = model.RoleTypeAdmin
	comment.Email = user.Email
	url, err := model.GetOptionByKey(model.OptionKeyBlogURL)
	if err != nil {
		log.Error("get option failed:", err.Error())
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	comment.URL = url
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
// @Success 200 {object} errmsg.Response{data=dto.Pagination{list=[]admindto.CommentListDTO}}
// @Router /api/admin/comments [get]
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
	comments, err := model.ListAllComments(pageNum, pageSize)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	var commentDTOs []admindto.CommentListDTO
	for _, comment := range comments {
		commentDTO := admindto.CommentListDTO{}
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
