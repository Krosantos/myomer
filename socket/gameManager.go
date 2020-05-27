package socket

import (
	"errors"
	"sync"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/krosantos/myomer/v2/game"
	"github.com/krosantos/myomer/v2/manager"
)

type gameManager struct {
	pool        *pgxpool.Pool
	activeGames map[string]*match
	register    chan *match
	remove      chan *match
}

// makeGameManager -- Return a pointer to a new gameManager instance, and set it in motion
func makeGameManager(pool *pgxpool.Pool) gameManager {
	gm := gameManager{
		pool:        pool,
		activeGames: make(map[string]*match),
		register:    make(chan *match),
		remove:      make(chan *match),
	}
	go gm.start()
	return gm
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

// reconnect -- attempt to replace a match connection with an incoming connection
func (gm gameManager) reconnect(c *client, uid string, gid string) error {
	match, ok := gm.activeGames[gid]
	if ok != true {
		for g := range gm.activeGames {
			println(g)
		}
		return errors.New("game not found")
	}
	player, ok := match.players[uid]
	if ok != true {
		return errors.New("user not found in game")
	}
	match.lock.Lock()
	player.client = c
	go match.listenToPlayer(player)
	match.lock.Unlock()
	match.broadcast("Player reconnected")
	return nil
}

// buildMatch -- given two successfully matched users, set up a game.
func (gm gameManager) buildMatch(mc1 *matchCandidate, mc2 *matchCandidate) error {
	match := &match{
		id:      uuid.New().String(),
		lock:    &sync.Mutex{},
		game:    game.BuildGame(),
		players: make(map[string]*player),
		active:  true,
	}
	go match.listenToGame()
	gm.activeGames[match.id] = match
	p1 := &player{
		id:     mc1.uid,
		name:   mc1.name,
		team:   0,
		client: mc1.client,
	}
	p2 := &player{
		id:     mc2.uid,
		name:   mc2.name,
		team:   1,
		client: mc2.client,
	}
	p1a, err := manager.FindArmyByID(gm.pool, mc1.aid)
	if err != nil {
		return err
	}
	p2a, err := manager.FindArmyByID(gm.pool, mc2.aid)
	if err != nil {
		return err
	}

	match.game.PopulateArmy(p1a.Cohort, 0)
	match.game.PopulateArmy(p2a.Cohort, 1)
	match.addPlayer(p1)
	match.addPlayer(p2)

	return nil
}
