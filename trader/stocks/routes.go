package stocks

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
	"tradingbot/lib/candle"
	"tradingbot/lib/finviz"
	"tradingbot/lib/options"
	"tradingbot/lib/polygon"
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
	polygonClient *polygon.Client
	yahooClient   *yahoo.Client
	finvizClient  *finviz.Client
	optionsClient *options.Client
}

func NewHandlers(logger *log.Logger, db *pgxpool.Pool, polygonClient *polygon.Client, yahooClient *yahoo.Client, finvizClient *finviz.Client, optionsClient *options.Client) *Handlers {
	return &Handlers{
		logger:        logger,
		db:            db,
		polygonClient: polygonClient,
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
	// TODO Handle error once Polygon issues are resolved
	info, _ := h.polygonClient.GetTickerDetails(symbol)
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
		SimilarStocks: info.Similar,
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
			OpenMicros:  utils.DollarsToMicros(bar.Open),
			HighMicros:  utils.DollarsToMicros(bar.High),
			LowMicros:   utils.DollarsToMicros(bar.Low),
			CloseMicros: utils.DollarsToMicros(bar.Close),
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

// TODO Use once Polygon issues are resolved
func (h *Handlers) getStockNews(w http.ResponseWriter, r *http.Request) {
	symbol := getSymbol(r)
	news, err := h.polygonClient.GetTickerNews(symbol, 10, 1)
	if err != nil {
		utils.JSONError(w, err)
		return
	}
	utils.JSONResponse(w, news)
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

func (h *Handlers) getGapStocks(w http.ResponseWriter, _ *http.Request) {
	stocksBySymbol, err := traderdb.GetTradableStocksBySymbol(h.db)
	if err != nil {
		utils.JSONError(w, err)
		return
	}
	snapshots := make(map[string][]Snapshot)
	gains, err := h.polygonClient.GetMovers(true)
	if err != nil {
		utils.JSONError(w, err)
		return
	}
	gapUpStocks := make([]Snapshot, 0, len(gains))
	for _, gain := range gains {
		if stock, ok := stocksBySymbol[gain.Ticker]; ok {
			gapUpStocks = append(gapUpStocks, Snapshot{
				Company:        stock.Company,
				Symbol:         gain.Ticker,
				Change:         gain.Change,
				ChangePercent:  gain.ChangePercent,
				PreviousVolume: int64(gain.PrevDay.Volume),
				PreviousClose:  gain.PrevDay.Close,
				UpdatedAt:      utils.ConvertUnixNanosToTime(gain.UpdatedAtUnixNanos),
			})
		}
	}
	snapshots["gapUp"] = gapUpStocks
	losses, err := h.polygonClient.GetMovers(false)
	if err != nil {
		utils.JSONError(w, err)
		return
	}
	gapDownStocks := make([]Snapshot, 0, len(losses))
	for _, loss := range losses {
		if stock, ok := stocksBySymbol[loss.Ticker]; ok {
			gapDownStocks = append(gapDownStocks, Snapshot{
				Company:        stock.Company,
				Symbol:         loss.Ticker,
				Change:         loss.Change,
				ChangePercent:  loss.ChangePercent,
				PreviousVolume: int64(loss.PrevDay.Volume),
				PreviousClose:  loss.PrevDay.Close,
				UpdatedAt:      utils.ConvertUnixNanosToTime(loss.UpdatedAtUnixNanos),
			})
		}
	}
	snapshots["gapDown"] = gapDownStocks
	utils.JSONResponse(w, snapshots)
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
	router.HandleFunc("/gaps", h.getGapStocks).Methods("GET")
	router.HandleFunc("/screened", h.getScreenedStocks).Methods("GET")
	router.HandleFunc("/{symbol}", h.getStockInfo).Methods("GET")
	router.HandleFunc("/{symbol}/charts", h.getStockCharts).Methods("GET")
	router.HandleFunc("/{symbol}/news", h.getStockNews).Methods("GET")
	router.HandleFunc("/{symbol}/options", h.getStockOptions).Methods("GET")
}

func getSymbol(r *http.Request) string {
	vars := mux.Vars(r)
	return strings.ToUpper(vars["symbol"])
}
