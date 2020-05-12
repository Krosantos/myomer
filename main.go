package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/krosantos/myomer/v2/handlers"
)

func main() {
	println(os.Getenv("DB_URL"))
	pool, err := pgxpool.Connect(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to the db: %v\n", err)
		os.Exit(1)
	}
	defer pool.Close()

	router := handlers.PrepareRouter(pool)
	router.Run(os.Getenv("PORT"))
}
