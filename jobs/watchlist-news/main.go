package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/t73liu/trading-bot/lib/polygon"
	"html/template"
	"net/http"
	"net/smtp"
	"os"
	"strings"
	"time"
)

type EmailParams struct {
	NewsByTicker map[string][]polygon.Article
}

func main() {
	recipientsFlag := flag.String("recipients", "", "Email addresses delimited by commas")
	flag.Parse()

	recipients := strings.TrimSpace(*recipientsFlag)
	if recipients == "" {
		fmt.Println("At least one recipient must be specified")
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
	apiKey := strings.TrimSpace(os.Getenv("ALPACA_API_KEY"))
	if apiKey == "" {
		fmt.Println("ALPACA_API_KEY environment variable is required")
		os.Exit(1)
	}

	emailTemplate, err := template.ParseFiles("email-template.html")
	if err != nil {
		fmt.Println("Failed to parse email template:", err)
		os.Exit(1)
	}

	now := time.Now()
	httpClient := &http.Client{Timeout: 15 * time.Second}
	polygonClient := polygon.NewClient(httpClient, apiKey)

	emailParams, err := getEmailParams(polygonClient)
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

func getEmailParams(polygonClient *polygon.Client) (params EmailParams, err error) {
	location, err := time.LoadLocation("EST")
	if err != nil {
		return params, err
	}

	newsByTicker := make(map[string][]polygon.Article)
	tickers := []string{"TSLA", "AAPL", "MSFT"}
	for _, ticker := range tickers {
		articles, err := polygonClient.GetTickerNews(ticker, 10, 1)
		if err != nil {
			return params, err
		}
		for index := range articles {
			articles[index].Timestamp = articles[index].Timestamp.In(location)
		}
		newsByTicker[ticker] = articles
	}

	params = EmailParams{
		NewsByTicker: newsByTicker,
	}
	return params, nil
}
