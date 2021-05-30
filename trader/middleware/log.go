package middleware

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func LogResponseTime(logger *log.Logger) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			startTime := time.Now()
			logger.Printf("Starting: %s - %s\n", r.Method, r.URL)
			defer func() {
				logger.Printf(
					"Completed (%dms): %s - %s\n",
					time.Since(startTime).Milliseconds(),
					r.Method,
					r.URL.Path,
				)
			}()
			next.ServeHTTP(w, r)
		})
	}
}
