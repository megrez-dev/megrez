package router

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/megrez/pkg/dao"
)

func routeOption(g *gin.Engine, dao *dao.DAO) {
	DAO = dao
	g.PUT("/admin/option/:key", setOption)
}

func setOption(c *gin.Context) {
	if c.Param("key") == "" {
		c.JSON(500, "empty key")
		return
	}
	err := DAO.SetOption(c.Param("key"), c.PostForm("value"))
	if err != nil {
		log.Println("set option failed, err: ", err)
		c.JSON(500, "failed")
		return
	}
	data := fmt.Sprintf("key:%s, value:%s", c.Param("key"), c.PostForm("value"))
	c.JSON(200, data)
}
