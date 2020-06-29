package account

import (
	"context"
	"encoding/json"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
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

const watchlistsQuery = `
SELECT wl.id, wl.name, s.id as stock_id, s.symbol FROM watchlists wl
INNER JOIN watchlist_stocks wls ON wl.id = wls.watchlist_id
INNER JOIN stocks s ON wls.stock_id = s.id
WHERE wl.user_id = $1
`

func (h *Handlers) handleGetWatchlists(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	watchlists, err := h.getWatchlists(1)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		if err = json.NewEncoder(w).Encode(watchlists); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func (h *Handlers) getWatchlists(userId int) (watchlists []Watchlist, err error) {
	rows, err := h.db.Query(context.Background(), watchlistsQuery, userId)
	if err != nil {
		return watchlists, err
	}
	defer rows.Close()

	watchlistsById := make(map[int]Watchlist)
	for rows.Next() {
		var watchlistId int
		var name string
		var stockId int
		var symbol string
		err = rows.Scan(&watchlistId, &name, &stockId, &symbol)
		if err != nil {
			return watchlists, err
		}
		_, ok := watchlistsById[watchlistId]
		if !ok {
			watchlistsById[watchlistId] = Watchlist{
				Id:     watchlistId,
				Name:   name,
				Stocks: make([]StockSymbol, 0),
			}
		}
		watchlist := watchlistsById[watchlistId]
		watchlist.Stocks = append(
			watchlist.Stocks,
			StockSymbol{
				Id:     stockId,
				Symbol: symbol,
			},
		)
		watchlistsById[watchlistId] = watchlist
	}

	if rows.Err() != nil {
		return watchlists, rows.Err()
	}

	for _, watchlist := range watchlistsById {
		watchlists = append(watchlists, watchlist)
	}
	return watchlists, nil
}

func (h *Handlers) AddRoutes(router *httprouter.Router) {
	router.GET("/api/account/watchlists", h.handleGetWatchlists)
	//router.POST("/api/account/watchlists", h.handleUpdateWatchlists)
}
