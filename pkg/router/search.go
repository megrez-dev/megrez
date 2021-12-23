package router

import (
	"github.com/flosch/pongo2/v4"
	"github.com/gin-gonic/gin"
	"github.com/megrez/pkg/dao"
	"github.com/megrez/pkg/entity/vo"
)

func routeSearch(g *gin.Engine, dao *dao.DAO) {
	DAO = dao
	g.GET("/search", search)
}

func search(c *gin.Context) {

	// TODO: list tags
	// tagPOs, err := DAO.ListAllTags()
	// if err != nil {
	// 	c.Redirect(500, "/error")
	// }
	// var tags []*vo.Tag
	// for _, tagPO := range tagPOs {
	// 	tag, err := vo.GetBriefTagFromPO(tagPO)
	// 	if err != nil {
	// 		c.Redirect(500, "/error")
	// 	}
	// 	tags = append(tags, tag)
	// }

	globalOption, err := vo.GetGlobalOption()
	if err != nil {
		c.Redirect(500, "/error")
	}
	c.HTML(200, "search.html", pongo2.Context{"global": globalOption})
}
