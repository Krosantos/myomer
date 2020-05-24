package manager

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Army -- A template for loading preset units
type Army struct {
	ID        string    `json:"id"`
	UserID    string    `json:"userId"`
	Cohort    string    `json:"cohort"`
	Auxiliary string    `json:"auxiliary"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// FindArmyByID -- Acquire a single Army based on ID
func FindArmyByID(pool *pgxpool.Pool, aid string) (Army, error) {
	army := Army{}
	err := pool.QueryRow(context.Background(), "select * from armies where id = $1", aid).Scan(&army.ID, &army.UserID, &army.Cohort, &army.Auxiliary, &army.CreatedAt, &army.UpdatedAt)
	return army, err
}
