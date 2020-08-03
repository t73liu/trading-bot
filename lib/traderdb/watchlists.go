package traderdb

import (
	"context"
	"github.com/jackc/pgx/v4"
)

type Watchlist struct {
	Id     int     `json:"id"`
	Name   string  `json:"name"`
	Stocks []Stock `json:"stocks"`
}

const watchlistsQuery = `
SELECT wl.id, wl.name, s.id as stock_id, s.symbol, s.company FROM watchlists wl
INNER JOIN watchlist_stocks wls ON wl.id = wls.watchlist_id
INNER JOIN stocks s ON wls.stock_id = s.id
WHERE wl.user_id = $1
`

func GetWatchlistsByUserId(db PGConnection, userId int) (watchlists []Watchlist, err error) {
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
		var company string
		if err = rows.Scan(&watchlistId, &name, &stockId, &symbol, &company); err != nil {
			return watchlists, err
		}
		_, ok := watchlistsById[watchlistId]
		if !ok {
			watchlistsById[watchlistId] = Watchlist{
				Id:     watchlistId,
				Name:   name,
				Stocks: make([]Stock, 0),
			}
		}
		watchlist := watchlistsById[watchlistId]
		watchlist.Stocks = append(
			watchlist.Stocks,
			Stock{
				Id:      stockId,
				Symbol:  symbol,
				Company: company,
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

func GetWatchlistStocksByUserId(db PGConnection, userId int) (stocks []Stock, err error) {
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

const watchlistExistsQuery = `
SELECT EXISTS(SELECT 1 FROM watchlists WHERE id = $1 AND user_id = $2)
`

func HasWatchlistWithIdAndUserId(db PGConnection, watchlistId int, userId int) (bool, error) {
	var exists bool
	err := db.QueryRow(
		context.Background(),
		watchlistExistsQuery,
		watchlistId,
		userId,
	).Scan(&exists)
	if err != nil {
		return exists, err
	}
	return exists, err
}

const watchlistStocksQuery = `
SELECT s.id, s.symbol, s.company FROM watchlist_stocks wls
INNER JOIN stocks s ON s.id = wls.stock_id
WHERE wls.watchlist_id = $1
`

func GetWatchlistById(db PGConnection, watchlistId int) (watchlist Watchlist, err error) {
	var watchlistName string
	err = db.QueryRow(
		context.Background(),
		"SELECT name FROM watchlists WHERE id = $1",
		watchlistId,
	).Scan(&watchlistName)
	if err != nil {
		return watchlist, err
	}

	rows, err := db.Query(context.Background(), watchlistStocksQuery, watchlistId)
	if err != nil {
		return watchlist, err
	}
	defer rows.Close()

	stocks := make([]Stock, 0)
	for rows.Next() {
		var stockId int
		var symbol string
		var company string
		if err = rows.Scan(&stockId, &symbol, &company); err != nil {
			return watchlist, err
		}
		stocks = append(stocks, Stock{
			Id:      stockId,
			Symbol:  symbol,
			Company: company,
		})
	}

	if rows.Err() != nil {
		return watchlist, rows.Err()
	}

	watchlist.Name = watchlistName
	watchlist.Id = watchlistId
	watchlist.Stocks = stocks
	return watchlist, nil
}

func CreateWatchlist(db PGConnection, userId int, watchlistName string, stockIds []int) (watchlistId int, err error) {
	tx, err := db.Begin(context.Background())
	if err != nil {
		return watchlistId, err
	}
	defer tx.Rollback(context.Background())

	err = tx.QueryRow(
		context.Background(),
		"INSERT INTO watchlists (user_id, name) VALUES ($1, $2) RETURNING id",
		userId,
		watchlistName,
	).Scan(&watchlistId)
	if err != nil {
		return watchlistId, err
	}

	if len(stockIds) > 0 {
		rows := make([][]interface{}, 0, len(stockIds))
		for _, stockId := range stockIds {
			rows = append(rows, []interface{}{watchlistId, stockId})
		}
		_, err = tx.CopyFrom(
			context.Background(),
			pgx.Identifier{"watchlist_stocks"},
			[]string{"watchlist_id", "stock_id"},
			pgx.CopyFromRows(rows),
		)
		if err != nil {
			return watchlistId, err
		}
	}

	if err = tx.Commit(context.Background()); err != nil {
		return watchlistId, err
	}

	return watchlistId, err
}

func UpdateWatchlistById(db PGConnection, watchlistId int, watchlistName string, stockIds []int) error {
	tx, err := db.Begin(context.Background())
	if err != nil {
		return err
	}
	defer tx.Rollback(context.Background())

	_, err = tx.Exec(
		context.Background(),
		"UPDATE watchlists SET name = $1 WHERE id = $2",
		watchlistName,
		watchlistId,
	)
	if err != nil {
		return err
	}

	_, err = tx.Exec(
		context.Background(),
		"DELETE FROM watchlist_stocks WHERE watchlist_id = $1",
		watchlistId,
	)
	if err != nil {
		return err
	}

	if len(stockIds) > 0 {
		rows := make([][]interface{}, 0, len(stockIds))
		for _, stockId := range stockIds {
			rows = append(rows, []interface{}{watchlistId, stockId})
		}
		_, err = tx.CopyFrom(
			context.Background(),
			pgx.Identifier{"watchlist_stocks"},
			[]string{"watchlist_id", "stock_id"},
			pgx.CopyFromRows(rows),
		)
		if err != nil {
			return err
		}
	}

	if err = tx.Commit(context.Background()); err != nil {
		return err
	}

	return nil
}

func DeleteWatchlistById(db PGConnection, watchlistId int) error {
	tx, err := db.Begin(context.Background())
	if err != nil {
		return err
	}
	defer tx.Rollback(context.Background())

	_, err = tx.Exec(
		context.Background(),
		"DELETE FROM watchlist_stocks WHERE watchlist_id = $1",
		watchlistId,
	)
	if err != nil {
		return err
	}

	_, err = tx.Exec(
		context.Background(),
		"DELETE FROM watchlists WHERE id = $1",
		watchlistId,
	)
	if err != nil {
		return err
	}

	if err = tx.Commit(context.Background()); err != nil {
		return err
	}

	return nil
}
