package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/t73liu/tradingbot/lib/traderdb"
	"github.com/t73liu/tradingbot/lib/utils"

	"github.com/gorilla/mux"
)

type WatchlistRequestBody struct {
	Name     string `json:"name"`
	StockIDs []int  `json:"stockIDs"`
}

func (t *trader) getWatchlists(w http.ResponseWriter, _ *http.Request) {
	watchlists, err := traderdb.GetWatchlistsWithUserID(t.db, userID)
	if err != nil {
		utils.JSONError(w, err)
		return
	}
	utils.JSONResponse(w, watchlists)
}

func (t *trader) deleteWatchlist(w http.ResponseWriter, r *http.Request) {
	watchlistID, err := getWatchlistID(r)
	if err != nil {
		utils.JSONError(w, err)
		return
	}

	exists, err := traderdb.HasWatchlistWithIDAndUserID(t.db, watchlistID, userID)
	if err != nil {
		utils.JSONError(w, err)
		return
	}
	if !exists {
		utils.JSONError(w, errors.New("watchlist does not exist"))
		return
	}

	err = traderdb.DeleteWatchlistWithID(t.db, watchlistID)
	if err != nil {
		utils.JSONError(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (t *trader) updateWatchlist(w http.ResponseWriter, r *http.Request) {
	watchlistID, err := getWatchlistID(r)
	if err != nil {
		utils.JSONError(w, err)
		return
	}

	exists, err := traderdb.HasWatchlistWithIDAndUserID(t.db, watchlistID, userID)
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

	err = traderdb.UpdateWatchlistWithID(t.db, watchlistID, body.Name, body.StockIDs)
	if err != nil {
		utils.JSONError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (t *trader) createWatchlist(w http.ResponseWriter, r *http.Request) {
	var body WatchlistRequestBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		utils.JSONError(w, err)
		return
	}

	id, err := traderdb.CreateWatchlist(t.db, userID, body.Name, body.StockIDs)
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

func (t *trader) AddAccountRoutes(router *mux.Router) {
	router.HandleFunc("/watchlists", t.getWatchlists).Methods("GET")
	router.HandleFunc("/watchlists", t.createWatchlist).Methods("POST")
	router.HandleFunc("/watchlists/{id}", t.updateWatchlist).Methods("PUT")
	router.HandleFunc("/watchlists/{id}", t.deleteWatchlist).Methods("DELETE")
}

func getWatchlistID(r *http.Request) (int, error) {
	vars := mux.Vars(r)
	return strconv.Atoi(vars["id"])
}
