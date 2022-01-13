package view

import (
	"github.com/gin-gonic/gin"
	"github.com/megrez/pkg/model"
)

func RouteTag(g *gin.Engine) {
	g.POST("/admin/tag", createTag)
}

func createTag(c *gin.Context) {
	name := c.PostForm("name")
	slug := c.PostForm("slug")
	description := c.PostForm("description")

	tag := &model.Tag{
		Name:        name,
		Slug:        slug,
		Description: description,
	}
	err := model.CreateTag(tag)
	if err != nil {
		c.JSON(500, "failed to create link")
		return
	}
	c.JSON(200, "success")
}
