package view

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/megrez/pkg/log"
	"github.com/megrez/pkg/model"
)

func RouteFavicon(g *gin.Engine) {
	g.GET("/favicon.ico", favicon)
}

func favicon(c *gin.Context) {
	// TODO： 这里用的重定向方式实现，后面看看有没有更好的实现方式
	url, err := model.GetOptionByKey(model.OptionKeyBlogFavicon)
	if err != nil {
		log.Error("incorrect param pageNum, err:", err.Error())
		c.Redirect(http.StatusInternalServerError, "/error")
	}
	c.Redirect(http.StatusFound, url)
}
