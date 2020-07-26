package middleware

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"time"
)

func LogResponseTime(next httprouter.Handle, logger *log.Logger) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		startTime := time.Now()
		logger.Printf("Starting: %s - %s\n", r.Method, r.URL)
		defer func() {
			logger.Printf(
				"Completed (%dms): %s - %s\n",
				time.Now().Sub(startTime).Milliseconds(),
				r.Method,
				r.URL.Path,
			)
		}()
		next(w, r, p)
	}
}
