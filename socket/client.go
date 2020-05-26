package socket

import (
	"net"
)

type client struct {
	conn net.Conn
	data chan []byte
}

func (c client) write(s string) {
	c.conn.Write([]byte(s))
}

func (c client) read() (string, error) {
	tooLong := make([]byte, 4096)
	len, err := c.conn.Read(tooLong)
	if err != nil {
		return "", err
	}
	raw := tooLong[:len]
	return string(raw), nil
}
