package models

type Response struct {
	Status  int         `json:"status"`
	Message int         `json:"message"`
	Errors  []string    `json:"errors"`
	Data    interface{} `json:"data"`
}

type Meta struct {
	TotalPages      int  `json:"total_pages"`
	CurrentPage     int  `json:"current_page"`
	TotalItems      int  `json:"total_items"`
	HasNextPage     bool `json:"has_next_page"`
	HasPreviousPage bool `json:"has_previous_page"`
}

type ResponseWithMeta struct {
	Response
	Meta Meta `json:"meta"`
}
