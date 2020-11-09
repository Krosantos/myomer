package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/krosantos/myomer/v2/auth"
	"github.com/krosantos/myomer/v2/manager"
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
	id, err := manager.CreateUser(dbPool, requestData.Email, requestData.Username, requestData.Password)
	if err == nil {
		claims := make(map[string]interface{})
		claims["userId"] = id
		tokenString, err := auth.WriteJwt(claims, 60)
		if err != nil {
			c.Status(200)
		}
		c.String(200, `{"token":"%v"}`, tokenString)
	} else {
		fmt.Println(err)

		c.String(400, err.Error())
	}
}
