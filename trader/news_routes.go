package main

import (
	"net/http"

	"github.com/t73liu/tradingbot/lib/newsapi"
	"github.com/t73liu/tradingbot/lib/traderdb"
	"github.com/t73liu/tradingbot/lib/utils"

	"github.com/gorilla/mux"
)

func (t *trader) getTopHeadlines(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	newsSourceIDs, err := traderdb.GetNewsSourceIDsWithUserID(t.db, userID)
	if err != nil {
		utils.JSONError(w, err)
		return
	}
	data, err := t.newsClient.GetTopHeadlinesWithSources(
		newsapi.ArticlesQueryParams{
			Query:   queryValues.Get("q"),
			Sources: newsSourceIDs,
		},
	)
	if err != nil {
		utils.JSONError(w, err)
		return
	}
	utils.JSONResponse(w, data)
}

func (t *trader) getUserNewsSources(w http.ResponseWriter, _ *http.Request) {
	data, err := traderdb.GetNewsSourcesWithUserID(t.db, userID)
	if err != nil {
		utils.JSONError(w, err)
		return
	}
	utils.JSONResponse(w, data)
}

func (t *trader) AddNewsRoutes(router *mux.Router) {
	router.HandleFunc("/headlines", t.getTopHeadlines).Methods("GET")
	router.HandleFunc("/sources/active", t.getUserNewsSources).Methods("GET")
}
