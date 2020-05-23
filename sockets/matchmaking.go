package sockets

import (
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/krosantos/myomer/v2/managers/users"
)

type matchmaking struct {
	gameManager *gameManager
	pool        *pgxpool.Pool
	candidates  map[*matchCandidate]bool
	register    chan *matchCandidate
	remove      chan *matchCandidate
}

type matchCandidate struct {
	uid    string
	aid    string
	name   string
	elo    int
	ts     time.Time
	client *client
}

// makeMatchMaking -- Return a pointer to a new matchmaking instance, and set it in motion
func makeMatchMaking(gm *gameManager, pool *pgxpool.Pool) *matchmaking {
	mm := matchmaking{
		gameManager: gm,
		pool:        pool,
		candidates:  make(map[*matchCandidate]bool),
		register:    make(chan *matchCandidate),
		remove:      make(chan *matchCandidate),
	}
	go mm.listen()
	go mm.match()
	return &mm
}

// listen -- Continuously listen for stuff
func (m matchmaking) listen() {
	for {
		select {
		case mc := <-m.register:
			m.candidates[mc] = true
		case mc := <-m.remove:
			delete(m.candidates, mc)
		}
	}
}

// match -- Continuously try to make matches
func (m matchmaking) match() {
	// In the future, this is an entire sub-project's worth of thinking,
	// and optimization, and fun. For now, just repeatedly pluck the top
	// two candidates.
	var cc *matchCandidate
	for {
		for mc := range m.candidates {
			if cc == nil {
				cc = mc
			}
			if cc != mc {
				// Match found, baby! ezpz
				m.buildMatch(cc, mc)
			}
		}
	}
}

// enqueue -- Attempt to enqueue a user, and a client, into matchmaking. Returns bool of success.
func (m matchmaking) enqueue(c *client, uid string, aid string) bool {
	u, err := users.FindUserByID(m.pool, uid)
	if err != nil {
		return false
	}
	mc := &matchCandidate{
		uid:    u.ID,
		aid:    aid,
		name:   u.Username,
		elo:    u.Elo,
		ts:     time.Now(),
		client: c,
	}
	m.register <- mc
	return true
}

//  cancel -- Remove a user from queue when they stop looking for a match
func (m matchmaking) cancel(mc *matchCandidate) {
	m.remove <- mc
	mc.client.conn.Close()
}

// buildMatch -- Once a match is made, alert players, and pass them off to the game manager
func (m matchmaking) buildMatch(p1 *matchCandidate, p2 *matchCandidate) {
	p1.client.conn.Write([]byte("MATCH FOUND"))
	p2.client.conn.Write([]byte("MATCH FOUND"))
}
