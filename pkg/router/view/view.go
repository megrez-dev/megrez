package view

import "github.com/gin-gonic/gin"

func RouteView(g *gin.Engine) {
	RouteArticle(g)
	RouteCategory(g)
	RouteComment(g)
	RouteLink(g)
	RouteJournal(g)
	RouteSearch(g)
	RoutePage(g)
	RouteFavicon(g)
}
