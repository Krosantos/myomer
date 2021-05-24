package handler

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/krosantos/myomer/v2/auth"
	"github.com/krosantos/myomer/v2/manager"
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
		return
	}
	dbPool := c.MustGet(cxtDbPool).(*pgxpool.Pool)
	valid := manager.ValidateLogin(dbPool, requestData.Email, requestData.Password)
	if !valid {
		c.AbortWithStatus(401)
		return
	}
	user, findErr := manager.FindUserByEmail(dbPool, requestData.Email)
	if findErr != nil {
		c.AbortWithError(400, findErr)
		return
	}
	claims := make(map[string]interface{})
	claims["userId"] = user.ID
	tokenString, err := auth.WriteJwt(claims, 60)
	if err != nil {
		c.AbortWithError(400, err)
		return
	}
	c.String(200, `{"token":"%v"}`, tokenString)
}
