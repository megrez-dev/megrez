package router

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/megrez/internal/dao"
	"github.com/megrez/internal/entity/po"
)

func routeCategory(g *gin.Engine, dao *dao.DAO) {
	DAO = dao
	g.POST("/admin/category", createCategory)
}

func createCategory(c *gin.Context) {
	category := &po.Category{
		Name:   "默认分类",
		Slug:   "default",
		Status: 0,
	}
	err := DAO.CreateCategory(category)
	if err != nil {
		log.Println("create category failed, err: ", err)
		c.JSON(500, "failed")
	}
	c.JSON(200, "success")
}
