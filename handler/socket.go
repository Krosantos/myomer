package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krosantos/myomer/v2/socket"
	"golang.org/x/net/websocket"
)

func getSocket(c *gin.Context) {
	server := websocket.Server{
		Handler:   socket.Handler,
		Handshake: func(*websocket.Config, *http.Request) error { return nil },
	}
	server.ServeHTTP(c.Writer, c.Request)
}
