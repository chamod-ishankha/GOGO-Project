package main

import (
	"context"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/chamod-ishankha/gogo-project/gogo-backend/internal/config"
	"github.com/chamod-ishankha/gogo-project/gogo-backend/internal/middleware"
	"github.com/chamod-ishankha/gogo-project/gogo-backend/pkg/utils"
	"github.com/rs/cors"
)

func createProxy(target string) *httputil.ReverseProxy {
	destination, _ := url.Parse(target)
	proxy := httputil.NewSingleHostReverseProxy(destination)

	// UPGRADE: Custom Error Handler
	// If the backend service is down, return 503 instead of a generic error
	proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
		log.Printf("Proxy error: %v", err)
		utils.WriteJSONError(w, http.StatusServiceUnavailable, "Backend service is unreachable")
	}

	return proxy
}

func main() {
	cfg, err := config.LoadConfig("gateway")
	if err != nil {
		log.Fatalf("Critical: Could not load gateway config: %v", err)
	}

	mux := http.NewServeMux()

	for name, svc := range cfg.Services {
		// Redeclare variables inside the loop to create a unique copy for the closure
		name := name
		svc := svc

		proxy := createProxy(svc.URL)
		prefix := svc.Prefix

		log.Printf("[Gateway] %s Routing %s -> %s", name, prefix, svc.URL)

		mux.HandleFunc(prefix+"/", func(w http.ResponseWriter, r *http.Request) {
			// Now 'proxy' refers to the correct instance for this specific prefix
			proxy.ServeHTTP(w, r)
		})
	}

	// UPGRADE: Wrap with Middleware
	// We use the same Recovery and Logging we built for the microservices
	finalHandler := middleware.RecoveryMiddleware(mux)
	finalHandler = middleware.LoggingMiddleware(finalHandler)

	// UPGRADE: Apply CORS
	finalHandler = cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		Debug:            false,
	}).Handler(finalHandler)

	// UPGRADE: Structured Server with Timeouts
	srv := &http.Server{
		Addr:         cfg.Server.Port,
		Handler:      finalHandler,
		ReadTimeout:  10 * time.Second, // Gateway needs slightly higher timeouts
		WriteTimeout: 20 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// UPGRADE: Graceful Shutdown
	go func() {
		log.Printf("API Gateway running on %s", cfg.Server.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down Gateway...")
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Gateway forced to shutdown:", err)
	}
	log.Println("Gateway exited gracefully")
}
