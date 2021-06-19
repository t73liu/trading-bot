package main

import (
	"bytes"
	"context"
	"flag"
	"html/template"
	"log"
	"net/smtp"
	"strings"
	"time"

	"github.com/t73liu/tradingbot/lib/newsapi"
	"github.com/t73liu/tradingbot/lib/traderdb"
	"github.com/t73liu/tradingbot/lib/utils"

	"github.com/jackc/pgx/v4/pgxpool"
)

type EmailParams struct {
	GeneralNewsByCompany map[string][]newsapi.Article
}

const userID = 1

func main() {
	recipients := flag.String("recipients", "", "Email addresses delimited by commas")
	dbURL := flag.String("db.url", "", "URL to connect to traderdb")
	newsApiKey := flag.String("news.key", "", "News API Key")
	alpacaApiKey := flag.String("alpaca.key", "", "Alpaca API Key")
	senderEmail := flag.String("sender.email", "", "Sender gmail address")
	senderPassword := flag.String(
		"sender.password",
		"",
		"Sender's gmail password",
	)
	flag.Parse()

	if *recipients == "" {
		log.Fatalln("-recipients flag must be provided")
	}
	if *dbURL == "" {
		log.Fatalln("-db.url flag must be provided")
	}
	if *newsApiKey == "" {
		log.Fatalln("-news.key flag must be provided")
	}
	if *alpacaApiKey == "" {
		log.Fatalln("-alpaca.key flag must be provided")
	}
	if *senderEmail == "" {
		log.Fatalln("-sender.email flag must be provided")
	}
	if *senderPassword == "" {
		log.Fatalln("-sender.password flag must be provided")
	}

	emailTemplate, err := template.ParseFiles("email-template.html")
	if err != nil {
		log.Fatalln("Failed to parse email template:", err)
	}

	now := time.Now()
	httpClient := utils.NewHttpClient()
	newsClient := newsapi.NewClient(httpClient, *newsApiKey)

	pool, err := pgxpool.Connect(context.Background(), *dbURL)
	if err != nil {
		log.Fatalln("Failed to connect to DB:", err)
	}

	emailParams, err := getEmailParams(pool, newsClient, now)
	if err != nil {
		log.Fatalln("Failed to fetch news items:", err)
	}

	var body bytes.Buffer
	subject := "Subject: Watchlist News " + now.Format("Jan 02") + "\n"
	headers := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.WriteString(subject + headers)
	err = emailTemplate.Execute(&body, emailParams)
	if err != nil {
		log.Fatalln("Failed to populate email template:", err)
	}

	auth := smtp.PlainAuth("", *senderEmail, *senderPassword, "smtp.gmail.com")
	err = smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		*senderEmail,
		strings.Split(*recipients, ","),
		body.Bytes(),
	)
	if err != nil {
		log.Fatalln("Failed to send email:", err)
	}
}

func getEmailParams(
	db *pgxpool.Pool,
	newsClient *newsapi.Client,
	now time.Time,
) (params EmailParams, err error) {
	location, err := time.LoadLocation("America/New_York")
	if err != nil {
		return params, err
	}

	stocks, err := traderdb.GetWatchlistStocksWithUserID(db, userID)
	if err != nil {
		return params, err
	}

	startTime := utils.GetLastWeekday(now)

	generalNewsByCompany := make(map[string][]newsapi.Article)
	for _, stock := range stocks {
		// Using domains because some news sources are missing from /sources
		// (e.g. seekingalpha.com, yahoo.finance.com)
		response, err := newsClient.GetAllHeadlinesWithSources(newsapi.AllArticlesQueryParams{
			Query:     utils.TrimCompanyName(stock.Company) + " OR " + stock.Symbol + " Stock",
			StartTime: startTime,
			Domains:   utils.NewsDomains,
			PageSize:  20,
		})
		if err != nil {
			return params, err
		}
		for index := range response.Articles {
			response.Articles[index].PublishedAt = response.Articles[index].PublishedAt.In(location)
		}
		generalNewsByCompany[stock.Symbol] = response.Articles
	}

	params = EmailParams{
		GeneralNewsByCompany: generalNewsByCompany,
	}
	return params, nil
}
