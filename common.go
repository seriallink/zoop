package zoop

type Pagination struct {
	Limit      int  `json:"limit"`
	Offset     int  `json:"offset"`
	HasMore    bool `json:"has_more"`
	QueryCount int  `json:"query_count"`
	Total      int  `json:"total"`
}

type DeleteResponse struct {
	Id       string `json:"id"`
	Resource string `json:"resource"`
	Deleted  bool   `json:"deleted"`
}
