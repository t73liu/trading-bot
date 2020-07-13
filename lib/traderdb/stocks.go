package traderdb

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Stock struct {
	Id       int    `json:"id"`
	Symbol   string `json:"symbol"`
	Company  string `json:"company"`
	Exchange string `json:"exchange,omitempty"`
}

const tradableStocksQuery = `
SELECT id, symbol, company, exchange FROM stocks WHERE is_tradable = true
`

func GetTradableStocks(db *pgxpool.Pool) (stocks []Stock, err error) {
	rows, err := db.Query(context.Background(), tradableStocksQuery)
	if err != nil {
		return stocks, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var symbol string
		var company string
		var exchange string
		if err = rows.Scan(&id, &symbol, &company, &exchange); err != nil {
			return stocks, err
		}
		stocks = append(stocks, Stock{
			Id:       id,
			Symbol:   symbol,
			Company:  company,
			Exchange: exchange,
		})
	}

	if rows.Err() != nil {
		return stocks, rows.Err()
	}
	return stocks, err
}
