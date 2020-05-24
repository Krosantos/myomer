package socket

import (
	"sync"

	"github.com/krosantos/myomer/v2/game"
)

type match struct {
	id      string
	lock    *sync.Mutex
	game    *game.Game
	players map[string]*player
}

type player struct {
	id     string
	name   string
	team   int
	client *client
}

// addPlayer -- Add a player to the slice, start listening to their messages
func (m match) addPlayer(p *player) {
	m.lock.Lock()
	m.players[p.id] = p
	m.lock.Unlock()
	p.client.conn.Write([]byte("Successfully added to game " + m.id))
	go m.listen(p)
}

// listen -- Listen to incoming messages
func (m match) listen(p *player) {
	for {
		bloat := make([]byte, 4096)
		len, err := p.client.conn.Read(bloat)
		if err != nil {
			m.broadcast(err.Error())
			break
		}
		raw := bloat[:len]
		m.broadcast("COPY!: " + string(raw))
	}
}

// broadcast -- Send a message to all players
func (m match) broadcast(s string) {
	m.lock.Lock()
	for _, p := range m.players {
		p.client.conn.Write([]byte(s))
	}
	m.lock.Unlock()
}
