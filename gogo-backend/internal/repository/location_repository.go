package repository

import (
	"strconv"

	redisclient "github.com/chamod-ishankha/gogo-project/gogo-backend/pkg/redis"
	"github.com/redis/go-redis/v9"
)

type LocationRepository struct{}

func (r *LocationRepository) UpdateDriverLocation(driverID int64, lat, lng float64) error {
	return redisclient.Client.GeoAdd(
		redisclient.Ctx,
		"drivers:locations",
		&redis.GeoLocation{
			Name:      strconv.Itoa(int(driverID)),
			Latitude:  lat,
			Longitude: lng,
		},
	).Err()
}

func (r *LocationRepository) RemoveDriver(driverID int64) error {
	return redisclient.Client.ZRem(
		redisclient.Ctx,
		"drivers:locations",
		strconv.Itoa(int(driverID)),
	).Err()
}
