package pagination

type Pagination struct {
	Limit  int
	Offset int
}

type RequestPagination struct {
	CurrentPage int `json:"currentPage" query:"current_page"`
	PerPage     int `json:"perPage" query:"per_page"`
}

type ResponsePagination struct {
	RequestPagination
	TotalPages int `json:"totalPages"`
}

func NewResponsePagination(rp RequestPagination, total int) *ResponsePagination {
	return &ResponsePagination{
		RequestPagination: rp,
		TotalPages:        total,
	}
}

func (rp RequestPagination) ToPagination() *Pagination {
	if rp.CurrentPage > 0 {
		rp.CurrentPage--
	}

	return &Pagination{
		Limit:  rp.PerPage - rp.CurrentPage,
		Offset: rp.CurrentPage,
	}
}
