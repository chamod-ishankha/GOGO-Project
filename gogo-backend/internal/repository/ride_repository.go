package repository

import (
	"github.com/chamod-ishankha/gogo-project/gogo-backend/internal/model"
	"github.com/jmoiron/sqlx"
)

type RideRepository struct {
	DB *sqlx.DB
}

func (r *RideRepository) CreateRide(ride *model.Ride) error {
	query := `
		INSERT INTO gogo.rides
		(rider_id, pickup_latitude, pickup_longitude, drop_latitude, drop_longitude)
		VALUES ($1,$2,$3,$4,$5)
		RETURNING id
	`
	return r.DB.QueryRow(
		query,
		ride.RiderID, ride.PickupLat, ride.PickupLng, ride.DropLat, ride.DropLng,
	).Scan(&ride.ID)
}

func (r *RideRepository) AssignDriver(rideID, driverID int64) error {
	query := `
		UPDATE gogo.rides
		SET driver_id=$1, status='accepted'
		WHERE id=$2
	`
	_, err := r.DB.Exec(query, driverID, rideID)
	return err
}

func (r *RideRepository) UpdateStatus(rideID int64, status string) error {
	query := `
		UPDATE gogo.rides
		SET status=$1
		WHERE id=$2
	`
	_, err := r.DB.Exec(query, status, rideID)
	return err
}

func (r *RideRepository) UpdateFare(rideID int64, fare float64) error {
	query := `
		UPDATE gogo.rides
		SET fare=$1
		WHERE id=$2
	`
	_, err := r.DB.Exec(query, fare, rideID)
	return err
}

func (r *RideRepository) GetRideByID(rideID int64) (*model.Ride, error) {
	var ride model.Ride
	query := `
		SELECT * FROM gogo.rides
		WHERE id=$1
	`
	err := r.DB.Get(&ride, query, rideID)
	return &ride, err
}
