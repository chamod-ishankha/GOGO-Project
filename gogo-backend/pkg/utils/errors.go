package utils

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/lib/pq"
)

func HandleDBError(err error) (int, string) {
	// ✅ Handle "no rows" FIRST
	if errors.Is(err, sql.ErrNoRows) {
		return http.StatusBadRequest, "Record not found."
	}

	var pqErr *pq.Error
	if errors.As(err, &pqErr) {
		switch pqErr.Code {
		case "23505":
			return http.StatusBadRequest, "Duplicate entry already exists."
		case "23503":
			return http.StatusBadRequest, "Related record not found."
		case "23502":
			return http.StatusBadRequest, "Missing required field."
		case "28P01":
			return http.StatusUnauthorized, "Database authentication failed."
		default:
			return http.StatusInternalServerError, "Database error occurred."
		}
	}

	// Fallback
	return http.StatusInternalServerError, err.Error()
}

func WriteJSONError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(map[string]string{
		"error": message,
	})
}
