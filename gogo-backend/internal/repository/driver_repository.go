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
