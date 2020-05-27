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
	active  bool
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
	p.client.write("Successfully added to game " + m.id)
	go m.listenToPlayer(p)
}

// listenToGame -- Listen to instructions from the game, broadcast them out
func (m match) listenToGame() {
	for {
		select {
		case i := <-m.game.Out:
			m.broadcast(i.ToString())
		case <-m.game.End:
			for _, p := range m.players {
				p.client.conn.Close()
			}
			m.active = false
			break
		}
	}
}

// listenToPlayer -- Listen to incoming messages
func (m match) listenToPlayer(p *player) {
	for {
		raw, err := p.client.read()
		if err != nil {
			m.broadcast(err.Error() + "-" + p.name)
			break
		}
		cmd, err := game.FormatCommand(raw)
		m.game.ParseCommand(cmd)
	}
}

// broadcast -- Send a message to all players
func (m match) broadcast(s string) {
	m.lock.Lock()
	for _, p := range m.players {
		p.client.write(s)
	}
	m.lock.Unlock()
}
