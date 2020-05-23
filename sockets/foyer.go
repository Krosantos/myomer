package sockets

import (
	"encoding/json"
	"net"
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
			close(c.data)
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
		raw := make([]byte, 4096)
		_, err := c.conn.Read(raw)
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

// Instantiate -- Start up the cluster of sockets and structures designed to wrangle them.
func Instantiate() {
	listener, error := net.Listen("tcp", ":4500")
	if error != nil {
		panic("SOCKET DEATH")
	}
	foyer := foyer{
		clients:  make(map[*client]time.Time),
		register: make(chan *client),
		remove:   make(chan *client),
	}
	go foyer.start()
	go foyer.prune(time.Second * time.Duration(5))
	for {
		conn, _ := listener.Accept()
		client := &client{conn: conn, data: make(chan []byte)}
		foyer.register <- client
		go foyer.receive(client)
	}
}
