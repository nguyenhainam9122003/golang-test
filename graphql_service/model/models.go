package model

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type PaginatedData struct {
	Page  int         `json:"page"`
	Limit int         `json:"limit"`
	Items interface{} `json:"items"`
}
