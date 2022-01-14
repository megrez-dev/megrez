package site

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/megrez/pkg/model"
)

func RouteOption(g *gin.Engine) {
	g.PUT("/admin/option/:key", setOption)
}

func setOption(c *gin.Context) {
	if c.Param("key") == "" {
		c.JSON(500, "empty key")
		return
	}
	err := model.SetOption(c.Param("key"), c.PostForm("value"))
	if err != nil {
		log.Println("set option failed, err: ", err)
		c.JSON(500, "failed")
		return
	}
	data := fmt.Sprintf("key:%s, value:%s", c.Param("key"), c.PostForm("value"))
	c.JSON(200, data)
}
