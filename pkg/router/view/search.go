package view

import (
	"github.com/flosch/pongo2/v4"
	"github.com/gin-gonic/gin"
	"github.com/megrez/pkg/entity/vo"
	"github.com/megrez/pkg/model"
	"net/http"
)

func RouteSearch(g *gin.Engine) {
	g.GET("/search", search)
}

func search(c *gin.Context) {
	tagPOs, err := model.ListAllTags()
	if err != nil {
		c.Redirect(http.StatusInternalServerError, "/error")
	}
	var tags []*vo.TagWithArticlesNum
	for _, tagPO := range tagPOs {
		tag, err := vo.GetTagWithArticlesNumFromPO(tagPO)
		if err != nil {
			c.Redirect(http.StatusInternalServerError, "/error")
		}
		tags = append(tags, tag)
	}

	globalOption, err := vo.GetGlobalOption()
	if err != nil {
		c.Redirect(http.StatusInternalServerError, "/error")
	}
	c.HTML(http.StatusOK, "search.html", pongo2.Context{"tags": tags, "global": globalOption})
}
