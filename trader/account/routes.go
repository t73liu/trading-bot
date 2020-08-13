package account

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"net/http"
	"strconv"
	"tradingbot/lib/traderdb"
	"tradingbot/lib/utils"
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

func (h *Handlers) getWatchlists(w http.ResponseWriter, _ *http.Request) {
	watchlists, err := traderdb.GetWatchlistsByUserId(h.db, userId)
	if err != nil {
		utils.JSONError(w, err)
		return
	}
	utils.JSONResponse(w, watchlists)
}

func (h *Handlers) deleteWatchlist(w http.ResponseWriter, r *http.Request) {
	watchlistId, err := getWatchlistID(r)
	if err != nil {
		utils.JSONError(w, err)
		return
	}

	exists, err := traderdb.HasWatchlistWithIdAndUserId(h.db, watchlistId, userId)
	if err != nil {
		utils.JSONError(w, err)
		return
	}
	if !exists {
		utils.JSONError(w, errors.New("watchlist does not exist"))
		return
	}

	err = traderdb.DeleteWatchlistById(h.db, watchlistId)
	if err != nil {
		utils.JSONError(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handlers) updateWatchlist(w http.ResponseWriter, r *http.Request) {
	watchlistId, err := getWatchlistID(r)
	if err != nil {
		utils.JSONError(w, err)
		return
	}

	exists, err := traderdb.HasWatchlistWithIdAndUserId(h.db, watchlistId, userId)
	if err != nil {
		utils.JSONError(w, err)
		return
	}
	if !exists {
		utils.JSONError(w, errors.New("watchlist does not exist"))
		return
	}

	var body WatchlistRequestBody
	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		utils.JSONError(w, err)
		return
	}

	err = traderdb.UpdateWatchlistById(h.db, watchlistId, body.Name, body.StockIds)
	if err != nil {
		utils.JSONError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handlers) createWatchlist(w http.ResponseWriter, r *http.Request) {
	var body WatchlistRequestBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		utils.JSONError(w, err)
		return
	}

	id, err := traderdb.CreateWatchlist(h.db, userId, body.Name, body.StockIds)
	if err != nil {
		utils.JSONError(w, err)
		return
	}

	watchlist, err := traderdb.GetWatchlistById(h.db, id)
	utils.JSONResponse(w, watchlist)
}

func (h *Handlers) AddRoutes(router *mux.Router) {
	router.HandleFunc("/watchlists", h.getWatchlists).Methods("GET")
	router.HandleFunc("/watchlists", h.createWatchlist).Methods("POST")
	router.HandleFunc("/watchlists/{id}", h.updateWatchlist).Methods("PUT")
	router.HandleFunc("/watchlists/{id}", h.deleteWatchlist).Methods("DELETE")
}

func getWatchlistID(r *http.Request) (int, error) {
	vars := mux.Vars(r)
	return strconv.Atoi(vars["id"])
}
