package news

import (
	"log"
	"net/http"
	"tradingbot/lib/newsapi"
	"tradingbot/lib/traderdb"
	"tradingbot/lib/utils"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Handlers struct {
	logger *log.Logger
	client *newsapi.Client
	db     *pgxpool.Pool
}

func NewHandlers(logger *log.Logger, db *pgxpool.Pool, client *newsapi.Client) *Handlers {
	return &Handlers{
		logger: logger,
		client: client,
		db:     db,
	}
}

const userID = 1

func (h *Handlers) getTopHeadlines(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	newsSourceIDs, err := traderdb.GetNewsSourceIDsWithUserID(h.db, userID)
	if err != nil {
		utils.JSONError(w, err)
		return
	}
	data, err := h.client.GetTopHeadlinesWithSources(
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

func (h *Handlers) getUserNewsSources(w http.ResponseWriter, _ *http.Request) {
	data, err := traderdb.GetNewsSourcesWithUserID(h.db, userID)
	if err != nil {
		utils.JSONError(w, err)
		return
	}
	utils.JSONResponse(w, data)
}

func (h *Handlers) AddRoutes(router *mux.Router) {
	router.HandleFunc("/headlines", h.getTopHeadlines).Methods("GET")
	router.HandleFunc("/sources/active", h.getUserNewsSources).Methods("GET")
}
