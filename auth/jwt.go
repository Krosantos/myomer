package auth

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func keyFunc(token *jwt.Token) (interface{}, error) {
	_, ok := token.Method.(*jwt.SigningMethodHMAC)
	if !ok {
		return nil, errors.New("Invalid signature")
	}
	return []byte(os.Getenv("JWT_SECRET")), nil
}

// WriteJwt -- Given a map of claims, and an expiry duration in minutes, return a token
func WriteJwt(claims map[string]interface{}, expiryInMinutes int) (string, error) {
	mappedClaims := jwt.MapClaims{}
	for key, val := range claims {
		mappedClaims[key] = val
	}
	mappedClaims["issueDate"] = time.Now().Unix()
	mappedClaims["expiryDate"] = time.Now().Add(time.Minute * time.Duration(expiryInMinutes)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mappedClaims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	return tokenString, err
}

// JwtIsValid -- Given a string, assess whether or not it is a valid, signed JWT
func JwtIsValid(rawToken string) (map[string]interface{}, error) {
	token, err := jwt.Parse(rawToken, keyFunc)
	if err != nil {
		return nil, err
	}
	claims := token.Claims.(jwt.MapClaims)
	timestamp := int64(claims["expiryDate"].(float64))
	expTime := time.Unix(timestamp, 0)

	if time.Now().After(expTime) {
		return nil, errors.New("Token Expired")
	}
	return claims, nil
}

// JwtMatchesUser -- Convenience method to ensure a JWT is valid, and has a userId claim which matches the passed userId
func JwtMatchesUser(rawToken string, userID int) bool {
	claims, err := JwtIsValid(rawToken)
	if err != nil {
		return false
	}
	foundID, ok := claims["userId"].(int)
	return ok && foundID == userID

}
