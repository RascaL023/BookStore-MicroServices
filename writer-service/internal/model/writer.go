package model

type Writer struct {
	Id 			int64	`json:"id"`
	Name 		string	`json:"name"`
	City 		string	`json:"city"`
	Email 		string	`json:"email"`
	IsActive 	bool	`json:"isActive"`
}
