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

	claims, ok := r.Context().Value(middleware.UserContextKey).(*utils.Claims)
	if !ok {
		utils.WriteJSONError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	driver, err := h.DriverRepo.GetByUserID(claims.UserID)
	if err != nil {
		code, msg := utils.HandleDBError(err)
		utils.WriteJSONError(w, code, msg)
		return
	}

	var req struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteJSONError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if req.Latitude == 0 || req.Longitude == 0 {
		utils.WriteJSONError(w, http.StatusBadRequest, "Latitude and longitude are required")
		return
	}

	if err := h.LocationRepo.UpdateDriverLocation(driver.ID, req.Latitude, req.Longitude); err != nil {
		code, msg := utils.HandleDBError(err)
		utils.WriteJSONError(w, code, msg)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]string{
		"message": "Location updated",
	})

	response := map[string]string{
		"message": "Location updated successfully",
	}

	utils.WriteJSONResponse(w, http.StatusOK, response)
}
