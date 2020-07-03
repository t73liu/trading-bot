package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/t73liu/trading-bot/lib/yahoo-finance"
	"html/template"
	"net/http"
	"net/smtp"
	"os"
	"strings"
	"time"
)

type EmailParams struct {
	Earnings []yahoo.EarningsCall
	IPOs     []yahoo.IPO
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

	emailTemplate, err := template.ParseFiles("email-template.html")
	if err != nil {
		fmt.Println("Failed to parse email template:", err)
		os.Exit(1)
	}

	now := time.Now()
	httpClient := &http.Client{Timeout: 15 * time.Second}
	yahooClient := yahoo.NewClient(httpClient)

	emailParams, err := getEmailParams(yahooClient, now)
	if err != nil {
		fmt.Println("Failed to fetch news items:", err)
		os.Exit(1)
	}

	var body bytes.Buffer
	subject := "Subject: Earnings/IPO News " + now.Format("Jan 02") + "\n"
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
	yahooClient *yahoo.Client,
	now time.Time,
) (params EmailParams, err error) {
	earnings, err := getEarnings(yahooClient, now)
	if err != nil {
		return params, err
	}
	earningTickers := make(map[string]struct{})
	filteredEarnings := make([]yahoo.EarningsCall, 0, len(earnings))
	for _, earningsCall := range earnings {
		_, ok := earningTickers[earningsCall.Ticker]
		if !ok {
			earningTickers[earningsCall.Ticker] = struct{}{}
			filteredEarnings = append(filteredEarnings, earningsCall)
		}
	}

	ipos, err := getIPOs(yahooClient, now)
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

func getEarnings(client *yahoo.Client, current time.Time) (earnings []yahoo.EarningsCall, err error) {
	for days := 0; days < 14; days++ {
		date := current.AddDate(0, 0, days)
		earningsForDate, err := client.GetEarningsCall(date)
		if err != nil {
			return earnings, err
		}
		earnings = append(earnings, earningsForDate...)
	}
	return earnings, nil
}

func getIPOs(client *yahoo.Client, current time.Time) (ipos []yahoo.IPO, err error) {
	for days := 0; days < 14; days++ {
		date := current.AddDate(0, 0, days)
		iposForDate, err := client.GetIPOs(date)
		if err != nil {
			return ipos, err
		}
		ipos = append(ipos, iposForDate...)
	}
	return ipos, nil
}
