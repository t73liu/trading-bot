package news

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"tradingbot/lib/newsapi"
	"tradingbot/lib/traderdb"
	"tradingbot/lib/utils"
	"tradingbot/trader/middleware"
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

const userId = 1

func (h *Handlers) getTopHeadlines(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	queryValues := r.URL.Query()
	newsSourceIds, err := traderdb.GetNewsSourceIdsByUserId(h.db, userId)
	if err != nil {
		utils.JSONError(w, err)
		return
	}
	data, err := h.client.GetTopHeadlinesBySources(
		newsapi.ArticlesQueryParams{
			Query:   queryValues.Get("q"),
			Sources: newsSourceIds,
		},
	)
	if err != nil {
		utils.JSONError(w, err)
		return
	}
	utils.JSONResponse(w, data)
}

func (h *Handlers) getUserNewsSources(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	data, err := traderdb.GetNewsSourcesByUserId(h.db, userId)
	if err != nil {
		utils.JSONError(w, err)
		return
	}
	utils.JSONResponse(w, data)
}

func (h *Handlers) AddRoutes(router *httprouter.Router) {
	router.GET(
		"/api/news/headlines",
		middleware.LogResponseTime(h.getTopHeadlines, h.logger),
	)
	router.GET(
		"/api/news/sources/active",
		middleware.LogResponseTime(h.getUserNewsSources, h.logger),
	)
}
