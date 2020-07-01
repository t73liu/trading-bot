package account

import (
	"encoding/json"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/julienschmidt/httprouter"
	"github.com/t73liu/trading-bot/lib/traderdb"
	"log"
	"net/http"
)

type Handlers struct {
	logger *log.Logger
	db     *pgxpool.Pool
}

func NewHandlers(logger *log.Logger, db *pgxpool.Pool) *Handlers {
	return &Handlers{
		logger: logger,
		db:     db,
	}
}

func (h *Handlers) handleGetWatchlists(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	watchlists, err := traderdb.GetWatchlistsByUserId(h.db, 1)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		if err = json.NewEncoder(w).Encode(watchlists); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func (h *Handlers) AddRoutes(router *httprouter.Router) {
	router.GET("/api/account/watchlists", h.handleGetWatchlists)
	//router.POST("/api/account/watchlists", h.handleUpdateWatchlists)
}
