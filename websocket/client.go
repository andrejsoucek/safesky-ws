package websocket

import (
	"github.com/andrejsoucek/safesky-ws/geography"
	socketio "github.com/googollee/go-socket.io"
)

type Client struct {
	Conn socketio.Conn
	Bb   geography.BoundingBox
}
