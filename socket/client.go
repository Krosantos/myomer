package socket

import (
	"encoding/json"

	"golang.org/x/net/websocket"
)

type client struct {
	conn *websocket.Conn
}

func (c client) write(s string) error {
	return websocket.Message.Send(c.conn, s)
}

func (c client) read() (string, error) {
	var m string
	err := websocket.Message.Receive(c.conn, &m)
	return m, err
}

// Convenience wrapper to return as marshaled JSON instead of a string.
func (c client) marshal(m interface{}) error {
	raw, err := c.read()
	if err != nil {
		println("Error in client.Marshal:", err.Error())
		return err
	}
	err = json.Unmarshal([]byte(raw), m)
	return err
}
