package logrequest

import (
	"log"
	"net/http"
	"time"
)

// Handler will log the HTTP requests
func Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Save start time to calculate duration
		startTime := time.Now()

		next.ServeHTTP(w, r)

		// Calculate request duration and print to screen
		duration := time.Now().Sub(startTime)
		log.Println(r.RemoteAddr, r.Method, r.URL, duration)
	})
}
