package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/krosantos/myomer/v2/managers/users"
)

type newUserRequestData struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"email"`
	Password string `json:"password" binding:"required"`
}

func postUsers(c *gin.Context) {
	var requestData = newUserRequestData{}
	bindErr := c.ShouldBindJSON(&requestData)
	if bindErr != nil {
		fmt.Println(bindErr.Error())
		return
	}
	dbPool := c.MustGet(cxtDbPool).(*pgxpool.Pool)
	err := users.CreateUser(dbPool, requestData.Email, requestData.Username, requestData.Password)
	if err == nil {
		c.Status(200)
	} else {
		fmt.Println(err)
		c.Status(400)
	}
}
