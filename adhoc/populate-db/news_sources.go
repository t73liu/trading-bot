package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/t73liu/trading-bot/lib/newsapi"
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
	apiKey := strings.TrimSpace(os.Getenv("NEWS_API_KEY"))
	if apiKey == "" {
		fmt.Println("NEWS_API_KEY environment variable is required")
		os.Exit(1)
	}

	conn, err := pgx.Connect(context.Background(), databaseUrl)
	if err != nil {
		fmt.Println("Failed to connect to DB:", err)
		os.Exit(1)
	}

	client := newsapi.NewClient(
		&http.Client{
			Timeout: 15 * time.Second,
		},
		apiKey,
	)

	sources, err := client.GetSources("", "en", "")
	if err != nil {
		fmt.Println("Failed to fetch news sources:", err)
		os.Exit(1)
	}

	err = bulkInsert(conn, sources)
	if err != nil {
		fmt.Println("Failed to populate DB with news sources:", err)
		os.Exit(1)
	}
}

func bulkInsert(conn *pgx.Conn, sources []newsapi.Source) error {
	tx, err := conn.Begin(context.Background())
	if err != nil {
		return err
	}
	// Rollback is safe to call even if the tx is already closed, so if
	// the tx commits successfully, this is a no-op
	defer tx.Rollback(context.Background())

	rows := make([][]interface{}, 0, len(sources))
	for _, source := range sources {
		rows = append(rows, []interface{}{source.Id, source.Name, source.Description, source.Url})
	}

	_, err = tx.CopyFrom(
		context.Background(),
		pgx.Identifier{"news_sources"},
		[]string{"id", "name", "description", "url"},
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
