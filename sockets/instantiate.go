package sockets

import (
	"net"
	"time"
)

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
