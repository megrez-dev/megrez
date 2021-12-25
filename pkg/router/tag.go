package router

import (
	"github.com/gin-gonic/gin"
	"github.com/megrez/pkg/dao"
	"github.com/megrez/pkg/entity/po"
)

func routeTag(g *gin.Engine, dao *dao.DAO) {
	DAO = dao
	g.POST("/admin/tag", createTag)
}

func createTag(c *gin.Context) {
	name := c.PostForm("name")
	slug := c.PostForm("slug")
	description := c.PostForm("description")

	tag := &po.Tag{
		Name: name,
		Slug: slug,
		Description: description,
	}
	err := DAO.CreateTag(tag)
	if err != nil {
		c.JSON(500, "failed to create link")
		return
	}
	c.JSON(200, "success")
}
