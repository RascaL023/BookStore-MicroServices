package dto

type ErrorResponse struct {
	Error ErrorDetail `json:"error"`
}

type ErrorDetail struct {
	Code	string 				`json:"code"`
	Message	string 				`json:"message"`
	Status	int 				`json:"status"`
	Errors []FieldError		 	`json:"errors,omitempty"`
}

type FieldError struct {
	Field string	`json:"field"`
	Message string	`json:"message"`
}
