package vo

import (
	"strings"
	"time"

	"github.com/megrez/pkg/model"
)

type Journal struct {
	Content     string
	Images      []string
	Private     bool
	Likes       int64
	Visits      int64
	Status      int
	PublishTime time.Time
}

func GetJournalFromPO(journal model.Journal) Journal {
	var images []string
	if journal.Images != "" {
		images = strings.Split(journal.Images, ";")
	}
	journalVO := Journal{
		Content:     journal.Content,
		Images:      images,
		Private:     journal.Private,
		Likes:       journal.Likes,
		Visits:      journal.Visits,
		Status:      journal.Status,
		PublishTime: journal.CreateTime,
	}
	return journalVO
}
