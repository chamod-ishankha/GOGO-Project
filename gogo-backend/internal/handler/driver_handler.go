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
	Repo         *repository.DriverRepository
	LocationRepo *repository.LocationRepository
}

func (h *DriverHandler) RegisterDriver(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Register endpoint hit")

	claims, ok := r.Context().Value(middleware.UserContextKey).(*utils.Claims)
	if !ok {
		utils.WriteJSONError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var req struct {
		LicenseNumber string `json:"license_number"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteJSONError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if req.LicenseNumber == "" {
		utils.WriteJSONError(w, http.StatusBadRequest, "License number is required")
		return
	}

	driver := model.Driver{
		UserID:        int64(claims.UserID),
		LicenseNumber: req.LicenseNumber,
		IsActive:      true,
	}

	if err := h.Repo.CreateDriver(&driver); err != nil {
		code, msg := utils.HandleDBError(err)
		utils.WriteJSONError(w, code, msg)
		return
	}

	utils.WriteJSONResponse(w, http.StatusCreated, driver)
}

func (h *DriverHandler) SetAvailability(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Set availability endpoint hit")

	claims, ok := r.Context().Value(middleware.UserContextKey).(*utils.Claims)
	if !ok {
		utils.WriteJSONError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var req struct {
		IsAvailable bool `json:"is_available"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteJSONError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := h.Repo.SetAvailability(claims.UserID, req.IsAvailable); err != nil {
		utils.WriteJSONError(w, http.StatusInternalServerError, "Failed to update availability")
		return
	}

	driver, err := h.Repo.GetByUserID(claims.UserID)
	if err != nil {
		code, msg := utils.HandleDBError(err)
		utils.WriteJSONError(w, code, msg)
		return
	}

	status := "offline"
	if req.IsAvailable {
		status = "online"
	} else {
		fmt.Println("Driver is now offline, removing from location tracking")
		if err := h.LocationRepo.RemoveDriver(driver.ID); err != nil {
			utils.WriteJSONError(
				w,
				http.StatusInternalServerError,
				"Failed to remove driver from location tracking",
			)
			return
		}
	}

	utils.WriteJSONResponse(w, http.StatusOK, map[string]string{
		"message": "You are now " + status,
	})
}
