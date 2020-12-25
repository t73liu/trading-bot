package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"html/template"
	"net/smtp"
	"os"
	"strings"
	"time"
	"tradingbot/lib/traderdb"
	"tradingbot/lib/utils"
	"tradingbot/lib/yahoo-finance"

	"github.com/jackc/pgx/v4/pgxpool"
)

type EmailParams struct {
	Earnings []yahoo.EarningsCall
	IPOs     []yahoo.IPO
}

func main() {
	recipientsFlag := flag.String(
		"recipients",
		"",
		"Email addresses delimited by commas",
	)
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
	startTime := now.AddDate(0, 0, 7)
	endTime := now.AddDate(0, 0, 21)
	httpClient := utils.NewHttpClient()
	yahooClient := yahoo.NewClient(httpClient)

	pool, err := pgxpool.Connect(context.Background(), databaseUrl)
	if err != nil {
		fmt.Println("Failed to connect to DB:", err)
		os.Exit(1)
	}

	emailParams, err := getEmailParams(pool, yahooClient, startTime, endTime)
	if err != nil {
		fmt.Println("Failed to fetch news items:", err)
		os.Exit(1)
	}

	var body bytes.Buffer
	subject := fmt.Sprintf(
		"Subject: Earnings/IPO News %s to %s\n",
		formatTime(startTime),
		formatTime(endTime),
	)
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
	yahooClient *yahoo.Client,
	startTime time.Time,
	endTime time.Time,
) (params EmailParams, err error) {
	tradableStocksBySymbol, err := traderdb.GetTradableStocksBySymbol(db)
	if err != nil {
		return params, err
	}

	earnings, err := getEarnings(yahooClient, startTime, endTime)
	if err != nil {
		return params, err
	}
	earningTickers := make(map[string]struct{})
	filteredEarnings := make([]yahoo.EarningsCall, 0, len(earnings))
	for _, earningsCall := range earnings {
		_, ok := earningTickers[earningsCall.Ticker]
		if !ok {
			earningTickers[earningsCall.Ticker] = struct{}{}
			if _, ok := tradableStocksBySymbol[earningsCall.Ticker]; ok {
				filteredEarnings = append(filteredEarnings, earningsCall)
			}
		}
	}

	ipos, err := getIPOs(yahooClient, startTime, endTime)
	if err != nil {
		return params, err
	}
	ipoTickers := make(map[string]struct{})
	filteredIPOs := make([]yahoo.IPO, 0, len(ipos))
	for _, ipo := range ipos {
		_, ok := ipoTickers[ipo.Ticker]
		if !ok {
			ipoTickers[ipo.Ticker] = struct{}{}
			filteredIPOs = append(filteredIPOs, ipo)
		}
	}

	params = EmailParams{
		Earnings: filteredEarnings,
		IPOs:     filteredIPOs,
	}
	return params, nil
}

func getEarnings(
	client *yahoo.Client,
	startTime,
	endTime time.Time,
) (earnings []yahoo.EarningsCall, err error) {
	current := startTime
	for !current.Equal(endTime) {
		current = current.AddDate(0, 0, 1)
		earningsForDate, err := client.GetEarningsCall(current)
		if err != nil {
			return earnings, err
		}
		earnings = append(earnings, earningsForDate...)
	}
	return earnings, nil
}

func getIPOs(
	client *yahoo.Client,
	startTime,
	endTime time.Time,
) (ipos []yahoo.IPO, err error) {
	current := startTime
	for !current.Equal(endTime) {
		current = current.AddDate(0, 0, 1)
		iposForDate, err := client.GetIPOs(current)
		if err != nil {
			return ipos, err
		}
		ipos = append(ipos, iposForDate...)
	}
	return ipos, nil
}

func formatTime(moment time.Time) string {
	return moment.Format("2006-01-02")
}
