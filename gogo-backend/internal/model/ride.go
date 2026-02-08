package model

type Ride struct {
	ID        int64   `db:"id" json:"id"`
	RiderID   int64   `db:"rider_id" json:"rider_id"`
	DriverID  int64   `db:"driver_id" json:"driver_id"`
	PickupLat float64 `db:"pickup_latitude" json:"pickup_latitude"`
	PickupLng float64 `db:"pickup_longitude" json:"pickup_longitude"`
	DropLat   float64 `db:"drop_latitude" json:"drop_latitude"`
	DropLng   float64 `db:"drop_longitude" json:"drop_longitude"`
	Status    string  `db:"status" json:"status"`
	Fare      float64 `db:"fare" json:"fare"`
	CreatedAt string  `db:"created_at" json:"created_at"`
	UpdatedAt string  `db:"updated_at" json:"updated_at"`
}
