package matches

import (
	"github.com/krosantos/myomer/v2/elements"
)

// Match -- An in-progress match, with players, client connections, and a game
type Match struct {
	ID      string
	Game    elements.Game
	Players map[int]Player
}

// Player -- Placeholder player
type Player struct {
	Username string
	Team     int
	Client   Client
}
