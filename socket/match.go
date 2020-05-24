package socket

import "github.com/krosantos/myomer/v2/game"

type match struct {
	id      string
	game    game.Game
	players []*player
}

type player struct {
	id     string
	name   string
	team   int
	client *client
}
