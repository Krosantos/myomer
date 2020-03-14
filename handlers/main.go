package handlers

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/krosantos/myomer/v2/managers/users"
)

type newUserRequestData struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"email"`
	Password string `json:"password" binding:"required"`
}

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
		rawToken := split[1]
		token, err := jwt.Parse(rawToken, keyFunc)
		if err != nil {
			c.AbortWithError(403, err)
		}
		fmt.Println(token.Claims)
		// TODO: Check expiry
		c.Next()
	}
}

// PrepareRouter -- Gets a gin router, loaded with the appropriate routes and middleware.
func PrepareRouter(pool *pgxpool.Pool) *gin.Engine {
	router := gin.Default()

	router.POST("/login", func(c *gin.Context) {
		var requestData = loginRequestData{}
		bindErr := c.ShouldBindJSON(&requestData)
		if bindErr != nil {
			c.AbortWithError(400, bindErr)
		}
		valid := users.ValidateLogin(pool, requestData.Email, requestData.Password)
		if valid != true {
			c.AbortWithStatus(403)
		}
		user, findErr := users.FindUserByEmail(pool, requestData.Email)
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
	})

	router.POST("/users", func(c *gin.Context) {
		var requestData = newUserRequestData{}
		bindErr := c.ShouldBindJSON(&requestData)
		if bindErr != nil {
			// TODO: Securely surface  binding errors (e.g., not an email)
			fmt.Println(bindErr.Error())
			return
		}
		err := users.CreateUser(pool, requestData.Email, requestData.Username, requestData.Password)
		if err == nil {
			c.Status(200)
		} else {
			fmt.Println(err)
			c.Status(400)
		}
	})

	router.Use(jwtAuth())

	router.GET("/", func(c *gin.Context) {
		c.String(200, "We in it now, boys")
	})

	return router
}
