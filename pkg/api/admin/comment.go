package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/megrez/pkg/entity/dto"
	"github.com/megrez/pkg/entity/vo"
	"github.com/megrez/pkg/log"
	"github.com/megrez/pkg/model"
	"github.com/megrez/pkg/utils/errmsg"
	"net/http"
	"strconv"
	"time"
)

func CreateComment(c *gin.Context) {
	var data dto.CreateCommentForm
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
	comment.Role = 1
	comment.Email = user.Email
	site, err := model.GetOptionByKey(vo.OptionKeyBlogURL)
	if err != nil {
		log.Error("get option failed:", err.Error())
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	comment.Site = site
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
	var commentDTOs []dto.CommentListDTO
	for _, comment := range comments {
		commentDTO := dto.CommentListDTO{}
		err := commentDTO.LoadFromModel(comment)
		if err != nil {
			log.Error(err)
			c.JSON(http.StatusOK, errmsg.Error())
			return
		}
		commentDTOs = append(commentDTOs, commentDTO)
	}
	total, err := model.CountAllArticles()
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
