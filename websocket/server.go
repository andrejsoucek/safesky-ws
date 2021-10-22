package websocket

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/andrejsoucek/safesky-ws/geography"
	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/polling"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
)

func Listen(
	onConnect func(id string, bb geography.BoundingBox),
	onBBUpdate func(id string, bb geography.BoundingBox),
	onDisconnect func(id string),
) {
	server := createServer()

	server.OnConnect("/", func(s socketio.Conn) error {
		onConnect(s.ID(), geography.BoundingBox{})
		return nil
	})

	server.OnEvent("/", "bb", func(s socketio.Conn, data string) {
		onBBUpdate(s.ID(), createBoundingBox(data))
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("closed", reason)
		onDisconnect(s.ID())
	})

	go server.Serve()
	defer server.Close()

	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("../asset")))

	log.Println("Serving at localhost:8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func createServer() *socketio.Server {
	var allowOriginFunc = func(r *http.Request) bool {
		return true
	}
	return socketio.NewServer(&engineio.Options{
		Transports: []transport.Transport{
			&polling.Transport{
				CheckOrigin: allowOriginFunc,
			},
			&websocket.Transport{
				CheckOrigin: allowOriginFunc,
			},
		},
	})
}

func createBoundingBox(bb string) geography.BoundingBox {
	// TODO: errors and parse all 4 points
	f, _ := strconv.ParseFloat(bb, 64)
	return geography.BoundingBox{
		SouthWest: geography.LatLng{Lat: f, Lon: f},
		NorthEast: geography.LatLng{Lat: f, Lon: f},
	}
}
