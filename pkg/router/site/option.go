package site

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/megrez/pkg/model"
)

func RouteOption(g *gin.Engine) {
	g.PUT("/admin/option/:key", setOption)
}

func setOption(c *gin.Context) {
	if c.Param("key") == "" {
		c.JSON(http.StatusInternalServerError, "empty key")
		return
	}
	err := model.SetOption(c.Param("key"), c.PostForm("value"))
	if err != nil {
		log.Println("set option failed, err: ", err)
		c.JSON(http.StatusInternalServerError, "failed")
		return
	}
	data := fmt.Sprintf("key:%s, value:%s", c.Param("key"), c.PostForm("value"))
	c.JSON(http.StatusOK, data)
}
