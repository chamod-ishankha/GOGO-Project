package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/chamod-ishankha/gogo-project/gogo-backend/internal/config"
	"github.com/rs/cors"
)

func createProxy(target string) *httputil.ReverseProxy {
	destination, _ := url.Parse(target)
	return httputil.NewSingleHostReverseProxy(destination)
}

func main() {
	// 1. Load Gateway Config
	cfg, err := config.LoadConfig("gateway")
	if err != nil {
		log.Fatalf("Could not load gateway config: %v", err)
	}

	mux := http.NewServeMux()

	// 2. Set up dynamic routing based on config
	for name, svc := range cfg.Services {
		proxy := createProxy(svc.URL)
		prefix := svc.Prefix

		log.Printf("%s Routing %s to %s", name, prefix, svc.URL)

		mux.HandleFunc(prefix+"/", func(w http.ResponseWriter, r *http.Request) {
			proxy.ServeHTTP(w, r)
		})
	}

	// 3. Centralized CORS logic
	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"}, // Your React App
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}).Handler(mux)

	log.Printf("API Gateway running on %s", cfg.Server.Port)
	log.Fatal(http.ListenAndServe(cfg.Server.Port, handler))
}
