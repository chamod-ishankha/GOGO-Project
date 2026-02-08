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

	claims, ok := r.Context().Value(middleware.UserContextKey).(*utils.Claims)
	if !ok {
		utils.WriteJSONError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var req model.Vehicle
	req.IsActive = true

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteJSONError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	driver, err := h.RepoD.GetByUserID(claims.UserID)
	if err != nil {
		utils.WriteJSONError(w, http.StatusBadRequest, "Driver profile not found")
		return
	}

	req.DriverID = driver.ID

	exists, err := h.RepoV.VehicleExists(driver.ID)
	if err != nil {
		code, msg := utils.HandleDBError(err)
		utils.WriteJSONError(w, code, msg)
		return
	}

	if exists {
		utils.WriteJSONError(w, http.StatusBadRequest, "Vehicle already registered")
		return
	}

	if err := h.RepoV.CreateVehicle(&req); err != nil {
		fmt.Printf("Error creating vehicle: %v\n", err)
		code, msg := utils.HandleDBError(err)
		utils.WriteJSONError(w, code, msg)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(req)
}

func (h *VehicleHandler) GetMyVehicle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get My Vehicle endpoint hit")

	claims, ok := r.Context().Value(middleware.UserContextKey).(*utils.Claims)
	if !ok {
		utils.WriteJSONError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	driver, err := h.RepoD.GetByUserID(claims.UserID)
	if err != nil {
		code, msg := utils.HandleDBError(err)
		utils.WriteJSONError(w, code, msg)
		return
	}

	exists, err := h.RepoV.VehicleExists(driver.ID)
	if err != nil {
		code, msg := utils.HandleDBError(err)
		utils.WriteJSONError(w, code, msg)
		return
	}

	if !exists {
		utils.WriteJSONError(w, http.StatusBadRequest, "Vehicle not registered")
		return
	}

	vehicle, err := h.RepoV.GetByDriverID(driver.ID)
	if err != nil {
		code, msg := utils.HandleDBError(err)
		utils.WriteJSONError(w, code, msg)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(vehicle)
}

func (h *VehicleHandler) UpdateVehicle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update Vehicle endpoint hit")

	claims, ok := r.Context().Value(middleware.UserContextKey).(*utils.Claims)
	if !ok {
		utils.WriteJSONError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	driver, err := h.RepoD.GetByUserID(claims.UserID)
	if err != nil {
		utils.WriteJSONError(w, http.StatusBadRequest, "Driver profile not found")
		return
	}

	exists, err := h.RepoV.VehicleExists(driver.ID)
	if err != nil {
		code, msg := utils.HandleDBError(err)
		utils.WriteJSONError(w, code, msg)
		return
	}

	if !exists {
		utils.WriteJSONError(w, http.StatusBadRequest, "Vehicle not registered")
		return
	}

	vehicle, err := h.RepoV.GetByDriverID(driver.ID)
	if err != nil {
		code, msg := utils.HandleDBError(err)
		utils.WriteJSONError(w, code, msg)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(vehicle); err != nil {
		utils.WriteJSONError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := h.RepoV.Update(vehicle); err != nil {
		code, msg := utils.HandleDBError(err)
		utils.WriteJSONError(w, code, msg)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]string{
		"message": "Vehicle updated",
	})
}
