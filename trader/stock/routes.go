package stock

import (
	"encoding/json"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"time"
	"tradingbot/lib/candle"
	"tradingbot/lib/polygon"
	"tradingbot/lib/traderdb"
	"tradingbot/lib/utils"
	"tradingbot/trader/middleware"
)

type Handlers struct {
	logger        *log.Logger
	db            *pgxpool.Pool
	polygonClient *polygon.Client
}

func NewHandlers(logger *log.Logger, db *pgxpool.Pool, polygonClient *polygon.Client) *Handlers {
	return &Handlers{
		logger:        logger,
		db:            db,
		polygonClient: polygonClient,
	}
}

func (h *Handlers) getTradableStocks(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
}

func (h *Handlers) getStockInfo(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
}

func (h *Handlers) getStockCandles(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	symbol := p.ByName("symbol")
	location := utils.GetNYSELocation()
	endTime := time.Now().In(location)
	startTime := utils.GetMidnight(endTime, location)
	query := r.URL.Query()
	if query.Get("interval") == "intraday" {
		bars, err := h.polygonClient.GetTickerBars(polygon.TickerBarsQueryParams{
			Ticker:       symbol,
			TimeInterval: 1,
			TimeUnit:     "minute",
			StartDate:    startTime,
			EndDate:      endTime,
			Unadjusted:   false,
			Sort:         "asc",
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		candles := make([]candle.Candle, 0, len(bars))
		for _, bar := range bars {
			candles = append(candles, candle.Candle{
				OpenedAt:    utils.ConvertUnixSecondsToTime(bar.StartAtUnixMillis / 1000),
				Volume:      int64(bar.Volume),
				OpenMicros:  candle.DollarsToMicros(bar.Open),
				HighMicros:  candle.DollarsToMicros(bar.High),
				LowMicros:   candle.DollarsToMicros(bar.Low),
				CloseMicros: candle.DollarsToMicros(bar.Close),
			})
		}
		if err = json.NewEncoder(w).Encode(candles); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		candles, err := traderdb.GetStockCandles(h.db, symbol, startTime, endTime)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		if err = json.NewEncoder(w).Encode(candles); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
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
