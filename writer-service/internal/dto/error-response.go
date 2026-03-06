package dto

import (
	"encoding/json"
	"net/http"
)

func ServerError(w http.ResponseWriter, errs []FieldError) {
	err := ErrorResponse{
		Error: ErrorDetail{
			Code: "INTERNAL_SERVER_ERROR",
			Message: "Writer not found",
			Status: 500,
			Errors: errs,
		},
	}

	writeHeader(w, http.StatusInternalServerError, err)
}

func WriterNotFoundError(w http.ResponseWriter, errs []FieldError) {
	err := ErrorResponse{
		Error: ErrorDetail{
			Code: "NOT_FOUND",
			Message: "Writer not found",
			Status: 404,
			Errors: errs,
		},
	}

	writeHeader(w, http.StatusNotFound, err)
}

func writeHeader(
	w http.ResponseWriter, 
	status int, 
	err ErrorResponse,
) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(err)
}
