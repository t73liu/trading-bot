package news

import (
	"encoding/json"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/julienschmidt/httprouter"
	"github.com/t73liu/trading-bot/lib/newsapi"
	"log"
	"net/http"
	"time"
)

type Handlers struct {
	logger *log.Logger
	client *newsapi.Client
	dbPool *pgxpool.Pool
}

func NewHandlers(logger *log.Logger, client *newsapi.Client, dbPool *pgxpool.Pool) *Handlers {
	return &Handlers{
		logger: logger,
		client: client,
		dbPool: dbPool,
	}
}

func (h *Handlers) getTopHeadlines(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	queryValues := r.URL.Query()
	data, err := h.client.GetTopHeadlinesBySources(
		queryValues.Get("q"),
		queryValues.Get("sources"),
		50,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handlers) getSources(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	queryValues := r.URL.Query()
	data, err := h.client.GetSources(
		queryValues.Get("category"),
		queryValues.Get("language"),
		queryValues.Get("country"),
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handlers) logTime(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		startTime := time.Now()
		h.logger.Printf("Starting: %s - %s\n", r.Method, r.URL)
		defer func() {
			h.logger.Printf(
				"Completed (%dms): %s - %s\n",
				time.Now().Sub(startTime).Milliseconds(),
				r.Method,
				r.URL.Path,
			)
		}()
		next(w, r, p)
	}
}

func (h *Handlers) AddRoutes(router *httprouter.Router) {
	router.GET("/api/news/headlines", h.logTime(h.getTopHeadlines))
	router.GET("/api/news/sources", h.logTime(h.getSources))
}
