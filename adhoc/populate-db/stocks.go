package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"github.com/jackc/pgx/v4"
	"os"
	"strings"
	"tradingbot/lib/alpaca"
	"tradingbot/lib/traderdb"
	"tradingbot/lib/utils"
)

func main() {
	databaseUrl := strings.TrimSpace(os.Getenv("DATABASE_URL"))
	if databaseUrl == "" {
		fmt.Println("DATABASE_URL environment variable is required")
		os.Exit(1)
	}
	apiKey := strings.TrimSpace(os.Getenv("ALPACA_API_KEY"))
	if apiKey == "" {
		fmt.Println("ALPACA_API_KEY environment variable is required")
		os.Exit(1)
	}
	apiSecretKey := strings.TrimSpace(os.Getenv("ALPACA_API_SECRET_KEY"))
	if apiSecretKey == "" {
		fmt.Println("ALPACA_API_SECRET_KEY environment variable is required")
		os.Exit(1)
	}

	conn, err := pgx.Connect(context.Background(), databaseUrl)
	if err != nil {
		fmt.Println("Failed to connect to DB:", err)
		os.Exit(1)
	}

	httpClient := utils.NewHttpClient()

	alpacaClient := alpaca.NewClient(
		httpClient,
		apiKey,
		apiSecretKey,
		false,
	)

	// Provides sensible company names courtesy of IEXCloud
	companyByTicker, err := getCompanyNamesByTicker()
	if err != nil {
		fmt.Println("Failed to read symbols.csv:", err)
		os.Exit(1)
	}

	stocksBySymbol, err := traderdb.GetTradableStocksBySymbol(conn)
	if err != nil {
		fmt.Println("Failed to get existing stocks by symbol:", err)
		os.Exit(1)
	}

	assets, err := alpacaClient.GetAssets("active", "")
	if err != nil {
		fmt.Println("Failed to fetch Alpaca supported stocks:", err)
		os.Exit(1)
	}
	assetsBySymbol := make(map[string]alpaca.Asset)
	for _, asset := range assets {
		if company, ok := companyByTicker[asset.Symbol]; ok {
			asset.Name = company
		}
		assetsBySymbol[asset.Symbol] = asset
	}

	inactiveSymbols := make([]string, 0, len(assets))
	for _, stock := range stocksBySymbol {
		if _, ok := assetsBySymbol[stock.Symbol]; !ok {
			inactiveSymbols = append(inactiveSymbols, stock.Symbol)
		}
	}

	updatedAssets := make([]alpaca.Asset, 0, len(assets))
	newAssets := make([]alpaca.Asset, 0, len(assets))
	for symbol, asset := range assetsBySymbol {
		stock, ok := stocksBySymbol[symbol]
		if ok {
			if !equalAssetAndStock(asset, stock) {
				updatedAssets = append(updatedAssets, asset)
			}
		} else {
			newAssets = append(newAssets, asset)
		}
	}

	if err = updateStocks(conn, inactiveSymbols, updatedAssets, newAssets); err != nil {
		fmt.Println("Failed to populate DB with stocks:", err)
		os.Exit(1)
	}
}

func equalAssetAndStock(asset alpaca.Asset, stock traderdb.Stock) bool {
	return stock.Marginable == asset.Marginable ||
		stock.Exchange == asset.Exchange ||
		stock.Company == asset.Name ||
		stock.Shortable == (asset.Shortable && asset.EasyToBorrow)
}

func updateStocks(conn *pgx.Conn, inactiveSymbols []string, updatedAssets []alpaca.Asset, newAssets []alpaca.Asset) error {
	tx, err := conn.Begin(context.Background())
	if err != nil {
		return err
	}
	defer tx.Rollback(context.Background())

	// Update inactive stocks
	for _, inactiveSymbol := range inactiveSymbols {
		_, err = tx.Exec(context.Background(), "UPDATE stocks SET tradable = false WHERE symbol = $1", inactiveSymbol)
		if err != nil {
			return err
		}
	}

	// Update outdated assets
	for _, asset := range updatedAssets {
		_, err = tx.Exec(
			context.Background(),
			"UPDATE stocks SET company = $1, tradable = $2, shortable = $3, marginable = $4, exchange = $5 WHERE symbol = $6",
			asset.Name,
			asset.Tradable,
			asset.Shortable && asset.EasyToBorrow,
			asset.Marginable,
			asset.Exchange,
			asset.Symbol,
		)
		if err != nil {
			return err
		}
	}

	// Insert new assets
	if len(newAssets) != 0 {
		rows := make([][]interface{}, 0, len(newAssets))
		for _, asset := range newAssets {
			rows = append(rows, []interface{}{
				asset.Symbol,
				asset.Name,
				asset.Exchange,
				asset.Tradable,
				asset.Marginable,
				asset.Shortable && asset.EasyToBorrow,
			})
		}

		_, err = tx.CopyFrom(
			context.Background(),
			pgx.Identifier{"stocks"},
			[]string{"symbol", "company", "exchange", "tradable", "marginable", "shortable"},
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

func getCompanyNamesByTicker() (map[string]string, error) {
	file, err := os.Open("symbols.csv")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Skip column headers
	if _, err = reader.Read(); err != nil {
		return nil, err
	}
	lines, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	companyBySymbol := make(map[string]string)
	for _, line := range lines {
		symbol := line[0]
		company := line[1]
		companyBySymbol[symbol] = company
	}
	return companyBySymbol, nil
}
