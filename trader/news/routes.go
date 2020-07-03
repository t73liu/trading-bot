package news

import (
	"encoding/json"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/julienschmidt/httprouter"
	"github.com/t73liu/trading-bot/lib/newsapi"
	"github.com/t73liu/trading-bot/lib/traderdb"
	"log"
	"net/http"
	"time"
)

type Handlers struct {
	logger *log.Logger
	client *newsapi.Client
	db     *pgxpool.Pool
}

func NewHandlers(logger *log.Logger, client *newsapi.Client, db *pgxpool.Pool) *Handlers {
	return &Handlers{
		logger: logger,
		client: client,
		db:     db,
	}
}

const userId = 1

func (h *Handlers) getTopHeadlines(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	queryValues := r.URL.Query()
	newsSourceIds, err := traderdb.GetNewsSourceIdsByUserId(h.db, userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data, err := h.client.GetTopHeadlinesBySources(
		newsapi.ArticlesQueryParams{
			Query:   queryValues.Get("q"),
			Sources: newsSourceIds,
		},
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handlers) getUserNewsSources(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	data, err := traderdb.GetNewsSourcesByUserId(h.db, userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
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
	router.GET("/api/news/sources/active", h.logTime(h.getUserNewsSources))
}
