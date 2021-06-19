package main

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func LogResponseTime(logger *zap.SugaredLogger) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			startTime := time.Now()
			logger.Infof("Starting: %s - %s\n", r.Method, r.URL)
			defer func() {
				logger.Infof(
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
