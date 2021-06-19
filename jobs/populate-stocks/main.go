package main

import (
	"context"
	"flag"
	"log"

	"github.com/t73liu/tradingbot/lib/alpaca"
	"github.com/t73liu/tradingbot/lib/iex"
	"github.com/t73liu/tradingbot/lib/traderdb"
	"github.com/t73liu/tradingbot/lib/utils"

	"github.com/jackc/pgx/v4"
)

func main() {
	dbURL := flag.String("db.url", "", "URL to connect to traderdb")
	alpacaApiKey := flag.String("alpaca.key", "", "Alpaca API Key")
	alpacaApiSecret := flag.String("alpaca.secret", "", "Alpaca API Secret")
	iexApiToken := flag.String("iex.token", "", "IEX API Token")
	flag.Parse()

	if *dbURL == "" {
		log.Fatalln("-db.url flag must be provided")
	}
	if *alpacaApiKey == "" {
		log.Fatalln("-alpaca.key flag must be provided")
	}
	if *alpacaApiSecret == "" {
		log.Fatalln("-alpaca.secret flag must be provided")
	}
	if *iexApiToken == "" {
		log.Fatalln("-iex.token flag must be provided")
	}

	conn, err := pgx.Connect(context.Background(), *dbURL)
	if err != nil {
		log.Fatalln("Failed to connect to DB:", err)
	}

	httpClient := utils.NewHttpClient()
	alpacaClient := alpaca.NewClient(alpaca.ClientConfig{
		HttpClient:    httpClient,
		ApiKey:        *alpacaApiKey,
		ApiSecret:     *alpacaApiSecret,
		IsLiveTrading: false,
		IsPaidData:    false,
	})
	iexClient := iex.NewClient(httpClient, *iexApiToken)

	if err = populateStocks(alpacaClient, iexClient, conn); err != nil {
		log.Fatalln("Failed to populate stocks table:", err)
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

	log.Printf("Updating %d existing stocks", len(updatedStocks))
	for _, updatedStock := range updatedStocks {
		if err = traderdb.UpdateStock(tx, updatedStock); err != nil {
			return err
		}
	}

	log.Printf("Adding %d new stocks", len(newStocks))
	if err = traderdb.InsertNewStocks(tx, newStocks); err != nil {
		return err
	}

	if err = tx.Commit(context.Background()); err != nil {
		return err
	}
	return nil
}
