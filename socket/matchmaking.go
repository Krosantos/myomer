package socket

import (
	"errors"
	"sync"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/krosantos/myomer/v2/manager"
)

type matchmaking struct {
	pool       *pgxpool.Pool
	lock       *sync.Mutex
	candidates map[*matchCandidate]bool
	register   chan *matchCandidate
	remove     chan *matchCandidate
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
func makeMatchMaking(pool *pgxpool.Pool) matchmaking {
	mm := matchmaking{
		pool:       pool,
		lock:       &sync.Mutex{},
		candidates: make(map[*matchCandidate]bool),
		register:   make(chan *matchCandidate),
		remove:     make(chan *matchCandidate),
	}
	go mm.listen()
	go mm.match()
	return mm
}

// listen -- Continuously listen for candidates to enqueue/dequeue
func (m matchmaking) listen() {
	for {
		select {
		case mc := <-m.register:
			m.lock.Lock()
			m.candidates[mc] = true
			m.lock.Unlock()
			go m.receive(mc)
		case mc := <-m.remove:
			m.lock.Lock()
			delete(m.candidates, mc)
			m.lock.Unlock()
		}
	}
}

// match -- Continuously try to make matches
func (m matchmaking) match() {
	// In the future, this is an entire sub-project's worth of thinking,
	// and optimization, and fun. For now, just repeatedly pluck the top
	// two candidates.
	for range time.Tick(time.Second * time.Duration(1)) {
		var currentCandidate *matchCandidate
		m.lock.Lock()
		for matchCandidate := range m.candidates {
			if currentCandidate == nil {
				currentCandidate = matchCandidate
			}
			if currentCandidate != matchCandidate {
				// Match found, baby! ezpz
				err := gm.buildMatch(currentCandidate, matchCandidate)
				if err == nil {
					m.remove <- currentCandidate
					m.remove <- matchCandidate
				}
			}
		}
		m.lock.Unlock()
	}
}

// enqueue -- Attempt to enqueue a user, and a client, into matchmaking.r
func (m matchmaking) enqueue(c *client, uid string, aid string) error {
	u, err := manager.FindUserByID(m.pool, uid)
	if err != nil {
		return err
	}
	a, err := manager.FindArmyByID(m.pool, aid)
	if err != nil {
		return err
	}
	if a.UserID != uid {
		return errors.New("Unowned Army")
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
	mc.client.write("Enqueued")
	return nil
}

// receive -- Listen for incoming client messages
func (m matchmaking) receive(mc *matchCandidate) {
	for {
		s, err := mc.client.read()
		if err != nil {
			println("Error detected while in matchmaking")
			m.cancel(mc)
			break
		}
		if s == "cancel" {
			m.cancel(mc)
			break
		}
		if s == "ready" {
			break
		}
	}
}

//  cancel -- Remove a user from queue when they stop looking for a match
func (m matchmaking) cancel(mc *matchCandidate) {
	mc.client.conn.Close()
	m.remove <- mc
}
