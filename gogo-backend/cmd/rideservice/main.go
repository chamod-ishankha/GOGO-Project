package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/chamod-ishankha/gogo-project/gogo-backend/internal/config"
	"github.com/chamod-ishankha/gogo-project/gogo-backend/internal/handler"
	"github.com/chamod-ishankha/gogo-project/gogo-backend/internal/middleware"
	"github.com/chamod-ishankha/gogo-project/gogo-backend/internal/repository"
	redisclient "github.com/chamod-ishankha/gogo-project/gogo-backend/pkg/redis"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	// Load config
	cfg, err := config.LoadConfig("rideservice")
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}
	// Use DB DSN from config
	db, err := sqlx.Connect("postgres", cfg.Database.DSN)
	if err != nil {
		log.Fatalln(err)
	}

	var schema string
	db.Get(&schema, "SELECT current_schema()")
	log.Println("Connected to schema:", schema)

	// Initialize Redis
	redisCfg := redisclient.RedisConfig{
		Addr:     cfg.Redis.Addr,     // K8s service name
		Password: cfg.Redis.Password, // no password yet
		DB:       cfg.Redis.DB,       // default DB
	}
	redisclient.InitRedis(redisCfg)

	driverRepo := &repository.DriverRepository{DB: db}
	rideRepo := &repository.RideRepository{DB: db}

	rideHandler := &handler.RideHandler{RideRepo: rideRepo, DriverRepo: driverRepo}

	r := mux.NewRouter()

	// Use Prefix from config
	protected := r.PathPrefix(cfg.Server.Prefix).Subrouter()
	protected.Use(middleware.JWTMiddleware)
	protected.Use(middleware.RoleMiddleware("rider"))

	// Ride endpoints
	protected.HandleFunc("", rideHandler.RequestRide).Methods("POST")
	protected.HandleFunc("/status", rideHandler.ChangeStatusRide).Methods("PUT")

	log.Printf("Rider Service running at %s with prefix %s", cfg.Server.Port, cfg.Server.Prefix)
	err = http.ListenAndServe(cfg.Server.Port, r)
	if err != nil {
		fmt.Println("Failed to start server:", err)
		return
	}
}
