package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/t73liu/trading-bot/lib/newsapi"
	"github.com/t73liu/trading-bot/lib/polygon"
	"github.com/t73liu/trading-bot/lib/traderdb"
	"html/template"
	"net/http"
	"net/smtp"
	"os"
	"strings"
	"time"
)

type EmailParams struct {
	NewsByTicker         map[string][]polygon.Article
	GeneralNewsByCompany map[string][]newsapi.Article
}

const userId = 1

var domains = []string{
	"finance.yahoo.com",
	"investors.com",
	"seekingalpha.com",
	"fool.com",
	"reuters.com",
	"bloomberg.com",
}

func main() {
	recipientsFlag := flag.String("recipients", "", "Email addresses delimited by commas")
	flag.Parse()

	recipients := strings.TrimSpace(*recipientsFlag)
	if recipients == "" {
		fmt.Println("At least one recipient must be specified")
		os.Exit(1)
	}

	databaseUrl := strings.TrimSpace(os.Getenv("DATABASE_URL"))
	if databaseUrl == "" {
		fmt.Println("DATABASE_URL environment variable is required")
		os.Exit(1)
	}
	newsAPIKey := strings.TrimSpace(os.Getenv("NEWS_API_KEY"))
	if newsAPIKey == "" {
		fmt.Println("NEWS_API_KEY environment variable is required")
		os.Exit(1)
	}
	alpacaAPIKey := strings.TrimSpace(os.Getenv("ALPACA_API_KEY"))
	if alpacaAPIKey == "" {
		fmt.Println("ALPACA_API_KEY environment variable is required")
		os.Exit(1)
	}
	email := strings.TrimSpace(os.Getenv("GMAIL_USERNAME"))
	if email == "" {
		fmt.Println("GMAIL_USERNAME environment variable is required")
		os.Exit(1)
	}
	password := strings.TrimSpace(os.Getenv("GMAIL_PASSWORD"))
	if password == "" {
		fmt.Println("GMAIL_PASSWORD environment variable is required")
		os.Exit(1)
	}

	emailTemplate, err := template.ParseFiles("email-template.html")
	if err != nil {
		fmt.Println("Failed to parse email template:", err)
		os.Exit(1)
	}

	now := time.Now()
	httpClient := &http.Client{Timeout: 15 * time.Second}
	polygonClient := polygon.NewClient(httpClient, alpacaAPIKey)
	newsClient := newsapi.NewClient(httpClient, newsAPIKey)

	pool, err := pgxpool.Connect(context.Background(), databaseUrl)
	if err != nil {
		fmt.Println("Failed to connect to DB:", err)
		os.Exit(1)
	}

	emailParams, err := getEmailParams(pool, polygonClient, newsClient, now)
	if err != nil {
		fmt.Println("Failed to fetch news items:", err)
		os.Exit(1)
	}

	var body bytes.Buffer
	subject := "Subject: Watchlist News " + now.Format("Jan 02") + "\n"
	headers := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.WriteString(subject + headers)
	err = emailTemplate.Execute(&body, emailParams)
	if err != nil {
		fmt.Println("Failed to populate email template:", err)
		os.Exit(1)
	}

	auth := smtp.PlainAuth("", email, password, "smtp.gmail.com")
	err = smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		email,
		strings.Split(recipients, ","),
		body.Bytes(),
	)
	if err != nil {
		fmt.Println("Failed to send email:", err)
		os.Exit(1)
	}
}

func getEmailParams(
	db *pgxpool.Pool,
	_ *polygon.Client,
	newsClient *newsapi.Client,
	now time.Time,
) (params EmailParams, err error) {
	location, err := time.LoadLocation("America/New_York")
	if err != nil {
		return params, err
	}

	stocks, err := traderdb.GetWatchlistStocksByUserId(db, userId)
	if err != nil {
		return params, err
	}

	startTime := getLastWeekday(now)

	generalNewsByCompany := make(map[string][]newsapi.Article)
	for _, stock := range stocks {
		// Using domains because some news sources are missing from /sources
		// (e.g. seekingalpha.com, yahoo.finance.com)
		response, err := newsClient.GetAllHeadlinesBySources(newsapi.AllArticlesQueryParams{
			Query:     trimCompanyName(stock.Company) + " OR " + stock.Symbol + " Stock",
			StartTime: startTime,
			Domains:   domains,
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

	// TODO Enable after https://github.com/polygon-io/issues/issues/25
	//newsByTicker := make(map[string][]polygon.Article)
	//for _, stock := range stocks {
	//	articles, err := polygonClient.GetTickerNews(stock.Symbol, 10, 1)
	//	if err != nil {
	//		return params, err
	//	}
	//	for index := range articles {
	//		articles[index].Timestamp = articles[index].Timestamp.In(location)
	//	}
	//	newsByTicker[stock.Symbol] = articles
	//}

	params = EmailParams{
		GeneralNewsByCompany: generalNewsByCompany,
		//NewsByTicker:         newsByTicker,
	}
	return params, nil
}

func getLastWeekday(now time.Time) time.Time {
	prevDay := now.AddDate(0, 0, -1)
	for prevDay.Weekday() == time.Saturday || prevDay.Weekday() == time.Sunday {
		prevDay = prevDay.AddDate(0, 0, -1)
	}
	return prevDay
}

func trimCompanyName(company string) string {
	trimmedName := strings.TrimSpace(strings.Split(company, " Class ")[0])
	//trimmedName = strings.TrimSpace(strings.Split(company, " Inc.")[0])
	return trimmedName
}
