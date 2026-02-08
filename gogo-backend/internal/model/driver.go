package model

type Driver struct {
	ID            int64  `db:"id" json:"id"`
	UserID        int64  `db:"user_id" json:"user_id"`
	LicenseNumber string `db:"license_number" json:"license_number"`
	IsActive      bool   `db:"is_active" json:"is_active"`
	IsAvailable   bool   `db:"is_available" json:"is_available"`
}
