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
	RepoV *repository.VehicleRepository
	RepoD *repository.DriverRepository
}

func (h *VehicleHandler) RegisterVehicle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Register Vehicle endpoint hit")
	claims := r.Context().Value(middleware.UserContextKey).(*utils.Claims)

	var req model.Vehicle
	req.IsActive = true
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request"})
		return
	}

	// Get driver ID from user ID
	driver, err := h.RepoD.GetByUserID(claims.UserID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Driver profile not found"})
		return
	}

	// Assign driver ID to request body
	req.DriverID = driver.ID

	// Check if vehicle already exists
	exists, err := h.RepoV.VehicleExists(driver.ID)
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

	if err := h.RepoV.CreateVehicle(&req); err != nil {
		code, msg := utils.HandleDBError(err)
		fmt.Printf("Error creating vehicle: %v\n", err)
		w.WriteHeader(code)
		json.NewEncoder(w).Encode(map[string]string{"error": msg})
		return
	}

	json.NewEncoder(w).Encode(req)
}

func (h *VehicleHandler) GetMyVehicle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get My Vehicle endpoint hit")
	claims := r.Context().Value(middleware.UserContextKey).(*utils.Claims)

	driver, err := h.RepoD.GetByUserID(claims.UserID)
	if err != nil {
		code, msg := utils.HandleDBError(err)
		w.WriteHeader(code)
		json.NewEncoder(w).Encode(map[string]string{"error": msg})
		return
	}

	// Check if vehicle registered for this driver
	exists, err := h.RepoV.VehicleExists(driver.ID)
	if err != nil {
		code, msg := utils.HandleDBError(err)
		w.WriteHeader(code)
		json.NewEncoder(w).Encode(map[string]string{"error": msg})
		return
	}

	if !exists {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Vehicle Not Registered"})
		return
	}

	vehicle, err := h.RepoV.GetByDriverID(driver.ID)
	if err != nil {
		code, msg := utils.HandleDBError(err)
		w.WriteHeader(code)
		json.NewEncoder(w).Encode(map[string]string{"error": msg})
		return
	}

	json.NewEncoder(w).Encode(vehicle)
}

func (h *VehicleHandler) UpdateVehicle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update Vehicle endpoint hit")
	claims := r.Context().Value(middleware.UserContextKey).(*utils.Claims)

	driver, _ := h.RepoD.GetByUserID(claims.UserID)
	// Check if vehicle registered for this driver
	exists, err := h.RepoV.VehicleExists(driver.ID)
	if err != nil {
		code, msg := utils.HandleDBError(err)
		w.WriteHeader(code)
		json.NewEncoder(w).Encode(map[string]string{"error": msg})
		return
	}
	if !exists {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Vehicle Not Registered"})
		return
	}

	vehicle, err := h.RepoV.GetByDriverID(driver.ID)
	if err != nil {
		code, msg := utils.HandleDBError(err)
		w.WriteHeader(code)
		json.NewEncoder(w).Encode(map[string]string{"error": msg})
		return
	}

	if err := json.NewDecoder(r.Body).Decode(vehicle); err != nil {
		code, msg := utils.HandleDBError(err)
		w.WriteHeader(code)
		json.NewEncoder(w).Encode(map[string]string{"error": msg})
		return
	}

	if err := h.RepoV.Update(vehicle); err != nil {
		code, msg := utils.HandleDBError(err)
		w.WriteHeader(code)
		json.NewEncoder(w).Encode(map[string]string{"error": msg})
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Vehicle updated",
	})
}
