package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/chamod-ishankha/gogo-project/gogo-backend/internal/middleware"
	"github.com/chamod-ishankha/gogo-project/gogo-backend/internal/repository"
	"github.com/chamod-ishankha/gogo-project/gogo-backend/pkg/utils"
)

type LocationHandler struct {
	LocationRepo *repository.LocationRepository
	DriverRepo   *repository.DriverRepository
}

func (h *LocationHandler) UpdateLocation(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UpdateLocation endpoint hit")
	claims := r.Context().Value(middleware.UserContextKey).(*utils.Claims)

	driver, err := h.DriverRepo.GetByUserID(claims.UserID)
	if err != nil {
		code, msg := utils.HandleDBError(err)
		w.WriteHeader(code)
		json.NewEncoder(w).Encode(map[string]string{"error": msg})
		return
	}

	var req struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		code, msg := utils.HandleDBError(err)
		w.WriteHeader(code)
		json.NewEncoder(w).Encode(map[string]string{"error": msg})
		return
	}

	err = h.LocationRepo.UpdateDriverLocation(driver.ID, req.Latitude, req.Longitude)
	if err != nil {
		code, msg := utils.HandleDBError(err)
		w.WriteHeader(code)
		json.NewEncoder(w).Encode(map[string]string{"error": msg})
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Location updated",
	})
}
