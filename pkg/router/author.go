package router

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/megrez/pkg/dao"
	"github.com/megrez/pkg/entity/po"
)

func routeAuthor(g *gin.Engine, dao *dao.DAO) {
	DAO = dao
	g.POST("/admin/author", createAuthor)
}

func createAuthor(c *gin.Context) {
	author := &po.Author{
		Name:   "AlkaidChen",
		Mail:   "362774405@qq.com",
		Site:   "alkaidchen.com",
		Role:   1,
		Avatar: "https://cdn.rawchen.com/logo/alkaidchen.jpg",
	}
	err := DAO.CreateAuthor(author)
	if err != nil {
		log.Println("create author failed, err: ", err)
		c.JSON(500, "failed")
	}
	c.JSON(200, "success")
}
