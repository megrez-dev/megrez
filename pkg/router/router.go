package router

import (
	"github.com/gin-gonic/gin"
	"github.com/megrez/pkg/dao"
	"github.com/megrez/pkg/middleware/pongo2gin"
)

func NewRouter(d *dao.DAO) *gin.Engine {
	g := gin.Default()
	g.HTMLRender = pongo2gin.TemplatePath("web/site/view")
	// admin := g.Group("/admin")
	routeArticle(g, d)
	routeCategory(g, d)
	routeComment(g, d)
	routeAuthor(g, d)
	routeOption(g, d)
	routeLink(g, d)
	routeAbout(g, d)
	routeJournal(g, d)
	routeSearch(g, d)
	routeFavicon(g, d)
	return g
}
