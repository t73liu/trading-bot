package stock

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strconv"
	"time"
	"tradingbot/lib/candle"
	"tradingbot/lib/polygon"
	"tradingbot/lib/traderdb"
	"tradingbot/lib/utils"
	"tradingbot/lib/yahoo-finance"
	"tradingbot/trader/middleware"
)

type Handlers struct {
	logger        *log.Logger
	db            *pgxpool.Pool
	polygonClient *polygon.Client
	yahooClient   *yahoo.Client
}

func NewHandlers(logger *log.Logger, db *pgxpool.Pool, polygonClient *polygon.Client, yahooClient *yahoo.Client) *Handlers {
	return &Handlers{
		logger:        logger,
		db:            db,
		polygonClient: polygonClient,
		yahooClient:   yahooClient,
	}
}

func (h *Handlers) getTradableStocks(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	stocks, err := traderdb.GetTradableStocks(h.db)
	if err != nil {
		utils.JSONError(w, err)
		return
	}
	utils.JSONResponse(w, stocks)
}

func (h *Handlers) getStockInfo(w http.ResponseWriter, _ *http.Request, p httprouter.Params) {
	symbol := p.ByName("symbol")
	stock, err := traderdb.GetTradableStock(h.db, symbol)
	if err != nil {
		utils.JSONError(w, err)
		return
	}
	// TODO Handle error once Polygon issues are resolved
	info, _ := h.polygonClient.GetTickerDetails(symbol)
	details, err := h.yahooClient.GetStock(symbol)
	if err != nil {
		utils.JSONError(w, err)
		return
	}
	utils.JSONResponse(w, Detail{
		Company:       details.Company,
		Website:       details.Website,
		Description:   details.Description,
		Sector:        details.Sector,
		Industry:      details.Industry,
		Country:       details.Country,
		AverageVolume: details.AverageVolume,
		MarketCap:     details.MarketCap,
		SimilarStocks: info.Similar,
		Shortable:     stock.Shortable,
		Marginable:    stock.Marginable,
		News:          details.News,
	})
}

func (h *Handlers) getStockCandles(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	symbol := p.ByName("symbol")
	location := utils.GetNYSELocation()
	endTime := time.Now().In(location)
	if !utils.IsWeekday(endTime) {
		endTime = utils.GetLastWeekday(endTime)
	}
	startTime := utils.GetMidnight(endTime, location)
	query := r.URL.Query()
	showExtendedHours, _ := strconv.ParseBool(query.Get("showExtendedHours"))
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
			utils.JSONError(w, err)
		}
		candles := make([]candle.Candle, 0, len(bars))
		for _, bar := range bars {
			candles = append(candles, candle.Candle{
				OpenedAt:    utils.ConvertUnixMillisToTime(bar.StartAtUnixMillis).In(location),
				Volume:      int64(bar.Volume),
				OpenMicros:  candle.DollarsToMicros(bar.Open),
				HighMicros:  candle.DollarsToMicros(bar.High),
				LowMicros:   candle.DollarsToMicros(bar.Low),
				CloseMicros: candle.DollarsToMicros(bar.Close),
			})
		}
		if !showExtendedHours {
			candles = candle.FilterTradingHourCandles(candles)
		}
		utils.JSONResponse(w, candle.FillMinuteCandles(candles))
	} else {
		candles, err := traderdb.GetStockCandles(h.db, symbol, startTime, endTime)
		if err != nil {
			utils.JSONError(w, err)
		}
		utils.JSONResponse(w, candles)
	}
}

// TODO Use once Polygon issues are resolved
func (h *Handlers) getStockNews(w http.ResponseWriter, _ *http.Request, p httprouter.Params) {
	symbol := p.ByName("symbol")
	news, err := h.polygonClient.GetTickerNews(symbol, 10, 1)
	if err != nil {
		utils.JSONError(w, err)
		return
	}
	utils.JSONResponse(w, news)
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
		"/api/stocks/:symbol/news",
		middleware.LogResponseTime(h.getStockNews, h.logger),
	)
	router.GET(
		"/api/stocks/:symbol/indicators",
		middleware.LogResponseTime(h.getStockIndicators, h.logger),
	)
}
