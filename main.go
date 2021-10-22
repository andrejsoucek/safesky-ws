package main

import (
	"fmt"
	"time"

	"github.com/andrejsoucek/safesky-ws/aircraft"
	"github.com/andrejsoucek/safesky-ws/geography"
	"github.com/andrejsoucek/safesky-ws/safesky"
	"github.com/andrejsoucek/safesky-ws/websocket"
)

func doEvery(d time.Duration, f func() ([]aircraft.Aircraft, error)) {
	for range time.Tick(d) {
		aircrafts, err := f()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(aircrafts)
	}
}

func main() {
	clients := map[string]geography.BoundingBox{}
	onConnect := func(id string, bb geography.BoundingBox) {
		clients[id] = bb
		fmt.Println(clients)
	}
	onBBUpdate := func(id string, bb geography.BoundingBox) {
		clients[id] = bb
		fmt.Println(clients)
	}
	onDisconnect := func(id string) {
		delete(clients, id)
		fmt.Println(clients)
	}
	go websocket.Listen(onConnect, onBBUpdate, onDisconnect)
	doEvery(4000*time.Millisecond, safesky.GetAircrafts)
}
