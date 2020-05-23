package sockets

import (
	"net"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Instantiate -- Start up the cluster of sockets and structures designed to wrangle them.
func Instantiate(pool *pgxpool.Pool) {
	listener, error := net.Listen("tcp", ":4500")
	if error != nil {
		panic("Failed to Instantiate socket listener")
	}
	gm := makeGameManager()
	mm := makeMatchMaking(gm, pool)
	f := makeFoyer(mm)
	for {
		conn, _ := listener.Accept()
		client := &client{conn: conn, data: make(chan []byte)}
		f.register <- client
	}
}
