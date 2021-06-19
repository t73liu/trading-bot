package main

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/t73liu/tradingbot/lib/alpaca"
	"github.com/t73liu/tradingbot/lib/candle"
	analyze "github.com/t73liu/tradingbot/lib/technical-analysis"
	"github.com/t73liu/tradingbot/lib/traderdb"
	"github.com/t73liu/tradingbot/lib/utils"

	"github.com/gorilla/mux"
)

type Detail struct {
	Price         float64           `json:"price"`
	Company       string            `json:"company"`
	Website       *utils.NullString `json:"website"`
	Description   *utils.NullString `json:"description"`
	Sector        *utils.NullString `json:"sector"`
	Industry      *utils.NullString `json:"industry"`
	Country       *utils.NullString `json:"country"`
	AverageVolume int64             `json:"averageVolume"`
	MarketCap     int64             `json:"marketCap"`
	SimilarStocks []string          `json:"similarStocks"`
	Shortable     bool              `json:"shortable"`
	Marginable    bool              `json:"marginable"`
	News          interface{}       `json:"news"`
}

type Snapshot struct {
	Symbol         string    `json:"symbol"`
	Company        string    `json:"company"`
	Change         float64   `json:"change"`
	ChangePercent  float64   `json:"changePercent"`
	PreviousVolume int64     `json:"previousVolume"`
	PreviousClose  float64   `json:"previousClose"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

func (t *trader) getTradableStocks(w http.ResponseWriter, _ *http.Request) {
	stocks, err := traderdb.GetTradableStocks(t.db)
	if err != nil {
		utils.JSONError(w, err)
		return
	}
	utils.JSONResponse(w, stocks)
}

func (t *trader) getStockInfo(w http.ResponseWriter, r *http.Request) {
	symbol := getSymbol(r)
	stock, err := traderdb.GetTradableStock(t.db, symbol)
	if err != nil {
		utils.JSONError(w, err)
		return
	}
	details, err := t.yahooClient.GetStock(symbol)
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

func (t *trader) getStockCharts(w http.ResponseWriter, r *http.Request) {
	symbol := getSymbol(r)
	location := utils.GetNYSELocation()
	endTime := time.Now().In(location)
	if !utils.IsWeekday(endTime) {
		endTime = utils.GetLastWeekday(endTime)
	}
	startTime := utils.GetMidnight(endTime, location)
	query := r.URL.Query()

	charts := make(map[string]interface{})
	candlesResponse, err := t.alpacaClient.GetSymbolCandles(symbol, alpaca.CandleQueryParams{
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
	case "5min":
		candles, _ = candle.CompressCandles(candles, 5, "minute", location)
	case "10min":
		candles, _ = candle.CompressCandles(candles, 10, "minute", location)
	case "30min":
		candles, _ = candle.CompressCandles(candles, 30, "minute", location)
	case "1hour":
		candles, _ = candle.CompressCandles(candles, 1, "hour", location)
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

func (t *trader) getStockOptions(w http.ResponseWriter, r *http.Request) {
	symbol := getSymbol(r)
	stockOptions, err := t.optionsClient.GetOptions(symbol)
	if err != nil {
		utils.JSONError(w, err)
		return
	}
	utils.JSONResponse(w, stockOptions)
}

// Optionable, 2B+ market cap, Up 1% from open, Over 2M volume, Over 1 relative Volume
const screenerQuery = "v=111&f=cap_midover,sh_curvol_o2000,sh_opt_option,sh_relvol_o1,ta_changeopen_u1&ft=4&o=-change"

func (t *trader) getScreenedStocks(w http.ResponseWriter, _ *http.Request) {
	screenedStocks, err := t.finvizClient.ScreenStocksOverview(screenerQuery)
	if err != nil {
		utils.JSONError(w, err)
		return
	}
	utils.JSONResponse(w, screenedStocks)
}

func (t *trader) AddStockRoutes(router *mux.Router) {
	router.HandleFunc("", t.getTradableStocks).Methods("GET")
	router.HandleFunc("/screened", t.getScreenedStocks).Methods("GET")
	router.HandleFunc("/{symbol}", t.getStockInfo).Methods("GET")
	router.HandleFunc("/{symbol}/charts", t.getStockCharts).Methods("GET")
	router.HandleFunc("/{symbol}/options", t.getStockOptions).Methods("GET")
}

func getSymbol(r *http.Request) string {
	vars := mux.Vars(r)
	return strings.ToUpper(vars["symbol"])
}
