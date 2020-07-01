package traderdb

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

const watchlistsQuery = `
SELECT wl.id, wl.name, s.id as stock_id, s.symbol FROM watchlists wl
INNER JOIN watchlist_stocks wls ON wl.id = wls.watchlist_id
INNER JOIN stocks s ON wls.stock_id = s.id
WHERE wl.user_id = $1
`

type StockSymbol struct {
	Id     int    `json:"id"`
	Symbol string `json:"symbol"`
}

type Watchlist struct {
	Id     int           `json:"id"`
	Name   string        `json:"name"`
	Stocks []StockSymbol `json:"stocks"`
}

func GetWatchlistsByUserId(db *pgxpool.Pool, userId int) (watchlists []Watchlist, err error) {
	rows, err := db.Query(context.Background(), watchlistsQuery, userId)
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

func GetWatchlistStocksByUserId(db *pgxpool.Pool, userId int) (stocks []StockSymbol, err error) {
	watchlists, err := GetWatchlistsByUserId(db, userId)
	if err != nil {
		return stocks, err
	}
	stocksById := make(map[int]struct{})
	for _, watchlist := range watchlists {
		for _, stock := range watchlist.Stocks {
			if _, ok := stocksById[stock.Id]; !ok {
				stocksById[stock.Id] = struct{}{}
				stocks = append(stocks, stock)
			}
		}
	}
	return stocks, nil
}
