package repository

import (
	"github.com/chamod-ishankha/gogo-project/gogo-backend/internal/model"
	"github.com/jmoiron/sqlx"
)

type DriverRepository struct {
	DB *sqlx.DB
}

func (r *DriverRepository) CreateDriver(driver *model.Driver) error {
	query := `
        INSERT INTO gogo.drivers (user_id, license_number, is_active)
        VALUES ($1, $2, $3)
        RETURNING id
    `
	return r.DB.QueryRow(
		query,
		driver.UserID,
		driver.LicenseNumber,
		driver.IsActive,
	).Scan(&driver.ID)
}

func (r *DriverRepository) GetByUserID(userID int64) (*model.Driver, error) {
	var driver model.Driver
	query := `
		SELECT id, user_id, license_number, is_active, is_available
		FROM gogo.drivers
		WHERE user_id = $1
	`
	err := r.DB.Get(&driver, query, userID)
	if err != nil {
		return nil, err
	}
	return &driver, nil
}

func (r *DriverRepository) SetAvailability(userID int64, available bool) error {
	query := `
		UPDATE gogo.drivers
		SET is_available = $1
		WHERE user_id = $2
	`
	_, err := r.DB.Exec(query, available, userID)
	return err
}

func (r *DriverRepository) GetByDriverID(driverId int64) (*model.Driver, error) {
	var driver model.Driver
	query := `
		SELECT id, user_id, license_number, is_active, is_available
		FROM gogo.drivers
		WHERE driver_id = $1
	`
	err := r.DB.Get(&driver, query, driverId)
	if err != nil {
		return nil, err
	}
	return &driver, nil
}
