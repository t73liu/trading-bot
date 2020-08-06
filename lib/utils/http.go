package utils

import (
	"encoding/json"
	"net"
	"net/http"
	"time"
)

// Use sensible defaults for timeouts for HTTP server
func NewHttpServer(port string, handler *http.Handler) *http.Server {
	return &http.Server{
		Addr:              port,
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      2 * time.Minute,
		IdleTimeout:       5 * time.Minute,
		Handler:           *handler,
	}
}

// Use sensible defaults for timeouts for HTTP client
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

// Equivalent to http.Error but returns JSON instead of text/plain
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
