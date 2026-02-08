package main

import (
	"fmt"
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
	cfg, err := config.LoadConfig("userservice")
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

	repo := &repository.UserRepository{DB: db}
	userHandler := &handler.UserHandler{Repo: repo}

	r := mux.NewRouter()

	// Use Prefix from config
	apiPath := r.PathPrefix(cfg.Server.Prefix).Subrouter()
	protected := apiPath.NewRoute().Subrouter()
	protected.Use(middleware.JWTMiddleware)
	protected.Use(middleware.RoleMiddleware("rider"))

	// Auth routes
	apiPath.HandleFunc("/register", userHandler.Register).Methods("POST")
	apiPath.HandleFunc("/login", userHandler.Login).Methods("POST")

	// User routes
	protected.HandleFunc("/me", userHandler.GetMe).Methods("GET")
	protected.HandleFunc("/me", userHandler.UpdateMe).Methods("PUT")

	log.Printf("Driver Service running at %s with prefix %s", cfg.Server.Port, cfg.Server.Prefix)
	err = http.ListenAndServe(cfg.Server.Port, r)
	if err != nil {
		fmt.Println("Failed to start server:", err)
		return
	}
}
