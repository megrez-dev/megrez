package dto

type CreateJournalForm struct {
	Content string   `json:"content"`
	Images  []string `json:"images"`
	Private bool     `json:"private"`
	Status  int      `json:"status"`
}
