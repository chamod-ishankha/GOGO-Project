package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/chamod-ishankha/gogo-project/gogo-backend/internal/middleware"
	"github.com/chamod-ishankha/gogo-project/gogo-backend/internal/model"
	"github.com/chamod-ishankha/gogo-project/gogo-backend/internal/repository"
	"github.com/chamod-ishankha/gogo-project/gogo-backend/pkg/utils"
)

type DriverHandler struct {
	Repo *repository.DriverRepository
}

func (h *DriverHandler) RegisterDriver(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Register endpoint hit")
	claims := r.Context().Value(middleware.UserContextKey).(*utils.Claims)

	// role already validated by middleware
	var req struct {
		LicenseNumber string `json:"license_number"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request"})
		return
	}

	driver := model.Driver{
		UserID:        int64(claims.UserID),
		LicenseNumber: req.LicenseNumber,
		IsActive:      true,
	}

	if err := h.Repo.CreateDriver(&driver); err != nil {
		code, msg := utils.HandleDBError(err)
		w.WriteHeader(code)
		json.NewEncoder(w).Encode(map[string]string{"error": msg})
		return
	}

	json.NewEncoder(w).Encode(driver)
}
