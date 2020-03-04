package models

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

// User -- an individual in the DB
type User struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// FindUserByID -- Acquire a single User based on id
func FindUserByID(pool *pgxpool.Pool, id int64) (User, error) {
	user := User{}
	err := pool.QueryRow(context.Background(), "select * from users where id = $1", id).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	return user, err
}

// ValidateLogin -- Confirm login with a given email and plaintext password
func ValidateLogin(pool *pgxpool.Pool, email string, plaintext string) bool {
	var hash string
	readErr := pool.QueryRow(context.Background(), "select password from users where email = $1", email).Scan(&hash)
	if readErr != nil {
		return false
	}

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(plaintext))
	return err == nil
}

// CreateUser -- Create a new user in the DB, returning any errors
func CreateUser(pool *pgxpool.Pool, email string, username string, plaintext string) error {
	hashed, hashErr := bcrypt.GenerateFromPassword([]byte(plaintext), 14)
	if hashErr != nil {
		return hashErr
	}
	_, err := pool.Exec(context.Background(), "INSERT INTO users (username, email, password, created_at, updated_at) VALUES ($1, $2, $3, NOW(), NOW());", username, email, hashed)

	return err
}
