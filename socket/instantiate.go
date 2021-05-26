package socket

import (
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"golang.org/x/net/websocket"
)

var gm gameManager
var mm matchmaking
var f foyer

// Initialize -- Set up the data structures which manage sockets over their lifetimes
func Initialize(pool *pgxpool.Pool) {
	gm = makeGameManager(pool)
	mm = makeMatchMaking(pool)
	f = makeFoyer()
}

// Handler -- The main handler for connections
func Handler(c *websocket.Conn) {
	client := &client{conn: c}
	f.register <- client
	for {
		time.Sleep(time.Minute * 1)
	}
}
