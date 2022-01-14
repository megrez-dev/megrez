package router

import (
	"github.com/gin-gonic/gin"
	"github.com/megrez/pkg/middleware/cros"
	"github.com/megrez/pkg/middleware/pongo2gin"
	"github.com/megrez/pkg/router/admin"
	"github.com/megrez/pkg/router/site"
)

func NewRouter() *gin.Engine {
	g := gin.Default()
	g.Use(cros.Cors())
	// load pongo2 for gin
	g.HTMLRender = pongo2gin.TemplatePath("web/site/view")
	// load admin static files
	g.Static("/admin", "web/admin")
	// load static files
	g.Static("/css", "web/admin/css")
	g.Static("/js", "web/admin/js")
	g.Static("/assets", "web/admin/assets")
	// route for site template
	site.RouteSite(g)
	// route for admin api
	admin.RouteAdmin(g)
	return g
}
