package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"tradingbot/lib/alpaca"
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
	apiSecretKey := strings.TrimSpace(os.Getenv("ALPACA_API_SECRET"))
	if apiSecretKey == "" {
		fmt.Println("ALPACA_API_SECRET environment variable is required")
		os.Exit(1)
	}

	conn, err := pgx.Connect(context.Background(), databaseUrl)
	if err != nil {
		fmt.Println("Failed to connect to DB:", err)
		os.Exit(1)
	}

	alpacaClient := alpaca.NewClient(
		utils.NewHttpClient(),
		apiKey,
		apiSecretKey,
		false,
	)

	if err = populateStocks(alpacaClient, conn); err != nil {
		fmt.Println("Failed to populate stocks table:", err)
		os.Exit(1)
	}
}

func populateStocks(client *alpaca.Client, conn *pgx.Conn) error {
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
		stock := traderdb.Stock{
			Symbol:     asset.Symbol,
			Company:    asset.Name,
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

	fmt.Println(fmt.Sprintf("Updating %d existing stocks", len(updatedStocks)))
	for _, updatedStock := range updatedStocks {
		if err = traderdb.UpdateStock(tx, updatedStock); err != nil {
			return err
		}
	}

	fmt.Println(fmt.Sprintf("Adding %d new stocks", len(newStocks)))
	if err = traderdb.InsertNewStocks(tx, newStocks); err != nil {
		return err
	}

	if err = tx.Commit(context.Background()); err != nil {
		return err
	}
	return nil
}
