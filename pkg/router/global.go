package router

import (
	"github.com/gin-gonic/gin"
	"github.com/megrez/pkg/dao"
)

func routeFavicon(g *gin.Engine, dao *dao.DAO) {
	DAO = dao
	g.GET("/favicon.ico", favicon)
}

func favicon(c *gin.Context) {
	// TODO： 这里需要跨域，以后换一种方式
	c.Redirect(302, "http://139.9.201.209/favicon.ico")
}
