package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"
	"tradingbot/lib/traderdb"
	"tradingbot/lib/utils"

	"github.com/jackc/pgx/v4"
)

func main() {
	databaseUrl := strings.TrimSpace(os.Getenv("DATABASE_URL"))
	if databaseUrl == "" {
		fmt.Println("DATABASE_URL environment variable is required")
		os.Exit(1)
	}

	emailFlag := flag.String("email", "", "New user email")
	passwordFlag := flag.String("password", "", "New user password")
	flag.Parse()

	if *emailFlag == "" {
		fmt.Println("email must be specified")
		os.Exit(1)
	}

	if *passwordFlag == "" {
		fmt.Println("password must be specified")
		os.Exit(1)
	}

	conn, err := pgx.Connect(context.Background(), databaseUrl)
	if err != nil {
		fmt.Println("Failed to connect to DB:", err)
		os.Exit(1)
	}

	hashedPassword, err := utils.HashPassword(*passwordFlag)
	if err != nil {
		fmt.Println("Failed to hash password:", err)
		os.Exit(1)
	}

	if err = traderdb.InsertNewUser(conn, traderdb.User{
		Email:          *emailFlag,
		HashedPassword: hashedPassword,
		IsActive:       true,
	}); err != nil {
		fmt.Println("Failed to insert new user:", err)
		os.Exit(1)
	}
}
