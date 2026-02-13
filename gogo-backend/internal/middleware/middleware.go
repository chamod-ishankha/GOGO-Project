package middleware

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
	"time"
)

// LoggingMiddleware records the details of each HTTP request
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Pass the request to the next handler
		next.ServeHTTP(w, r)

		// Log the results after the request is finished
		log.Printf(
			"METHOD: %s | PATH: %s | DURATION: %v | IP: %s",
			r.Method,
			r.URL.Path,
			time.Since(start),
			r.RemoteAddr,
		)
	})
}

// RecoveryMiddleware catches panics and prevents the server from crashing
func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				// Log the error and the stack trace for debugging
				log.Printf("PANIC RECOVERED: %v\n%s", err, debug.Stack())

				// Send a 500 Internal Server Error to the client
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"error": "An internal server error occurred"}`)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
