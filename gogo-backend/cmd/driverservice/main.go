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
	cfg, err := config.LoadConfig("driverservice")
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
	vehicleRepo := &repository.VehicleRepository{DB: db}
	locationRepo := &repository.LocationRepository{}

	driverHandler := &handler.DriverHandler{Repo: driverRepo, LocationRepo: locationRepo}
	vehicleHandler := &handler.VehicleHandler{RepoV: vehicleRepo, RepoD: driverRepo}
	locationHandler := &handler.LocationHandler{LocationRepo: locationRepo, DriverRepo: driverRepo}

	r := mux.NewRouter()

	// Use Prefix from config
	protected := r.PathPrefix(cfg.Server.Prefix).Subrouter()
	protected.Use(middleware.JWTMiddleware)
	protected.Use(middleware.RoleMiddleware("driver"))

	// Driver routes
	protected.HandleFunc("/register", driverHandler.RegisterDriver).Methods("POST")
	protected.HandleFunc("/availability", driverHandler.SetAvailability).Methods("PUT")

	// Vehicle routes
	protected.HandleFunc("/vehicle", vehicleHandler.RegisterVehicle).Methods("POST")
	protected.HandleFunc("/vehicle/me", vehicleHandler.GetMyVehicle).Methods("GET")
	protected.HandleFunc("/vehicle/me", vehicleHandler.UpdateVehicle).Methods("PUT")

	// Location routes
	protected.HandleFunc("/location", locationHandler.UpdateLocation).Methods("PUT")

	log.Printf("Driver Service running at %s with prefix %s", cfg.Server.Port, cfg.Server.Prefix)
	err = http.ListenAndServe(cfg.Server.Port, r)
	if err != nil {
		fmt.Println("Failed to start server:", err)
		return
	}
}
