package model

type Vehicle struct {
	ID          int    `db:"id" json:"id"`
	DriverID    int64  `db:"driver_id" json:"driver_id"`
	VehicleType string `db:"vehicle_type" json:"vehicle_type"`
	Make        string `db:"make" json:"make"`
	Model       string `db:"model" json:"model"`
	Year        int    `db:"year" json:"year"`
	PlateNumber string `db:"plate_number" json:"plate_number"`
	Color       string `db:"color" json:"color"`
	IsActive    bool   `db:"is_active" json:"is_active"`
}
