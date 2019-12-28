package stock

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/t73liu/trading-bot/trader/candle"
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

func (h *Handlers) getCandles(w http.ResponseWriter, _ *http.Request, params httprouter.Params) {
	symbol := params.ByName("symbol")
	fmt.Println(symbol)
	data := candle.Candle{}
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handlers) getInfo(w http.ResponseWriter, _ *http.Request, params httprouter.Params) {
	symbol := params.ByName("symbol")
	fmt.Println(symbol)
	data := candle.Candle{}
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handlers) AddRoutes(router *httprouter.Router) {
	router.GET("/api/stocks/:symbol/candles", h.getCandles)
	router.GET("/api/stocks/:symbol/info", h.getInfo)
	// router.GET("/api/stocks", h.getAllStocks)
	// router.PATCH/POST add to watch list, update price, ...
}
