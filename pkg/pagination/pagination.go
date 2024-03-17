package pagination

import "math"

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
	TotalPages   int `json:"totalPages"`
	TotalEntries int `json:"totalEntries"`
}

func NewResponsePagination(rp RequestPagination, total int) *ResponsePagination {
	totalPages := math.Ceil(float64(total) / float64(rp.PerPage))

	return &ResponsePagination{
		RequestPagination: rp,
		TotalPages:        int(totalPages),
		TotalEntries:      total,
	}
}

func (rp RequestPagination) ToPagination() *Pagination {
	if rp.CurrentPage > 0 {
		rp.CurrentPage--
	}

	return &Pagination{
		Limit:  rp.PerPage,
		Offset: rp.CurrentPage * rp.PerPage,
	}
}
