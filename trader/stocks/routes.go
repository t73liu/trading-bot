package stocks

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
	"tradingbot/lib/alpaca"
	"tradingbot/lib/candle"
	"tradingbot/lib/finviz"
	"tradingbot/lib/options"
	analyze "tradingbot/lib/technical-analysis"
	"tradingbot/lib/traderdb"
	"tradingbot/lib/utils"
	"tradingbot/lib/yahoo-finance"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Handlers struct {
	logger        *log.Logger
	db            *pgxpool.Pool
	alpacaClient  *alpaca.Client
	yahooClient   *yahoo.Client
	finvizClient  *finviz.Client
	optionsClient *options.Client
}

func NewHandlers(logger *log.Logger, db *pgxpool.Pool, alpacaClient *alpaca.Client, yahooClient *yahoo.Client, finvizClient *finviz.Client, optionsClient *options.Client) *Handlers {
	return &Handlers{
		logger:        logger,
		db:            db,
		alpacaClient:  alpacaClient,
		yahooClient:   yahooClient,
		finvizClient:  finvizClient,
		optionsClient: optionsClient,
	}
}

func (h *Handlers) getTradableStocks(w http.ResponseWriter, _ *http.Request) {
	stocks, err := traderdb.GetTradableStocks(h.db)
	if err != nil {
		utils.JSONError(w, err)
		return
	}
	utils.JSONResponse(w, stocks)
}

func (h *Handlers) getStockInfo(w http.ResponseWriter, r *http.Request) {
	symbol := getSymbol(r)
	stock, err := traderdb.GetTradableStock(h.db, symbol)
	if err != nil {
		utils.JSONError(w, err)
		return
	}
	details, err := h.yahooClient.GetStock(symbol)
	if err != nil {
		utils.JSONError(w, err)
		return
	}
	utils.JSONResponse(w, Detail{
		Price:         details.Price,
		Company:       details.Company,
		Website:       details.Website,
		Description:   details.Description,
		Sector:        details.Sector,
		Industry:      details.Industry,
		Country:       details.Country,
		AverageVolume: details.AverageVolume,
		MarketCap:     details.MarketCap,
		Shortable:     stock.Shortable,
		Marginable:    stock.Marginable,
		News:          details.News,
	})
}

func (h *Handlers) getStockCharts(w http.ResponseWriter, r *http.Request) {
	symbol := getSymbol(r)
	location := utils.GetNYSELocation()
	endTime := time.Now().In(location)
	if !utils.IsWeekday(endTime) {
		endTime = utils.GetLastWeekday(endTime)
	}
	startTime := utils.GetMidnight(endTime, location)
	query := r.URL.Query()

	charts := make(map[string]interface{})
	candlesResponse, err := h.alpacaClient.GetSymbolCandles(symbol, alpaca.CandleQueryParams{
		Limit:      10000,
		CandleSize: alpaca.OneMin,
		StartTime:  startTime,
		EndTime:    endTime,
	})
	if err != nil {
		utils.JSONError(w, err)
	}
	candles := make([]candle.Candle, 0, len(candlesResponse.Candles))
	for _, bar := range candlesResponse.Candles {
		candles = append(candles, candle.Candle{
			OpenedAt:    bar.StartAt.In(location),
			Volume:      int64(bar.Volume),
			OpenMicros:  utils.DollarsToMicros(float64(bar.Open)),
			HighMicros:  utils.DollarsToMicros(float64(bar.High)),
			LowMicros:   utils.DollarsToMicros(float64(bar.Low)),
			CloseMicros: utils.DollarsToMicros(float64(bar.Close)),
		})
	}

	// Format Candles
	showExtendedHours, _ := strconv.ParseBool(query.Get("showExtendedHours"))
	if !showExtendedHours {
		candles = candle.FilterTradingHourCandles(candles)
	}
	candles = candle.FillMinuteCandles(candles)
	candleSize := query.Get("candleSize")
	switch candleSize {
	case "3min":
		candles, _ = candle.CompressCandles(candles, 3, "minute", location)
		break
	case "5min":
		candles, _ = candle.CompressCandles(candles, 5, "minute", location)
		break
	case "10min":
		candles, _ = candle.CompressCandles(candles, 10, "minute", location)
		break
	case "30min":
		candles, _ = candle.CompressCandles(candles, 30, "minute", location)
		break
	case "1hour":
		candles, _ = candle.CompressCandles(candles, 1, "hour", location)
		break
	}

	// Add candles to response
	closingPrices := candle.GetClosingPrices(candles)
	volumes := candle.GetVolumes(candles)
	charts["candles"] = candles
	charts["volume"] = volumes
	charts["currentVolume"] = utils.Sum(volumes...)
	// Add indicators to response
	charts["ema"] = analyze.EMA(closingPrices, 9)
	charts["vwap"] = analyze.VWAP(candles)
	charts["ttmSqueeze"] = analyze.TTMSqueeze(candles)
	charts["macd"] = analyze.StandardMACD(closingPrices)
	charts["rsi"] = analyze.RSI(closingPrices, 14)
	utils.JSONResponse(w, charts)
}

func (h *Handlers) getStockOptions(w http.ResponseWriter, r *http.Request) {
	symbol := getSymbol(r)
	stockOptions, err := h.optionsClient.GetOptions(symbol)
	if err != nil {
		utils.JSONError(w, err)
		return
	}
	utils.JSONResponse(w, stockOptions)
}

// Optionable, 2B+ market cap, Up 1% from open, Over 2M volume, Over 1 relative Volume
const screenerQuery = "v=111&f=cap_midover,sh_curvol_o2000,sh_opt_option,sh_relvol_o1,ta_changeopen_u1&ft=4&o=-change"

func (h *Handlers) getScreenedStocks(w http.ResponseWriter, _ *http.Request) {
	screenedStocks, err := h.finvizClient.ScreenStocksOverview(screenerQuery)
	if err != nil {
		utils.JSONError(w, err)
		return
	}
	utils.JSONResponse(w, screenedStocks)
}

func (h *Handlers) AddRoutes(router *mux.Router) {
	router.HandleFunc("", h.getTradableStocks).Methods("GET")
	router.HandleFunc("/screened", h.getScreenedStocks).Methods("GET")
	router.HandleFunc("/{symbol}", h.getStockInfo).Methods("GET")
	router.HandleFunc("/{symbol}/charts", h.getStockCharts).Methods("GET")
	router.HandleFunc("/{symbol}/options", h.getStockOptions).Methods("GET")
}

func getSymbol(r *http.Request) string {
	vars := mux.Vars(r)
	return strings.ToUpper(vars["symbol"])
}
