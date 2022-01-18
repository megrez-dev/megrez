package site

import "github.com/gin-gonic/gin"

func RouteSite(g *gin.Engine) {
	RouteArticle(g)
	RouteCategory(g)
	RouteComment(g)
	RouteOption(g)
	RouteLink(g)
	RouteJournal(g)
	RouteSearch(g)
	RoutePage(g)
	RouteFavicon(g)
}
