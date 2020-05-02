package handlers

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/krosantos/myomer/v2/managers/users"
)

type loginRequestData struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func keyFunc(token *jwt.Token) (interface{}, error) {
	_, ok := token.Method.(*jwt.SigningMethodHMAC)
	if !ok {
		return nil, fmt.Errorf("Invalid signature")
	}
	return []byte(os.Getenv("JWT_SECRET")), nil
}

func jwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("authorization")
		split := strings.Split(auth, " ")
		if len(split) < 2 {
			c.AbortWithError(401, errors.New("Invalid Authentication"))
		}
		rawToken := split[1]
		token, err := jwt.Parse(rawToken, keyFunc)
		if err != nil {
			c.AbortWithError(401, err)
		}
		claims := token.Claims.(jwt.MapClaims)
		timestamp := int64(claims["expiryDate"].(float64))
		expTime := time.Unix(timestamp, 0)

		if time.Now().After(expTime) {
			c.AbortWithError(401, errors.New("Token Expired"))
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
	claims := jwt.MapClaims{
		"userId":     user.ID,
		"issueDate":  time.Now().Unix(),
		"expiryDate": time.Now().Add(time.Hour * time.Duration(1)).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		c.AbortWithError(400, err)
	}
	c.String(200, `{"token":"%v"}`, tokenString)
}
