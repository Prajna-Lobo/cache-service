package model

type Data struct {
	Key   string
	Value interface{}
}

type Pagination struct {
	PageNumber int
	PageSize   int
}

type PaginationResponse struct {
	TotalPageCount int
}

type Response struct {
	Data       []Data             `json:"data"`
	Pagination PaginationResponse `json:"pagination"`
}
