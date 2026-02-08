package model

type Vehicle struct {
	ID          int    `db:"id" json:"id"`
	DriverID    int64  `db:"driver_id" json:"driver_id"`
	Make        string `db:"make" json:"make"`
	Model       string `db:"model" json:"model"`
	Year        int    `db:"year" json:"year"`
	PlateNumber string `db:"plate_number" json:"plate_number"`
}
