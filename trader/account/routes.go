package account

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"tradingbot/lib/traderdb"
	"tradingbot/lib/utils"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
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
	StockIDs []int  `json:"stockIDs"`
}

const userID = 1

func (h *Handlers) getWatchlists(w http.ResponseWriter, _ *http.Request) {
	watchlists, err := traderdb.GetWatchlistsWithUserID(h.db, userID)
	if err != nil {
		utils.JSONError(w, err)
		return
	}
	utils.JSONResponse(w, watchlists)
}

func (h *Handlers) deleteWatchlist(w http.ResponseWriter, r *http.Request) {
	watchlistID, err := getWatchlistID(r)
	if err != nil {
		utils.JSONError(w, err)
		return
	}

	exists, err := traderdb.HasWatchlistWithIDAndUserID(h.db, watchlistID, userID)
	if err != nil {
		utils.JSONError(w, err)
		return
	}
	if !exists {
		utils.JSONError(w, errors.New("watchlist does not exist"))
		return
	}

	err = traderdb.DeleteWatchlistWithID(h.db, watchlistID)
	if err != nil {
		utils.JSONError(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handlers) updateWatchlist(w http.ResponseWriter, r *http.Request) {
	watchlistID, err := getWatchlistID(r)
	if err != nil {
		utils.JSONError(w, err)
		return
	}

	exists, err := traderdb.HasWatchlistWithIDAndUserID(h.db, watchlistID, userID)
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

	err = traderdb.UpdateWatchlistWithID(h.db, watchlistID, body.Name, body.StockIDs)
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

	id, err := traderdb.CreateWatchlist(h.db, userID, body.Name, body.StockIDs)
	if err != nil {
		utils.JSONError(w, err)
		return
	}
	utils.JSONResponse(w, traderdb.Watchlist{
		ID:       id,
		Name:     body.Name,
		StockIDs: body.StockIDs,
	})
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
