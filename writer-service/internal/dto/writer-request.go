package dto

type WriterRequest struct {
	Name string		`json:"name"`
	City string		`json:"city"`
	Email string	`json:"email"`
	IsActive bool	`json:"isActive"`
}

type WriterPatchRequest struct {
	Name *string	`json:"name,omitempty"`
	City *string	`json:"city,omitempty"`
	Email *string	`json:"email,omitempty"`
	IsActive *bool	`json:"isActive,omitempty"`
}

