package sockets

import "net"

type client struct {
	conn net.Conn
	data chan []byte
}
