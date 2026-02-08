package handler

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/chamod-ishankha/gogo-project/gogo-backend/internal/middleware"
	"github.com/chamod-ishankha/gogo-project/gogo-backend/internal/model"
	"github.com/chamod-ishankha/gogo-project/gogo-backend/internal/repository"
	redisclient "github.com/chamod-ishankha/gogo-project/gogo-backend/pkg/redis"
	"github.com/chamod-ishankha/gogo-project/gogo-backend/pkg/utils"
	"github.com/redis/go-redis/v9"
)

type RideHandler struct {
	RideRepo   *repository.RideRepository
	DriverRepo *repository.DriverRepository
}

// helper to return JSON error
func respondWithError(w http.ResponseWriter, code int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(map[string]string{
		"error": msg,
	})
}

// RequestRide handles ride requests
func (h *RideHandler) RequestRide(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request Ride endpoint hit")

	claims, ok := r.Context().Value(middleware.UserContextKey).(*utils.Claims)
	if !ok {
		respondWithError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var req struct {
		PickupLat float64 `json:"pickup_latitude"`
		PickupLng float64 `json:"pickup_longitude"`
		DropLat   float64 `json:"drop_latitude"`
		DropLng   float64 `json:"drop_longitude"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	ride := &model.Ride{
		RiderID:   claims.UserID,
		PickupLat: req.PickupLat,
		PickupLng: req.PickupLng,
		DropLat:   req.DropLat,
		DropLng:   req.DropLng,
		Status:    "requested",
	}

	if err := h.RideRepo.CreateRide(ride); err != nil {
		code, msg := utils.HandleDBError(err)
		respondWithError(w, code, msg)
		return
	}

	// Find nearest driver
	driverID, err := h.findNearestDriver(req.PickupLat, req.PickupLng)
	if err != nil {
		if err == redis.Nil {
			respondWithError(w, http.StatusNotFound, "No available drivers nearby")
		} else {
			respondWithError(w, http.StatusInternalServerError, "Failed to search nearby drivers")
		}
		return
	}

	if err := h.RideRepo.AssignDriver(ride.ID, driverID); err != nil {
		code, msg := utils.HandleDBError(err)
		respondWithError(w, code, msg)
		return
	}

	if err := h.DriverRepo.SetAvailability(driverID, false); err != nil {
		code, msg := utils.HandleDBError(err)
		respondWithError(w, code, msg)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(ride)
}

// findNearestDriver queries Redis to find the nearest driver
func (h *RideHandler) findNearestDriver(lat, lng float64) (int64, error) {
	locs, err := redisclient.Client.GeoRadius(
		redisclient.Ctx,
		"drivers:locations",
		lng,
		lat,
		&redis.GeoRadiusQuery{
			Radius:    5,
			Unit:      "km",
			Count:     1,
			Sort:      "ASC",
			WithCoord: false,
			WithDist:  false,
		},
	).Result()

	if err != nil {
		return 0, err
	}

	if len(locs) == 0 {
		return 0, redis.Nil
	}

	driverID, err := strconv.Atoi(locs[0].Name)
	if err != nil {
		return 0, fmt.Errorf("invalid driver id")
	}

	return int64(driverID), nil
}

// ChangeStatusRide marks a ride as completed or cancelled
func (h *RideHandler) ChangeStatusRide(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Change Ride Status endpoint hit")

	rideIDStr := r.URL.Query().Get("id")
	rideStatusStr := r.URL.Query().Get("status")

	if rideIDStr == "" {
		respondWithError(w, http.StatusBadRequest, "Ride ID is required")
		return
	}

	if rideStatusStr == "" {
		respondWithError(w, http.StatusBadRequest, "Ride status is required")
		return
	}

	if rideStatusStr != "completed" && rideStatusStr != "cancelled" {
		respondWithError(w, http.StatusBadRequest, "Ride status must be 'completed' or 'cancelled'")
		return
	}

	rideID, err := strconv.Atoi(rideIDStr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid ride ID")
		return
	}

	ride, err := h.RideRepo.GetRideByID(int64(rideID))
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Ride not found")
		return
	}

	fare := 0.0
	if rideStatusStr == "completed" {
		fare = 50.0 + 10*rand.Float64()*10
		if err := h.RideRepo.UpdateFare(int64(rideID), fare); err != nil {
			code, msg := utils.HandleDBError(err)
			respondWithError(w, code, msg)
			return
		}
	}

	if err := h.RideRepo.UpdateStatus(int64(rideID), rideStatusStr); err != nil {
		code, msg := utils.HandleDBError(err)
		respondWithError(w, code, msg)
		return
	}

	if err := h.DriverRepo.SetAvailability(ride.DriverID, true); err != nil {
		code, msg := utils.HandleDBError(err)
		respondWithError(w, code, msg)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"ride_id": rideID,
		"fare":    fare,
		"status":  rideStatusStr,
	})
}
