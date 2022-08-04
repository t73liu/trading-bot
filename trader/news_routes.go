package main

import (
	"net/http"

	"github.com/t73liu/tradingbot/lib/newsapi"
	"github.com/t73liu/tradingbot/lib/traderdb"
	"github.com/t73liu/tradingbot/lib/utils"

	"github.com/go-chi/chi/v5"
)

func (t *trader) getTopHeadlines(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	userID := getContextUserID(r)
	newsSourceIDs, err := traderdb.GetNewsSourceIDsWithUserID(t.db, userID)
	if err != nil {
		utils.InternalServerError(w, err)
		return
	}
	data, err := t.newsClient.GetTopHeadlinesWithSources(
		newsapi.ArticlesQueryParams{
			Query:   queryValues.Get("q"),
			Sources: newsSourceIDs,
		},
	)
	if err != nil {
		utils.InternalServerError(w, err)
		return
	}
	utils.JSONResponse(w, data)
}

func (t *trader) getUserNewsSources(w http.ResponseWriter, r *http.Request) {
	userID := getContextUserID(r)
	data, err := traderdb.GetNewsSourcesWithUserID(t.db, userID)
	if err != nil {
		utils.InternalServerError(w, err)
		return
	}
	utils.JSONResponse(w, data)
}

func (t *trader) addNewsRoutes(router chi.Router) {
	router.Get("/headlines", t.getTopHeadlines)
	router.Get("/sources/active", t.getUserNewsSources)
}
