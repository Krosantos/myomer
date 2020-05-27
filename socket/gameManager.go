package socket

import (
	"errors"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/krosantos/myomer/v2/game"
	"github.com/krosantos/myomer/v2/manager"
)

type gameManager struct {
	pool        *pgxpool.Pool
	lock        *sync.Mutex
	activeGames map[string]*match
	register    chan *match
	remove      chan *match
}

// makeGameManager -- Return a pointer to a new gameManager instance, and set it in motion
func makeGameManager(pool *pgxpool.Pool) gameManager {
	gm := gameManager{
		pool:        pool,
		lock:        &sync.Mutex{},
		activeGames: make(map[string]*match),
		register:    make(chan *match),
		remove:      make(chan *match),
	}
	go gm.start()
	go gm.prune()
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

// prune -- periodically check for and remove inactive games
func (gm gameManager) prune() {
	for range time.Tick(time.Second * time.Duration(30)) {
		gm.lock.Lock()
		for gid, g := range gm.activeGames {
			hasPlayers := false
			for _, p := range g.players {
				if p.active {
					hasPlayers = true
				}
			}
			if !g.active || !hasPlayers {
				delete(gm.activeGames, gid)
			}
		}
		gm.lock.Unlock()
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
	player.active = true
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
	gm.lock.Lock()
	gm.activeGames[match.id] = match
	gm.lock.Unlock()
	p1 := &player{
		id:     mc1.uid,
		name:   mc1.name,
		team:   0,
		client: mc1.client,
		active: true,
	}
	p2 := &player{
		id:     mc2.uid,
		name:   mc2.name,
		team:   1,
		client: mc2.client,
		active: true,
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
