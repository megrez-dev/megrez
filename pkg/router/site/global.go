package site

import (
	"github.com/gin-gonic/gin"
)

func RouteFavicon(g *gin.Engine) {
	g.GET("/favicon.ico", favicon)
}

func favicon(c *gin.Context) {
	// TODO： 这里需要跨域，以后换一种方式
	c.Redirect(302, "https://alkaidchen-1257721976.cos.ap-guangzhou.myqcloud.com/blog/static/images/favicon.ico")
}
