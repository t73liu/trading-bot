package main

import (
	"context"
	"flag"
	"log"

	"github.com/t73liu/tradingbot/lib/traderdb"
	"github.com/t73liu/tradingbot/lib/utils"

	"github.com/jackc/pgx/v4"
)

func main() {
	dbURL := flag.String("db.url", "", "URL to connect to traderdb")
	email := flag.String("email", "", "New user email")
	password := flag.String("password", "", "New user password")
	flag.Parse()

	if *dbURL == "" {
		log.Fatalln("-db.url flag must be provided")
	}
	if *email == "" {
		log.Fatalln("-email flag must be specified")
	}
	if *password == "" {
		log.Fatalln("-password flag must be specified")
	}

	conn, err := pgx.Connect(context.Background(), *dbURL)
	if err != nil {
		log.Fatalln("Failed to connect to DB:", err)
	}

	hashedPassword, err := utils.HashPassword(*password)
	if err != nil {
		log.Fatalln("Failed to hash password:", err)
	}

	if err = traderdb.InsertNewUser(conn, traderdb.User{
		Email:          *email,
		HashedPassword: hashedPassword,
		IsActive:       true,
	}); err != nil {
		log.Fatalln("Failed to insert new user:", err)
	}
}
