package stock

import (
	"github.com/julienschmidt/httprouter"
	"github.com/t73liu/trading-bot/trader/middleware"
	"log"
	"net/http"
)

type Handlers struct {
	logger *log.Logger
}

func NewHandlers(logger *log.Logger) *Handlers {
	return &Handlers{
		logger: logger,
	}
}

func (h *Handlers) getTradableStocks(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
}

func (h *Handlers) getStockInfo(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
}

func (h *Handlers) getStockCandles(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
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
