package traderdb

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v4"
)

type Stock struct {
	Id         int    `json:"id"`
	Symbol     string `json:"symbol"`
	Company    string `json:"company"`
	Exchange   string `json:"exchange,omitempty"`
	Tradable   bool   `json:"tradable"`
	Shortable  bool   `json:"shortable"`
	Marginable bool   `json:"marginable"`
}

func (s *Stock) Equal(other Stock) bool {
	return s.Symbol == other.Symbol &&
		s.Company == other.Company &&
		s.Exchange == other.Exchange &&
		s.Tradable == other.Tradable &&
		s.Shortable == other.Shortable &&
		s.Marginable == other.Marginable
}

const stocksQuery = `
SELECT id, symbol, company, exchange, tradable, shortable, marginable
FROM stocks
%s
ORDER BY symbol
`

func getStocks(db PGConnection, query string) (stocks []Stock, err error) {
	rows, err := db.Query(context.Background(), query)
	if err != nil {
		return stocks, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var symbol string
		var company string
		var exchange string
		var tradable bool
		var shortable bool
		var marginable bool
		if err = rows.Scan(&id, &symbol, &company, &exchange, &tradable, &shortable, &marginable); err != nil {
			return stocks, err
		}
		stocks = append(stocks, Stock{
			Id:         id,
			Symbol:     symbol,
			Company:    company,
			Exchange:   exchange,
			Tradable:   tradable,
			Shortable:  shortable,
			Marginable: marginable,
		})
	}

	if rows.Err() != nil {
		return stocks, rows.Err()
	}
	return stocks, err
}

func GetAllStocks(db PGConnection) (stocks []Stock, err error) {
	return getStocks(db, fmt.Sprintf(stocksQuery, ""))
}

func GetTradableStocks(db PGConnection) (stocks []Stock, err error) {
	return getStocks(db, fmt.Sprintf(stocksQuery, " WHERE tradable = true"))
}

func GetTradableStocksBySymbol(db PGConnection) (map[string]Stock, error) {
	stocks, err := GetTradableStocks(db)
	if err != nil {
		return nil, err
	}
	return GroupStocksBySymbol(stocks), nil
}

func GroupStocksBySymbol(stocks []Stock) map[string]Stock {
	stocksBySymbol := make(map[string]Stock)
	for _, stock := range stocks {
		stocksBySymbol[stock.Symbol] = stock
	}
	return stocksBySymbol
}

const tradableStockQuery = `
SELECT id, company, exchange, tradable, shortable, marginable
FROM stocks
WHERE tradable = true AND symbol = $1
`

func GetTradableStock(db PGConnection, symbol string) (stock Stock, err error) {
	row := db.QueryRow(context.Background(), tradableStockQuery, symbol)
	var id int
	var company string
	var exchange string
	var tradable bool
	var shortable bool
	var marginable bool
	if err = row.Scan(&id, &company, &exchange, &tradable, &shortable, &marginable); err != nil {
		return stock, err
	}
	stock = Stock{
		Id:         id,
		Symbol:     symbol,
		Exchange:   exchange,
		Company:    company,
		Tradable:   tradable,
		Shortable:  shortable,
		Marginable: marginable,
	}

	return stock, nil
}

const updateUnsupportedStocksQuery = `
UPDATE stocks
SET tradable = false
WHERE symbol != ALL($1)
`

func UpdateUnsupportedStocks(db PGConnection, supportedSymbols []string) error {
	if len(supportedSymbols) == 0 {
		return errors.New("no supported symbols")
	}

	_, err := db.Exec(
		context.Background(),
		updateUnsupportedStocksQuery,
		supportedSymbols,
	)
	return err
}

const updateStockQuery = `
UPDATE stocks
SET company = $1, tradable = $2, shortable = $3, marginable = $4, exchange = $5
WHERE symbol = $6
`

func UpdateStock(db PGConnection, stock Stock) error {
	_, err := db.Exec(
		context.Background(),
		updateStockQuery,
		stock.Company,
		stock.Tradable,
		stock.Shortable,
		stock.Company,
		stock.Exchange,
		stock.Symbol,
	)
	return err
}

var columns = []string{"symbol", "company", "exchange", "tradable", "marginable", "shortable"}

func InsertNewStocks(db PGConnection, stocks []Stock) error {
	if len(stocks) == 0 {
		return nil
	}

	rows := make([][]interface{}, 0, len(stocks))
	for _, stock := range stocks {
		rows = append(rows, []interface{}{
			stock.Symbol,
			stock.Company,
			stock.Exchange,
			stock.Tradable,
			stock.Marginable,
			stock.Shortable,
		})
	}

	_, err := db.CopyFrom(
		context.Background(),
		pgx.Identifier{"stocks"},
		columns,
		pgx.CopyFromRows(rows),
	)
	return err
}
