package socket

import (
	"bufio"
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
	b := bufio.NewReader(c.conn)

	raw, err := b.ReadBytes(byte('\n'))
	if err != nil {
		println("Fuck you", err.Error())
		return "", err
	}
	s := string(raw)
	println("Internal register --", s)
	return s, nil
}

// Convenience wrapper to return as marshaled JSON instead of a string.
func (c client) marshal(m interface{}) error {
	raw, err := c.read()
	println(raw)
	if err != nil {
		println("Error in client.Marshal", err.Error())
		return err
	}
	err = json.Unmarshal([]byte(raw), m)
	return err
}
