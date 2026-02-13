package util

import "net/http"

// Pagination holds the metadata for the response
type Pagination struct {
	Page       int64 `json:"page"`
	Limit      int64 `json:"limit"`
	TotalItems int64 `json:"totalItems"`
	TotalPages int64 `json:"totalPages"`
}

// PaginatedData wraps the actual data and the pagination metadata
type PaginatedData struct {
	Data       any        `json:"data"`
	Pagination Pagination `json:"pagination"`
}

// SendPage assembles and sends the paginated JSON response
func SendPage(w http.ResponseWriter, data any, page, limit, cnt int64) {
	paginatedData := PaginatedData{
		Data: data,
		Pagination: Pagination{
			Page:       page,
			Limit:      limit,
			TotalItems: cnt,
			TotalPages: cnt / limit,
		},
	}

	SendData(w, http.StatusOK, paginatedData)
}