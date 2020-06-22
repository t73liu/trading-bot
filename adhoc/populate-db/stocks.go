package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/t73liu/trading-bot/lib/alpaca"
	"net/http"
	"os"
	"strings"
	"time"
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

	client := alpaca.NewClient(
		&http.Client{
			Timeout: 15 * time.Second,
		},
		apiKey,
		apiSecretKey,
		false,
	)

	assets, err := client.GetAssets("active", "")
	if err != nil {
		fmt.Println("Failed to fetch Alpaca supported stocks:", err)
		os.Exit(1)
	}

	err = bulkInsertStocks(conn, assets)
	if err != nil {
		fmt.Println("Failed to populate DB with stocks:", err)
		os.Exit(1)
	}
}

func bulkInsertStocks(conn *pgx.Conn, assets []alpaca.Asset) error {
	tx, err := conn.Begin(context.Background())
	if err != nil {
		return err
	}
	// Rollback is safe to call even if the tx is already closed, so if
	// the tx commits successfully, this is a no-op
	defer tx.Rollback(context.Background())

	rows := make([][]interface{}, 0, len(assets))
	for _, asset := range assets {
		rows = append(rows, []interface{}{
			asset.Symbol,
			asset.Name,
			asset.Exchange,
			asset.Tradable,
		})
	}

	_, err = tx.CopyFrom(
		context.Background(),
		pgx.Identifier{"stocks"},
		[]string{"symbol", "company", "exchange", "is_tradable"},
		pgx.CopyFromRows(rows),
	)
	if err != nil {
		return err
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return err
	}

	return nil
}
