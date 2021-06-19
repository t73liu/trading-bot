package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/smtp"
	"strings"
	"time"

	"github.com/t73liu/tradingbot/lib/traderdb"
	"github.com/t73liu/tradingbot/lib/utils"
	"github.com/t73liu/tradingbot/lib/yahoo-finance"

	"github.com/jackc/pgx/v4/pgxpool"
)

type EmailParams struct {
	Earnings []yahoo.EarningsCall
	IPOs     []yahoo.IPO
}

func main() {
	recipients := flag.String(
		"recipients",
		"",
		"Email addresses delimited by commas",
	)
	dbURL := flag.String("db.url", "", "URL to connect to traderdb")
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
	startTime := now.AddDate(0, 0, 7)
	endTime := now.AddDate(0, 0, 21)
	httpClient := utils.NewHttpClient()
	yahooClient := yahoo.NewClient(httpClient)

	pool, err := pgxpool.Connect(context.Background(), *dbURL)
	if err != nil {
		log.Fatalln("Failed to connect to DB:", err)
	}

	emailParams, err := getEmailParams(pool, yahooClient, startTime, endTime)
	if err != nil {
		log.Fatalln("Failed to fetch news items:", err)
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
	yahooClient *yahoo.Client,
	startTime time.Time,
	endTime time.Time,
) (params EmailParams, err error) {
	tradableStocksBySymbol, err := traderdb.GetStocksBySymbol(db)
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
