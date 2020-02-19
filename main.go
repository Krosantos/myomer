package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/krosantos/myomer/v2/models"
)

func main() {
	fmt.Fprintf(os.Stdout, "Attempting to connect to %v\n", os.Getenv("DB_URL"))

	pool, err := pgxpool.Connect(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to the db: %v\n", err)
		os.Exit(1)
	}
	defer pool.Close()

	user, err := models.FindUserByID(pool, 1)
	fmt.Println(user)
}
