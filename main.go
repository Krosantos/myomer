package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
	"github.com/krosantos/myomer/v2/handler"
	"github.com/krosantos/myomer/v2/socket"
)

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Fprintf(os.Stderr, "No .env file found")
	}
}

func main() {
	pool, err := pgxpool.Connect(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to the db: %v\n", err)
		os.Exit(1)
	}
	defer pool.Close()
	go socket.Initialize(pool)

	router := handler.PrepareRouter(pool)
	router.Run(os.Getenv("PORT"))
}
