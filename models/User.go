package models

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

// User -- an individual in the DB
type User struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Salt      string    `json:"salt"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// FindUserByID -- Acquire a single User based on id
func FindUserByID(pool *pgxpool.Pool, id int64) (User, error) {
	user := User{}
	err := pool.QueryRow(context.Background(), "select * from users where id = $1", id).Scan(&user.ID, &user.Username, &user.Email, &user.Salt, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	fmt.Println(user.CreatedAt.String())
	return user, err
}
