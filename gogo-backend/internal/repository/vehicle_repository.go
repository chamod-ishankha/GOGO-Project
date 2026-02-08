package repository

import (
	"github.com/chamod-ishankha/gogo-project/gogo-backend/internal/model"
	"github.com/jmoiron/sqlx"
)

type VehicleRepository struct {
	DB *sqlx.DB
}

func (r *VehicleRepository) CreateVehicle(vehicle *model.Vehicle) error {
	query := `
		INSERT INTO gogo.vehicles
		(driver_id, make, model, year, plate_number)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`
	return r.DB.QueryRow(
		query,
		vehicle.DriverID,
		vehicle.Make,
		vehicle.Model,
		vehicle.Year,
		vehicle.PlateNumber,
	).Scan(&vehicle.ID)
}

func (r *VehicleRepository) VehicleExists(driverID int64) (bool, error) {
	var count int
	err := r.DB.Get(&count,
		`SELECT COUNT(*) FROM gogo.vehicles WHERE driver_id=$1`,
		driverID,
	)
	return count > 0, err
}

func (r *VehicleRepository) GetDriverIDByUserID(userID int64) (int64, error) {
	var driverID int64
	err := r.DB.Get(&driverID,
		`SELECT id FROM gogo.drivers WHERE user_id=$1`,
		userID,
	)
	return driverID, err
}
