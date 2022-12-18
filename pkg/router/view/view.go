package view

import "github.com/gin-gonic/gin"

func RouteView(g *gin.Engine) {
	RouteArticle(g)
	RouteCategory(g)
	RouteLink(g)
	RouteJournal(g)
	RouteSearch(g)
	RouteAbout(g)
	RoutePage(g)
	RouteFavicon(g)
}
