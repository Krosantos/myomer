package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/krosantos/myomer/v2/managers/users"
)

type newUserRequestData struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// PrepareRouter -- Gets a gin router, loaded with the appropriate routes and middleware.
func PrepareRouter(pool *pgxpool.Pool) *gin.Engine {

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(200, "We in it now, boys")
	})
	router.POST("/users", func(c *gin.Context) {
		var requestData = newUserRequestData{}
		c.BindJSON(&requestData)
		err := users.CreateUser(pool, requestData.Email, requestData.Username, requestData.Password)
		if err == nil {
			c.Status(200)
		} else {
			c.Status(400)
		}
	})

	return router
}
