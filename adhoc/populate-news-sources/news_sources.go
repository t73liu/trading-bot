package main

import (
	"context"
	"flag"
	"log"

	"github.com/t73liu/tradingbot/lib/newsapi"
	"github.com/t73liu/tradingbot/lib/utils"

	"github.com/jackc/pgx/v4"
)

func main() {
	dbURL := flag.String("db.url", "", "URL to connect to traderdb")
	newsApiKey := flag.String("news.key", "", "News API Key")
	flag.Parse()

	if *dbURL == "" {
		log.Fatalln("-db.url flag must be provided")
	}
	if *newsApiKey == "" {
		log.Fatalln("-news.key flag must be provided")
	}

	conn, err := pgx.Connect(context.Background(), *dbURL)
	if err != nil {
		log.Fatalln("Failed to connect to DB:", err)
	}

	client := newsapi.NewClient(utils.NewHttpClient(), *newsApiKey)

	sources, err := client.GetSources("", "en", "")
	if err != nil {
		log.Fatalln("Failed to fetch news sources:", err)
	}

	if err = bulkInsertNewsSources(conn, sources); err != nil {
		log.Fatalln("Failed to populate DB with news sources:", err)
	}
}

func bulkInsertNewsSources(conn *pgx.Conn, sources []newsapi.Source) error {
	tx, err := conn.Begin(context.Background())
	if err != nil {
		return err
	}
	defer tx.Rollback(context.Background())

	rows := make([][]interface{}, 0, len(sources))
	for _, source := range sources {
		rows = append(rows, []interface{}{source.ID, source.Name, source.Description, source.Url})
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

	if err = tx.Commit(context.Background()); err != nil {
		return err
	}

	return nil
}
