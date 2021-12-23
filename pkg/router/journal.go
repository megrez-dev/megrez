package router

import (
	"log"
	"strconv"

	"github.com/flosch/pongo2/v4"
	"github.com/gin-gonic/gin"
	"github.com/megrez/pkg/dao"
	"github.com/megrez/pkg/entity/po"
	"github.com/megrez/pkg/entity/vo"
)

func routeJournal(g *gin.Engine, dao *dao.DAO) {
	DAO = dao
	g.GET("/journal", listJournal)
	g.GET("/journal/:pageNum", listJournal)
	g.POST("/admin/journal", createJournal)
}

func listJournal(c *gin.Context) {
	var pageNum, pageSize int
	var err error
	if c.Param("pageNum") == "" {
		pageNum = 1
	} else {
		pageNum, err = strconv.Atoi(c.Param("pageNum"))
		if err != nil {
			log.Println("incorrect param pageNum, err:", err)
			// TODO: 应该是 4XX?
			c.Redirect(500, "/error")
		}
	}
	if c.Param("pageSize") == "" {
		pageSize = 10
	} else {
		pageSize, err = strconv.Atoi(c.Param("pageSize"))
		if err != nil {
			log.Println("incorrect param pageSize, err:", err)
			c.Redirect(500, "/error")
		}
	}

	journalPOs, err := DAO.ListAllJournals(pageNum, pageSize)
	if err != nil {
		c.Redirect(500, "/error")
	}
	var journals []*vo.Journal
	for _, journalPO := range journalPOs {
		journal := vo.GetJournalFromPO(journalPO)
		journals = append(journals, &journal)
	}

	globalOption, err := vo.GetGlobalOption()
	if err != nil {
		c.Redirect(500, "/error")
	}
	c.HTML(200, "journal.html", pongo2.Context{"journals": journals, "global": globalOption})
}

func createJournal(c *gin.Context) {
	journal := &po.Journal{
		FormatContent: "测试journal测试journal测试journal测试journal",
		Images:        "https://rawchen.com/sgjpic/23468945914666148.jpg;https://rawchen.com/sgjpic/23468945914666148.jpg;https://rawchen.com/sgjpic/23468945914666148.jpg;https://rawchen.com/sgjpic/23468945914666148.jpg",
		Private:       false,
		Visits:        0,
		Likes:         0,
		Status:        0,
	}
	err := DAO.CreateJournal(journal)
	if err != nil {
		c.JSON(500, "create journal failed")
		return
	}
	c.JSON(200, "success")
}
