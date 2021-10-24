package main

import (
	"time"

	"github.com/andrejsoucek/safesky-ws/aircraft"
	"github.com/andrejsoucek/safesky-ws/geography"
	"github.com/andrejsoucek/safesky-ws/safesky"
	"github.com/andrejsoucek/safesky-ws/websocket"
	socketio "github.com/googollee/go-socket.io"
)

func doEvery(d time.Duration, f func() ([]aircraft.Aircraft, error), aircrafts []aircraft.Aircraft, clients map[socketio.Conn]geography.BoundingBox) {
	for range time.Tick(d) {
		if len(clients) == 0 {
			continue
		}
		data, err := f()
		if err == nil {
			aircrafts = data
		}

		websocket.EmitAircrafts(aircrafts, clients)
	}
}

func main() {
	aircrafts := []aircraft.Aircraft{}
	clients := map[socketio.Conn]geography.BoundingBox{}
	go websocket.Listen(clients)
	doEvery(4000*time.Millisecond, safesky.GetAircrafts, aircrafts, clients)
}
