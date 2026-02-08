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
		(driver_id, vehicle_type, make, model, year, plate_number, color, is_active)
		VALUES ($1, $2, $3, $4, $5, $6, $7, true)
		RETURNING id
	`
	return r.DB.QueryRow(
		query,
		vehicle.DriverID,
		vehicle.VehicleType,
		vehicle.Make,
		vehicle.Model,
		vehicle.Year,
		vehicle.PlateNumber,
		vehicle.Color,
	).Scan(&vehicle.ID)
}

func (r *VehicleRepository) VehicleExists(driverID int64) (bool, error) {
	var count int
	err := r.DB.Get(&count,
		`SELECT COUNT(*) FROM gogo.vehicles WHERE driver_id=$1 AND is_active=true`,
		driverID,
	)
	return count > 0, err
}

func (r *VehicleRepository) GetByDriverID(driverID int64) (*model.Vehicle, error) {
	var vehicle model.Vehicle
	query := `
		SELECT id, driver_id, vehicle_type, make, model, year, plate_number, color, is_active
		FROM gogo.vehicles
		WHERE driver_id = $1
	`
	err := r.DB.Get(&vehicle, query, driverID)
	if err != nil {
		return nil, err
	}
	return &vehicle, nil
}

func (r *VehicleRepository) Update(vehicle *model.Vehicle) error {
	query := `
		UPDATE gogo.vehicles
		SET vehicle_type = $1,
		    make = $2,
		    model = $3,
		    color = $4,
		    year = $5
		WHERE id = $6
	`
	_, err := r.DB.Exec(
		query,
		vehicle.VehicleType,
		vehicle.Make,
		vehicle.Model,
		vehicle.Color,
		vehicle.Year,
		vehicle.ID,
	)
	return err
}
