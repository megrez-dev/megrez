package router

import (
	"github.com/gin-gonic/gin"
	"github.com/megrez/pkg/middleware/pongo2gin"
	"github.com/megrez/pkg/router/view"
)

func NewRouter() *gin.Engine {
	g := gin.Default()
	g.HTMLRender = pongo2gin.TemplatePath("web/site/view")
	g.Static("/admin", "web/admin")
	g.Static("/css", "web/admin/css")
	g.Static("/js", "web/admin/js")
	g.Static("/assets", "web/admin/assets")
	// admin := g.Group("/admin")
	view.RouteArticle(g)
	view.RouteCategory(g)
	view.RouteComment(g)
	view.RouteOption(g)
	view.RouteLink(g)
	view.RouteJournal(g)
	view.RouteSearch(g)
	view.RouteFavicon(g)
	return g
}
