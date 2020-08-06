package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"github.com/jackc/pgx/v4"
	"os"
	"strings"
	"tradingbot/lib/alpaca"
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

	assets, err := alpacaClient.GetAssets("active", "")
	if err != nil {
		fmt.Println("Failed to fetch Alpaca supported stocks:", err)
		os.Exit(1)
	}

	if err = bulkInsertStocks(conn, assets, companyByTicker); err != nil {
		fmt.Println("Failed to populate DB with stocks:", err)
		os.Exit(1)
	}
}

func bulkInsertStocks(conn *pgx.Conn, assets []alpaca.Asset, companyByTicker map[string]string) error {
	tx, err := conn.Begin(context.Background())
	if err != nil {
		return err
	}
	defer tx.Rollback(context.Background())

	rows := make([][]interface{}, 0, len(assets))
	for _, asset := range assets {
		name := asset.Name
		if company, ok := companyByTicker[asset.Symbol]; ok {
			name = company
		}
		rows = append(rows, []interface{}{
			asset.Symbol,
			name,
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
