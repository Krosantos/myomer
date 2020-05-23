package sockets

import (
	"encoding/json"
	"time"

	"github.com/krosantos/myomer/v2/auth"
)

type foyer struct {
	clients  map[*client]time.Time
	register chan *client
	remove   chan *client
}

// start -- Start the foyer, which will listen to register/remove clients
func (f foyer) start() {
	for {
		select {
		case c := <-f.register:
			f.clients[c] = time.Now()
		case c := <-f.remove:
			delete(f.clients, c)
		}
	}
}

// prune -- If a client's been in the foyer for too long, kill 'em.
func (f foyer) prune(d time.Duration) {
	for range time.Tick(d) {
		for client, t := range f.clients {
			ttl := t.Add(time.Second * time.Duration(30))
			if time.Now().After(ttl) {
				f.remove <- client
				client.conn.Close()
			}
		}
	}
}

// receive -- Parse and act on messages from held clients
func (f foyer) receive(c *client) {
	abort := func() {
		f.remove <- c
		c.conn.Close()
	}
	for {
		bloat := make([]byte, 4096)
		len, err := c.conn.Read(bloat)
		raw := bloat[:len]
		if err != nil {
			abort()
			break
		}
		m := foyerMessage{}
		err = json.Unmarshal(raw, &m)
		if err != nil {
			abort()
			break
		}
		if auth.JwtMatchesUser(m.Auth, m.UserID) == false {
			abort()
			break
		}
		if m.Action == "matchmake" {
			println("MATCH ME")
		} else if m.Action == "reconnect" {
			println("RECONNECT")
		}
	}
}

type foyerMessage struct {
	Action string `json:"type"`
	Auth   string `json:"auth"`
	UserID string `json:"userId"`
	ArmyID string `json:"armyId"`
	GameID string `json:"gameId"`
}
