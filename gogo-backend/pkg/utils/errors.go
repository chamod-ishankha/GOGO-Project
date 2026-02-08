package utils

import (
	"errors"
	"net/http"

	"github.com/lib/pq"
)

// DBErrorResponse helps standardize error messages
type DBErrorResponse struct {
	StatusCode int
	Message    string
}

func HandleDBError(err error) (int, string) {
	var pqErr *pq.Error

	// Check if it's a Postgres driver error
	if errors.As(err, &pqErr) {
		switch pqErr.Code {
		case "23505": // unique_violation
			return http.StatusBadRequest, "The record already exists (duplicate entry)."
		case "23503": // foreign_key_violation
			return http.StatusBadRequest, "Related record not found (invalid ID)."
		case "23502": // not_null_violation
			return http.StatusBadRequest, "A required field is missing."
		case "28P01": // invalid_password
			return http.StatusUnauthorized, "Database authentication failed."
		}
	}

	// Default to 500 for generic SQL errors or connection issues
	return http.StatusInternalServerError, "An internal database error occurred."
}
