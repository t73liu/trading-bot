package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"tradingbot/lib/alpaca"
	"tradingbot/lib/iex"
	"tradingbot/lib/traderdb"
	"tradingbot/lib/utils"

	"github.com/jackc/pgx/v4"
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
	apiSecret := strings.TrimSpace(os.Getenv("ALPACA_API_SECRET"))
	if apiSecret == "" {
		fmt.Println("ALPACA_API_SECRET environment variable is required")
		os.Exit(1)
	}
	iexToken := strings.TrimSpace(os.Getenv("IEX_API_TOKEN"))
	if iexToken == "" {
		fmt.Println("IEX_API_TOKEN environment variable is required")
		os.Exit(1)
	}

	conn, err := pgx.Connect(context.Background(), databaseUrl)
	if err != nil {
		fmt.Println("Failed to connect to DB:", err)
		os.Exit(1)
	}

	httpClient := utils.NewHttpClient()
	alpacaClient, err := alpaca.NewClient(alpaca.ClientConfig{
		HttpClient: httpClient,
		ApiKey:     apiKey,
		ApiSecret:  apiSecret,
		IsLive:     false,
		IsPaid:     false,
	})
	if err != nil {
		fmt.Println("Failed to initialize Alpaca client:", err)
		os.Exit(1)
	}
	iexClient := iex.NewClient(httpClient, iexToken)

	if err = populateStocks(alpacaClient, iexClient, conn); err != nil {
		fmt.Println("Failed to populate stocks table:", err)
		os.Exit(1)
	}
}

func populateStocks(client *alpaca.Client, iexClient *iex.Client, conn *pgx.Conn) error {
	// Get sensible company names
	iexStocks, err := iexClient.GetReferenceSymbols()
	if err != nil {
		return err
	}
	companyBySymbol := make(map[string]string)
	for _, iexStock := range iexStocks {
		companyBySymbol[iexStock.Symbol] = iexStock.Name
	}

	assets, err := client.GetAssets("active", "")
	if err != nil {
		return err
	}

	existingStocks, err := traderdb.GetAllStocks(conn)
	if err != nil {
		return err
	}
	existingStocksBySymbol := traderdb.GroupStocksBySymbol(existingStocks)

	var newStocks []traderdb.Stock
	var updatedStocks []traderdb.Stock
	supportedSymbols := make([]string, 0, len(assets))
	for _, asset := range assets {
		company, ok := companyBySymbol[asset.Symbol]
		if !ok {
			company = asset.Name
		}
		stock := traderdb.Stock{
			Symbol:     asset.Symbol,
			Company:    company,
			Exchange:   asset.Exchange,
			Tradable:   asset.Tradable,
			Shortable:  asset.Shortable && asset.EasyToBorrow,
			Marginable: asset.Marginable,
		}
		supportedSymbols = append(supportedSymbols, stock.Symbol)
		if existingStock, ok := existingStocksBySymbol[stock.Symbol]; ok {
			if !existingStock.Equal(stock) {
				updatedStocks = append(updatedStocks, stock)
			}
		} else {
			newStocks = append(newStocks, stock)
		}
	}

	tx, err := conn.Begin(context.Background())
	if err != nil {
		return err
	}
	defer tx.Rollback(context.Background())

	if err = traderdb.UpdateUnsupportedStocks(tx, supportedSymbols); err != nil {
		return err
	}

	fmt.Printf("Updating %d existing stocks", len(updatedStocks))
	for _, updatedStock := range updatedStocks {
		if err = traderdb.UpdateStock(tx, updatedStock); err != nil {
			return err
		}
	}

	fmt.Printf("Adding %d new stocks", len(newStocks))
	if err = traderdb.InsertNewStocks(tx, newStocks); err != nil {
		return err
	}

	if err = tx.Commit(context.Background()); err != nil {
		return err
	}
	return nil
}
