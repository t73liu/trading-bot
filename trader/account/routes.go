package account

import (
	"encoding/json"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/julienschmidt/httprouter"
	"github.com/t73liu/trading-bot/lib/traderdb"
	"log"
	"net/http"
	"strconv"
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

type WatchlistRequestBody struct {
	Name     string `json:"name"`
	StockIds []int  `json:"stockIds"`
}

const userId = 1

func (h *Handlers) getWatchlists(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	watchlists, err := traderdb.GetWatchlistsByUserId(h.db, userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(watchlists); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handlers) deleteWatchlist(w http.ResponseWriter, _ *http.Request, params httprouter.Params) {
	watchlistId, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	exists, err := traderdb.HasWatchlistWithIdAndUserId(h.db, watchlistId, userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !exists {
		http.Error(w, "Watchlist does not exist", http.StatusNotFound)
		return
	}

	err = traderdb.DeleteWatchlistById(h.db, watchlistId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handlers) updateWatchlist(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	watchlistId, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	exists, err := traderdb.HasWatchlistWithIdAndUserId(h.db, watchlistId, userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !exists {
		http.Error(w, "Watchlist does not exist", http.StatusNotFound)
		return
	}

	var body WatchlistRequestBody
	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = traderdb.UpdateWatchlistById(h.db, watchlistId, body.Name, body.StockIds)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handlers) createWatchlist(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var body WatchlistRequestBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := traderdb.CreateWatchlist(h.db, userId, body.Name, body.StockIds)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	watchlist, err := traderdb.GetWatchlistById(h.db, id)
	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(watchlist); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handlers) AddRoutes(router *httprouter.Router) {
	router.GET("/api/account/watchlists", h.getWatchlists)
	router.PUT("/api/account/watchlists/:id", h.updateWatchlist)
	router.DELETE("/api/account/watchlists/:id", h.deleteWatchlist)
	router.POST("/api/account/watchlists", h.createWatchlist)
}
