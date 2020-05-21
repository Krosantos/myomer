package matches

import (
	"fmt"
	"net"

	"github.com/krosantos/myomer/v2/game"
)

// Client -- The bit that sends and receives data
type Client struct {
	socket net.Conn
	data   chan []byte
}

// ClientManager -- The bit that tracks the clients
type ClientManager struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
}

func (manager *ClientManager) start() {
	for {
		select {
		case connection := <-manager.register:
			manager.clients[connection] = true
			fmt.Println("Added new connection!")
		case connection := <-manager.unregister:
			if _, ok := manager.clients[connection]; ok {
				close(connection.data)
				delete(manager.clients, connection)
				fmt.Println("A connection has terminated!")
			}
		case message := <-manager.broadcast:
			for connection := range manager.clients {
				select {
				case connection.data <- message:
				default:
					close(connection.data)
					delete(manager.clients, connection)
				}
			}
		}
	}
}

func (manager *ClientManager) receive(client *Client) {
	for {
		message := make([]byte, 4096)
		length, err := client.socket.Read(message)
		if err != nil {
			manager.unregister <- client
			client.socket.Close()
			break
		}
		if length > 0 {
			fmt.Println("RECEIVED: " + string(message))
			manager.broadcast <- message
		}
	}
}
func (manager *ClientManager) send(client *Client) {
	defer client.socket.Close()
	for {
		select {
		case message, ok := <-client.data:
			if !ok {
				return
			}
			client.socket.Write(message)
		}
	}
}
func (client *Client) receive() {
	for {
		message := make([]byte, 4096)
		length, err := client.socket.Read(message)
		if err != nil {
			client.socket.Close()
			break
		}
		if length > 0 {
			fmt.Println("RECEIVED: " + string(message))
		}
	}
}

// ListenForMatches -- Spin up a listener and a client manager and chilllll
func ListenForMatches() {

	game := game.BuildGame()
	army := `{
		"units": {
			"0": {
				"name": "Horse Guy",
				"cost": 50,
				"color": "red",
				"strength": 3,
				"health": 2,
				"speed": 3,
				"moxie": 40,
				"attackRange": 1,
				"attackType": "melee",
				"moveType": "basic",
				"onAttack": [],
				"onDie": [],
				"onKill": [],
				"onMove": [],
				"onStrike": [],
				"onStruck": [],
				"onTurnEnd": [],
				"activeAbilities": []
			}
		}
	}`
	game.PopulateArmy(army, 1)

	listener, error := net.Listen("tcp", ":4500")
	if error != nil {
		panic("SOCKET DEATH")
	}
	manager := ClientManager{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
	go manager.start()
	for {
		connection, _ := listener.Accept()
		client := &Client{socket: connection, data: make(chan []byte)}
		manager.register <- client
		go manager.receive(client)
		go manager.send(client)
	}
}
