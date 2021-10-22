package main

import (
	"fmt"
	"time"

	"github.com/andrejsoucek/safesky-ws/aircraft"
	"github.com/andrejsoucek/safesky-ws/geography"
	"github.com/andrejsoucek/safesky-ws/safesky"
	"github.com/andrejsoucek/safesky-ws/websocket"
	socketio "github.com/googollee/go-socket.io"
)

func doEvery(d time.Duration, f func() ([]aircraft.Aircraft, error), aircrafts []aircraft.Aircraft, clients map[int]websocket.Client) {
	for range time.Tick(d) {
		if len(clients) == 0 {
			continue
		}
		data, err := f()
		if err != nil {
			fmt.Println(err)
			websocket.EmitAircrafts(aircrafts, clients)
			continue
		}
		aircrafts = data

		websocket.EmitAircrafts(aircrafts, clients)
	}
}

func main() {
	aircrafts := []aircraft.Aircraft{}
	clients := map[int]websocket.Client{}
	onBBUpdate := func(id int, conn socketio.Conn, bb geography.BoundingBox) {
		clients[id] = websocket.Client{Conn: conn, Bb: bb}
		fmt.Println(clients)
	}
	onDisconnect := func(id int) {
		delete(clients, id)
		fmt.Println(clients)
	}
	go websocket.Listen(onBBUpdate, onDisconnect)
	doEvery(4000*time.Millisecond, safesky.GetAircrafts, aircrafts, clients)
}
