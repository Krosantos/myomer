package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

func main() {
	fmt.Fprintf(os.Stdout, "Attempting to connect to %v\n", os.Getenv("DB_URL"))

	conn, err := pgx.Connect(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to the db: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), "select username from users")
	defer rows.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to scan row: %v\n", err)
		os.Exit(1)
	}
	for rows.Next() {
		var username string
		err = rows.Scan(&username)

		fmt.Println(username)
	}
}
