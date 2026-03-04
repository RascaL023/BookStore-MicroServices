package dto

type Meta struct {
	Page int `json:"page"`
	Size int `json:"size"`
	Total int `json:"total"`
	Current int `json:"current"`
}

type Response[T any] struct {
	Data T 		`json:"data"`
	Meta *Meta 	`json:"meta,omitempty"`
}

type WriterResponse struct {
	Id 		int64	`json:"id"`
	Name 	string	`json:"name"`
	City 	string	`json:"city"`
	Email 	string	`json:"email"`
}

