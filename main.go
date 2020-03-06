package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	fmt.Fprintf(os.Stdout, "Attempting to connect to %v\n", os.Getenv("DB_URL"))

	pool, err := pgxpool.Connect(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to the db: %v\n", err)
		os.Exit(1)
	}
	defer pool.Close()

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(200, "We in it now, boys")
	})

	router.Run(os.Getenv("PORT"))

}
