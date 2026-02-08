package main

import (
	"log"
	"net/http"

	"github.com/chamod-ishankha/gogo-project/gogo-backend/internal/config"
	"github.com/chamod-ishankha/gogo-project/gogo-backend/internal/handler"
	"github.com/chamod-ishankha/gogo-project/gogo-backend/internal/middleware"
	"github.com/chamod-ishankha/gogo-project/gogo-backend/internal/repository"
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

	driverRepo := &repository.DriverRepository{DB: db}
	driverHandler := &handler.DriverHandler{Repo: driverRepo}

	vehicleRepo := &repository.VehicleRepository{DB: db}
	vehicleHandler := &handler.VehicleHandler{Repo: vehicleRepo}

	r := mux.NewRouter()

	// Use Prefix from config
	protected := r.PathPrefix(cfg.Server.Prefix).Subrouter()
	protected.Use(middleware.JWTMiddleware)
	protected.Use(middleware.RoleMiddleware("driver"))

	protected.HandleFunc("/register", driverHandler.RegisterDriver).Methods("POST")
	protected.HandleFunc("/vehicle", vehicleHandler.RegisterVehicle).Methods("POST")

	log.Printf("Driver Service running at %s with prefix %s", cfg.Server.Port, cfg.Server.Prefix)
	http.ListenAndServe(cfg.Server.Port, r)
}
