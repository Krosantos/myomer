package socket

import (
	"encoding/json"
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

// Convenience wrapper to return as marshaled JSON instead of a string.
func (c client) marshal(m interface{}) error {
	raw, err := c.read()
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(raw), m)
	return err
}
