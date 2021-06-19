package utils

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"time"
)

// NewHttpServer creates a http.Server with sensible timeouts
func NewHttpServer(port int, handler *http.Handler) *http.Server {
	return &http.Server{
		Addr:              fmt.Sprintf(":%d", port),
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      2 * time.Minute,
		IdleTimeout:       5 * time.Minute,
		Handler:           *handler,
	}
}

// NewHttpClient creates a http.Client with sensible timeouts
func NewHttpClient() *http.Client {
	return &http.Client{
		Timeout: 15 * time.Second,
		Transport: &http.Transport{
			// Time spent establishing TCP connection
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			// Time spent performing TLS handshake
			TLSHandshakeTimeout: 10 * time.Second,
			// Time spent reading response headers
			ResponseHeaderTimeout: 10 * time.Second,
		},
	}
}

// JSONError similar to http.Error but returns application/json instead of text/plain
func JSONError(w http.ResponseWriter, err error) {
	SetJSONHeader(w)
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusInternalServerError)
	_ = json.NewEncoder(w).Encode(err.Error())
}

func JSONResponse(w http.ResponseWriter, data interface{}) {
	SetJSONHeader(w)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		JSONError(w, err)
	}
}

func SetJSONHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
}
