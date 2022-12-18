package view

import (
	"log"
	"net/http"
	"strconv"

	"github.com/flosch/pongo2/v6"
	"github.com/gin-gonic/gin"
	"github.com/megrez/pkg/entity/vo"
	"github.com/megrez/pkg/model"
)

func RouteJournal(g *gin.Engine) {
	g.GET("/journal", listJournal)
	g.GET("/journal/:pageNum", listJournal)
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
			c.Redirect(http.StatusInternalServerError, "/error")
		}
	}
	if c.Param("pageSize") == "" {
		pageSize = 10
	} else {
		pageSize, err = strconv.Atoi(c.Param("pageSize"))
		if err != nil {
			log.Println("incorrect param pageSize, err:", err)
			c.Redirect(http.StatusInternalServerError, "/error")
		}
	}

	journalPOs, err := model.ListAllJournals(pageNum, pageSize)
	if err != nil {
		c.Redirect(http.StatusInternalServerError, "/error")
	}
	var journals []*vo.Journal
	for _, journalPO := range journalPOs {
		journal := vo.GetJournalFromPO(journalPO)
		journals = append(journals, &journal)
	}

	globalOption, err := vo.GetGlobalOption()
	if err != nil {
		c.Redirect(http.StatusInternalServerError, "/error")
	}

	journalsNum, err := model.CountAllJournals()
	if err != nil {
		c.Redirect(http.StatusInternalServerError, "/error")
	}
	page, err := model.GetPageBySlugAndType("journal", model.PageTypeBuildIn)
	if err != nil {
		c.Redirect(http.StatusInternalServerError, "/error")
	}
	pageVO := vo.GetPageFromPO(page)
	pagination := vo.CalculatePagination(pageNum, pageSize, int(journalsNum))
	c.HTML(http.StatusOK, "journal.html", pongo2.Context{
		"page":       pageVO,
		"journals":   journals,
		"pagination": pagination,
		"global":     globalOption,
	})
}
