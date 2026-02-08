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

type VehicleHandler struct {
	Repo *repository.VehicleRepository
}

func (h *VehicleHandler) RegisterVehicle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Register Vehicle endpoint hit")
	claims := r.Context().Value(middleware.UserContextKey).(*utils.Claims)

	var req model.Vehicle
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request"})
		return
	}

	// Check if vehicle already exists
	exists, err := h.Repo.VehicleExists(int64(claims.UserID))
	if err != nil {
		code, msg := utils.HandleDBError(err)
		w.WriteHeader(code)
		json.NewEncoder(w).Encode(map[string]string{"error": msg})
		return
	}
	if exists {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Vehicle already registered"})
		return
	}

	// Get driver ID from user ID
	driverID, err := h.Repo.GetDriverIDByUserID(int64(claims.UserID))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Driver profile not found"})
		return
	}

	// Assign driver ID to request body
	req.DriverID = driverID

	if err := h.Repo.CreateVehicle(&req); err != nil {
		code, msg := utils.HandleDBError(err)
		fmt.Printf("Error creating vehicle: %v\n", err)
		w.WriteHeader(code)
		json.NewEncoder(w).Encode(map[string]string{"error": msg})
		return
	}

	json.NewEncoder(w).Encode(req)
}
