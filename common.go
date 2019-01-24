package zoop

type Info struct {
	Resource string `json:"resource"`
	Uri      string `json:"uri"`
}

type Pagination struct {
	Limit      int  `json:"limit"`
	Offset     int  `json:"offset"`
	HasMore    bool `json:"has_more"`
	QueryCount int  `json:"query_count"`
	Total      int  `json:"total"`
}