package redisclient

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()
var Client *redis.Client

type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}

// InitRedis initializes Redis client with provided config
func InitRedis(cfg RedisConfig) {
	Client = redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	// Test the connection
	_, err := Client.Ping(Ctx).Result()
	if err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}

	log.Println("Redis connected successfully")
}
