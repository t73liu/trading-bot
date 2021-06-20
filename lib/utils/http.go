package utils

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"
)

type ContextKey string

// NewHttpServer creates a http.Server with sensible timeouts
func NewHttpServer(port int, handler *http.Handler) *http.Server {
	return &http.Server{
		Addr:              fmt.Sprintf(":%d", port),
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      time.Minute,
		IdleTimeout:       time.Minute,
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

func IsJSONContentType(r *http.Request) bool {
	contentType := r.Header.Get("Content-Type")
	return strings.Contains(contentType, "application/json")
}

// InternalServerError similar to http.Error but returns application/json
// instead of text/plain
func InternalServerError(w http.ResponseWriter, err error) {
	JSONError(w, err, http.StatusInternalServerError)
}

func UnauthenticatedError(w http.ResponseWriter, err error) {
	JSONError(w, err, http.StatusForbidden)
}

func NotFoundError(w http.ResponseWriter, err error) {
	JSONError(w, err, http.StatusNotFound)
}

func JSONError(w http.ResponseWriter, err error, statusCode int) {
	SetJSONHeader(w)
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(err.Error())
}

func JSONResponse(w http.ResponseWriter, data interface{}) {
	SetJSONHeader(w)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		InternalServerError(w, err)
	}
}

func SetJSONHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
}

func SecureHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		w.Header().Set("X-Frame-Options", "SAMEORIGIN")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		next.ServeHTTP(w, r)
	})
}

func PanicRecovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			// Recover from panic in handler goroutine
			if err := recover(); err != nil {
				w.Header().Set("Connection", "close")
				InternalServerError(w, fmt.Errorf("%s", err))
			}
		}()
		next.ServeHTTP(w, r)
	})
}
