package vo

import (
	"strings"
	"time"

	"github.com/megrez/pkg/model"
)

type Journal struct {
	FormatContent string
	Images        []string
	Private       bool
	Likes         int64
	Visits        int64
	Status        int
	PublishTime   time.Time
}

func GetJournalFromPO(journal model.Journal) Journal {
	var images []string
	if journal.Images != "" {
		images = strings.Split(journal.Images, ";")
	}
	journalVO := Journal{
		FormatContent: journal.FormatContent,
		Images:        images,
		Private:       journal.Private,
		Likes:         journal.Likes,
		Visits:        journal.Visits,
		Status:        journal.Status,
		PublishTime:   journal.CreatedAt,
	}
	return journalVO
}
