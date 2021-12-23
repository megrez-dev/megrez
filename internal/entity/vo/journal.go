package vo

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/megrez/internal/entity/po"
)

type Journal struct {
	FormatContent string
	Images        []string
	Private       bool
	Likes         uint
	Visits        uint
	Status        int
	PublishTime   time.Time
}

func GetJournalFromPO(po po.Journal) Journal {
	var images []string
	if po.Images != "" {
		images = strings.Split(po.Images, ";")
	}
	journal := Journal{
		FormatContent: po.FormatContent,
		Images:        images,
		Private:       po.Private,
		Likes:         po.Likes,
		Visits:        po.Visits,
		Status:        po.Status,
		PublishTime:   po.CreatedAt,
	}
	b, _ := json.Marshal(journal)
	fmt.Println(string(b))
	return journal
}
