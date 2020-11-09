package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

const cxtDbPool = "DB_POOL"

func useDbPool(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(cxtDbPool, pool)
		c.Next()
	}
}

// PrepareRouter -- Gets a gin router, loaded with the appropriate routes and middleware.
func PrepareRouter(pool *pgxpool.Pool) *gin.Engine {
	router := gin.Default()

	router.Use(useDbPool(pool))

	router.POST("/login", postLogin)
	router.POST("/users", postUsers)
	
	router.Use(jwtAuth())
	router.GET("/socket", getSocket)

	router.GET("/", func(c *gin.Context) {
		c.String(200, "We in it now, boys")
	})

	return router
}
