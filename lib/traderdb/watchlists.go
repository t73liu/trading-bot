package traderdb

import (
	"context"

	"github.com/t73liu/tradingbot/lib/utils"

	"github.com/jackc/pgx/v4"
)

type Watchlist struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	StockIDs []int  `json:"stockIDs"`
}

const watchlistsQuery = `
SELECT wl.id, wl.name, wls.stock_id as stock_id FROM watchlists wl
LEFT JOIN watchlist_stocks wls ON wl.id = wls.watchlist_id
WHERE user_id = $1
ORDER BY wl.id
`

func GetWatchlistsWithUserID(db PGConnection, userID int) (watchlists []Watchlist, err error) {
	rows, err := db.Query(context.Background(), watchlistsQuery, userID)
	if err != nil {
		return watchlists, err
	}
	defer rows.Close()

	watchlistsByID := make(map[int]Watchlist)
	for rows.Next() {
		var watchlistID int
		var name string
		var stockID utils.NullInt64
		if err = rows.Scan(&watchlistID, &name, &stockID); err != nil {
			return watchlists, err
		}
		if _, ok := watchlistsByID[watchlistID]; !ok {
			watchlistsByID[watchlistID] = Watchlist{
				ID:       watchlistID,
				Name:     name,
				StockIDs: make([]int, 0),
			}
		}
		watchlist := watchlistsByID[watchlistID]
		if stockID.Valid {
			watchlist.StockIDs = append(watchlist.StockIDs, int(stockID.Int64))
		}
		watchlistsByID[watchlistID] = watchlist
	}

	if rows.Err() != nil {
		return watchlists, rows.Err()
	}

	watchlists = make([]Watchlist, 0, len(watchlistsByID))
	for _, watchlist := range watchlistsByID {
		watchlists = append(watchlists, watchlist)
	}
	return watchlists, nil
}

const watchlistExistsQuery = `
SELECT EXISTS(SELECT 1 FROM watchlists WHERE id = $1 AND user_id = $2)
`

func HasWatchlistWithIDAndUserID(db PGConnection, watchlistID int, userID int) (bool, error) {
	var exists bool
	err := db.QueryRow(
		context.Background(),
		watchlistExistsQuery,
		watchlistID,
		userID,
	).Scan(&exists)
	if err != nil {
		return exists, err
	}
	return exists, err
}

func CreateWatchlist(db PGConnection, userID int, watchlistName string, stockIDs []int) (watchlistID int, err error) {
	tx, err := db.Begin(context.Background())
	if err != nil {
		return watchlistID, err
	}
	defer tx.Rollback(context.Background())

	err = tx.QueryRow(
		context.Background(),
		"INSERT INTO watchlists (user_id, name) VALUES ($1, $2) RETURNING id",
		userID,
		watchlistName,
	).Scan(&watchlistID)
	if err != nil {
		return watchlistID, err
	}

	if len(stockIDs) > 0 {
		rows := make([][]interface{}, 0, len(stockIDs))
		for _, stockID := range stockIDs {
			rows = append(rows, []interface{}{watchlistID, stockID})
		}
		_, err = tx.CopyFrom(
			context.Background(),
			pgx.Identifier{"watchlist_stocks"},
			[]string{"watchlist_id", "stock_id"},
			pgx.CopyFromRows(rows),
		)
		if err != nil {
			return watchlistID, err
		}
	}

	if err = tx.Commit(context.Background()); err != nil {
		return watchlistID, err
	}

	return watchlistID, err
}

func UpdateWatchlistWithID(db PGConnection, watchlistID int, watchlistName string, stockIDs []int) error {
	tx, err := db.Begin(context.Background())
	if err != nil {
		return err
	}
	defer tx.Rollback(context.Background())

	_, err = tx.Exec(
		context.Background(),
		"UPDATE watchlists SET name = $1 WHERE id = $2",
		watchlistName,
		watchlistID,
	)
	if err != nil {
		return err
	}

	_, err = tx.Exec(
		context.Background(),
		"DELETE FROM watchlist_stocks WHERE watchlist_id = $1",
		watchlistID,
	)
	if err != nil {
		return err
	}

	if len(stockIDs) > 0 {
		rows := make([][]interface{}, 0, len(stockIDs))
		for _, stockID := range stockIDs {
			rows = append(rows, []interface{}{watchlistID, stockID})
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

func DeleteWatchlistWithID(db PGConnection, watchlistID int) error {
	tx, err := db.Begin(context.Background())
	if err != nil {
		return err
	}
	defer tx.Rollback(context.Background())

	_, err = tx.Exec(
		context.Background(),
		"DELETE FROM watchlist_stocks WHERE watchlist_id = $1",
		watchlistID,
	)
	if err != nil {
		return err
	}

	_, err = tx.Exec(
		context.Background(),
		"DELETE FROM watchlists WHERE id = $1",
		watchlistID,
	)
	if err != nil {
		return err
	}

	if err = tx.Commit(context.Background()); err != nil {
		return err
	}

	return nil
}

func GetAllWatchlistStocks(db PGConnection) ([]Stock, error) {
	rows, err := db.Query(context.Background(), "SELECT DISTINCT stock_id FROM watchlist_stocks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stockIDs []int
	var stockID int
	for rows.Next() {
		if err = rows.Scan(&stockID); err != nil {
			return nil, err
		}
		stockIDs = append(stockIDs, stockID)
	}
	return GetStocksWithIDs(db, stockIDs)
}

func GetWatchlistStocksWithUserID(db PGConnection, userID int) (stocks []Stock, err error) {
	watchlists, err := GetWatchlistsWithUserID(db, userID)
	if err != nil {
		return stocks, err
	}
	stocksByID := make(map[int]struct{})
	var stockIDs []int
	for _, watchlist := range watchlists {
		for _, stockID := range watchlist.StockIDs {
			if _, ok := stocksByID[stockID]; !ok {
				stocksByID[stockID] = struct{}{}
				stockIDs = append(stockIDs, stockID)
			}
		}
	}
	return GetStocksWithIDs(db, stockIDs)
}
