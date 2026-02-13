package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/chamod-ishankha/gogo-project/gogo-backend/internal/config"
	"github.com/chamod-ishankha/gogo-project/gogo-backend/internal/handler"
	"github.com/chamod-ishankha/gogo-project/gogo-backend/internal/middleware"
	"github.com/chamod-ishankha/gogo-project/gogo-backend/internal/repository"
	redisclient "github.com/chamod-ishankha/gogo-project/gogo-backend/pkg/redis"
	"github.com/gorilla/mux"
	"github.com/hellofresh/health-go/v5"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	// 1. Load Config
	cfg, err := config.LoadConfig("driverservice")
	if err != nil {
		log.Fatalf("Critical: Could not load config: %v", err)
	}

	// 2. Database with Connection Pooling
	db, err := sqlx.Connect("postgres", cfg.Database.DSN)
	if err != nil {
		log.Fatalf("Critical: Database connection failed: %v", err)
	}
	db.SetMaxOpenConns(25) // Prevent leaking connections
	db.SetMaxIdleConns(5)

	// 3. Initialize Redis
	redisclient.InitRedis(redisclient.RedisConfig{
		Addr:     cfg.Redis.Addr,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	// 4. Setup Health Checks
	h, _ := health.New(
		health.WithComponent(health.Component{
			Name:    "driver-service",
			Version: "1.0.0",
		}),
		health.WithSystemInfo(), // Adds uptime and host info
	)
	// Example: Add a simple ping check for DB (Simplified for brevity)
	h.Register(health.Config{
		Name:  "postgres",
		Check: func(ctx context.Context) error { return db.PingContext(ctx) },
	})

	// 5. Initialize Repos & Handlers
	driverRepo := &repository.DriverRepository{DB: db}
	vehicleRepo := &repository.VehicleRepository{DB: db}
	locationRepo := &repository.LocationRepository{}

	driverHandler := &handler.DriverHandler{Repo: driverRepo, LocationRepo: locationRepo}
	vehicleHandler := &handler.VehicleHandler{RepoV: vehicleRepo, RepoD: driverRepo}
	locationHandler := &handler.LocationHandler{LocationRepo: locationRepo, DriverRepo: driverRepo}

	// 6. Router Setup
	r := mux.NewRouter()

	// 1. GLOBAL MIDDLEWARE
	// Recovery goes FIRST to protect the whole stack
	r.Use(middleware.RecoveryMiddleware)
	// Logging goes SECOND to record every request
	r.Use(middleware.LoggingMiddleware)

	// Public routes
	public := r.PathPrefix(cfg.Server.Prefix).Subrouter()

	// Routes
	public.Handle("/health", h.Handler())

	// Protected routes (Require Authentication)
	protected := r.PathPrefix(cfg.Server.Prefix).Subrouter()
	protected.Use(middleware.JWTMiddleware)
	protected.Use(middleware.RoleMiddleware("driver"))

	// Routes
	// Driver routes
	protected.HandleFunc("/register", driverHandler.RegisterDriver).Methods("POST")
	protected.HandleFunc("/availability", driverHandler.SetAvailability).Methods("PUT")

	// Vehicle routes
	protected.HandleFunc("/vehicle", vehicleHandler.RegisterVehicle).Methods("POST")
	protected.HandleFunc("/vehicle/me", vehicleHandler.GetMyVehicle).Methods("GET")
	protected.HandleFunc("/vehicle/me", vehicleHandler.UpdateVehicle).Methods("PUT")

	// Location routes
	protected.HandleFunc("/location", locationHandler.UpdateLocation).Methods("PUT")

	// 7. HTTP Server Configuration
	srv := &http.Server{
		Addr:         cfg.Server.Port,
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// 8. Graceful Shutdown Logic
	go func() {
		log.Printf("Driver Service starting on %s%s", cfg.Server.Port, cfg.Server.Prefix)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen error: %s\n", err)
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting gracefully")
}
