package models

// Response ..
type Response struct {
	Data interface{} `json:"data,omitempty"`
	Meta Meta        `json:"meta"`
}

// Meta ...
type Meta struct {
	Status  bool   `json:"status" example:"true"`
	Code    int    `json:"code" example:"200"`
	Message string `json:"message" example:"OK"`
}

// ResponsePagination ..
type ResponsePagination struct {
	ResponseCode string      `json:"response_code" example:"00"`
	Message      string      `json:"message" example:"success"`
	Data         interface{} `json:"data"`
	Meta         interface{} `json:"meta"`
}

// MetaPagination ..
type MetaPagination struct {
	CurrentPage int64 `json:"current_page"`
	NextPage    int64 `json:"next_page"`
	PrevPage    int64 `json:"prev_page"`
	TotalPages  int64 `json:"total_pages"`
	TotalCount  int64 `json:"total_count"`
}

// MappingErrorCodes models
type MappingErrorCodes struct {
	Key     string           `json:"key"`
	Content ContentErrorCode `json:"content"`
}

// ContentErrorCode models
type ContentErrorCode struct {
	Status  bool   `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}
