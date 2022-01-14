package site

import (
	"github.com/flosch/pongo2/v4"
	"github.com/gin-gonic/gin"
	"github.com/megrez/pkg/entity/vo"
	"github.com/megrez/pkg/model"
)

func RouteSearch(g *gin.Engine) {
	g.GET("/search", search)
}

func search(c *gin.Context) {
	tagPOs, err := model.ListAllTags()
	if err != nil {
		c.Redirect(500, "/error")
	}
	var tags []*vo.TagWithArticlesNum
	for _, tagPO := range tagPOs {
		tag, err := vo.GetTagWithArticlesNumFromPO(tagPO)
		if err != nil {
			c.Redirect(500, "/error")
		}
		tags = append(tags, tag)
	}

	globalOption, err := vo.GetGlobalOption()
	if err != nil {
		c.Redirect(500, "/error")
	}
	c.HTML(200, "search.html", pongo2.Context{"tags": tags, "global": globalOption})
}
