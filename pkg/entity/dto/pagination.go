package dto

type Pagination struct {
	List     interface{} `json:"list"`
	Current  int         `json:"current"`
	PageSize int         `json:"pageSize"`
	Total    int64       `json:"total"`
}
