package handlers

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/krosantos/myomer/v2/auth"
	"github.com/krosantos/myomer/v2/managers/users"
)

type loginRequestData struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func jwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		ah := c.GetHeader("authorization")
		split := strings.Split(ah, " ")
		if len(split) < 2 {
			c.AbortWithError(401, errors.New("Invalid Authentication"))
		}
		rawToken := split[1]
		_, err := auth.JwtIsValid(rawToken)
		if err != nil {
			c.AbortWithError(401, err)
		}
		c.Next()
	}
}

func postLogin(c *gin.Context) {
	var requestData = loginRequestData{}
	bindErr := c.ShouldBindJSON(&requestData)
	if bindErr != nil {
		c.AbortWithError(400, bindErr)
	}
	dbPool := c.MustGet(cxtDbPool).(*pgxpool.Pool)
	valid := users.ValidateLogin(dbPool, requestData.Email, requestData.Password)
	if valid != true {
		c.AbortWithStatus(401)
	}
	user, findErr := users.FindUserByEmail(dbPool, requestData.Email)
	if findErr != nil {
		c.AbortWithError(400, findErr)
	}
	claims := make(map[string]interface{})
	claims["userId"] = user.ID
	tokenString, err := auth.WriteJwt(claims, 60)
	if err != nil {
		c.AbortWithError(400, err)
	}
	c.String(200, `{"token":"%v"}`, tokenString)
}
