package socket

import (
	"sync"
	"time"
)

type foyer struct {
	matchmaking *matchmaking
	lock        *sync.Mutex
	clients     map[*client]time.Time
	register    chan *client
	remove      chan *client
}

// makeFoyer -- Instantiate a new foyer instance, and set it in motion
func makeFoyer(mm *matchmaking) *foyer {
	foyer := foyer{
		matchmaking: mm,
		lock:        &sync.Mutex{},
		clients:     make(map[*client]time.Time),
		register:    make(chan *client),
		remove:      make(chan *client),
	}

	go foyer.listen()
	go foyer.prune(time.Second * time.Duration(5))

	return &foyer
}

// listen -- Listen to register/remove clients
func (f foyer) listen() {
	for {
		select {
		case c := <-f.register:
			f.lock.Lock()
			f.clients[c] = time.Now()
			f.lock.Unlock()
			go f.receive(c)
		case c := <-f.remove:
			f.lock.Lock()
			delete(f.clients, c)
			f.lock.Unlock()
		}
	}
}

// prune -- If a client's been in the foyer for too long, kill 'em.
func (f foyer) prune(d time.Duration) {
	for range time.Tick(d) {
		f.lock.Lock()
		for client, t := range f.clients {
			ttl := t.Add(time.Second * time.Duration(30))
			if time.Now().After(ttl) {
				f.deregister(client, true)
			}
		}
		f.lock.Unlock()
	}
}

// deregister -- cleanly remove a client from the foyer, optionally killing it
func (f foyer) deregister(c *client, kill bool) {
	if kill == true {
		c.conn.Close()
	}
	f.remove <- c
}

// receive -- Parse and act on messages from held clients
func (f foyer) receive(c *client) {
	for {
		m := foyerMessage{}
		err := c.marshal(&m)
		if err != nil {
			f.deregister(c, true)
			break
		}
		// if auth.JwtMatchesUser(m.Auth, m.UserID) == false {
		// 	f.deregister(c, true)
		// 	break
		// }
		if m.Action == "matchmake" {
			err := f.matchmaking.enqueue(c, m.UserID, m.ArmyID)
			if err == nil {
				println("Successfully enqueued" + m.UserID)
				f.deregister(c, false)
			} else {
				println("Enqueue error:", err.Error())
			}
			break
		} else if m.Action == "reconnect" {
			err := f.matchmaking.gameManager.reconnect(c, m.UserID, m.GameID)
			if err == nil {
				f.deregister(c, false)
			} else {
				println("Reconnect error:", err.Error())
			}
			break
		}
	}
}

type foyerMessage struct {
	Action string `json:"action"`
	Auth   string `json:"auth"`
	UserID string `json:"userId"`
	ArmyID string `json:"armyId"`
	GameID string `json:"gameId"`
}
