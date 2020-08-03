package stock

import (
	"encoding/json"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/julienschmidt/httprouter"
	analyze "github.com/t73liu/trading-bot/lib/technical-analysis"
	"github.com/t73liu/trading-bot/lib/traderdb"
	"github.com/t73liu/trading-bot/lib/utils"
	"github.com/t73liu/trading-bot/trader/middleware"
	"log"
	"net/http"
	"time"
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

func (h *Handlers) getTradableStocks(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
}

func (h *Handlers) getStockInfo(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
}

func (h *Handlers) getStockCandles(w http.ResponseWriter, _ *http.Request, p httprouter.Params) {
	symbol := p.ByName("symbol")
	location := utils.GetNYSELocation()
	endTime := time.Now().In(location)
	startTime := analyze.GetMidnight(endTime, location)
	candles, err := traderdb.GetStockCandles(h.db, symbol, startTime, endTime)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if err = json.NewEncoder(w).Encode(candles); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handlers) getStockIndicators(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
}

func (h *Handlers) AddRoutes(router *httprouter.Router) {
	router.GET(
		"/api/stocks",
		middleware.LogResponseTime(h.getTradableStocks, h.logger),
	)
	router.GET(
		"/api/stocks/:symbol",
		middleware.LogResponseTime(h.getStockInfo, h.logger),
	)
	router.GET(
		"/api/stocks/:symbol/candles",
		middleware.LogResponseTime(h.getStockCandles, h.logger),
	)
	router.GET(
		"/api/stocks/:symbol/indicators",
		middleware.LogResponseTime(h.getStockIndicators, h.logger),
	)
}
