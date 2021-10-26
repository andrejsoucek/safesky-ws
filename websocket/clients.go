package websocket

import (
	"encoding/json"
	"sync"

	"github.com/andrejsoucek/safesky-ws/aircraft"
	"github.com/andrejsoucek/safesky-ws/geography"
	socketio "github.com/googollee/go-socket.io"
	log "github.com/sirupsen/logrus"
)

type Clients struct {
	mutex       sync.Mutex
	Connections map[socketio.Conn]geography.BoundingBox
}

func (c *Clients) SetBoundingBox(conn socketio.Conn, bb geography.BoundingBox) {
	c.mutex.Lock()
	c.Connections[conn] = bb
	c.mutex.Unlock()
}

func (c *Clients) EmitAircrafts(aircrafts []aircraft.Aircraft) {
	c.mutex.Lock()
	for conn, bb := range c.Connections {
		visibleAircrafts := []aircraft.Aircraft{}
		for _, aircraft := range aircrafts {
			if geography.IsInBounds(bb, aircraft.LatLng) {
				visibleAircrafts = append(visibleAircrafts, aircraft)
			}
		}

		json, err := json.Marshal(visibleAircrafts)
		if err != nil {
			log.Error(err)
			continue
		}
		conn.Emit("aircrafts", string(json))
	}
	c.mutex.Unlock()
}

func (c *Clients) Remove(conn socketio.Conn) {
	c.mutex.Lock()
	delete(c.Connections, conn)
	c.mutex.Unlock()
}
