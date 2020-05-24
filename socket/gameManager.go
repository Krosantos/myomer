package socket

import (
	"github.com/jackc/pgx/v4/pgxpool"
)

type gameManager struct {
	pool        *pgxpool.Pool
	activeGames map[string]*match
	register    chan *match
	remove      chan *match
}

// makeGameManager -- Return a pointer to a new gameManager instance, and set it in motion
func makeGameManager(pool *pgxpool.Pool) *gameManager {
	gm := gameManager{
		pool:        pool,
		activeGames: make(map[string]*match),
		register:    make(chan *match),
		remove:      make(chan *match),
	}
	go gm.start()
	return &gm
}

// start -- prepare to add and remove games from the registry
func (gm gameManager) start() {
	for {
		select {
		case match := <-gm.register:
			gm.activeGames[match.id] = match
		case match := <-gm.remove:
			delete(gm.activeGames, match.id)
		}
	}
}

// buildMatch -- given two successfully matched users, set up a game.
func (gm gameManager) buildMatch(mc1 *matchCandidate, mc2 *matchCandidate) error {

	return nil
}
