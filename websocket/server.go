package websocket

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/andrejsoucek/safesky-ws/aircraft"
	"github.com/andrejsoucek/safesky-ws/authentication"
	"github.com/andrejsoucek/safesky-ws/geography"
	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/polling"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
)

func Listen(
	onBBUpdate func(id int, conn socketio.Conn, bb geography.BoundingBox),
	onDisconnect func(id int),
) {
	server := createServer()
	server.OnEvent("/", "auth", func(s socketio.Conn, data string) {
		credentials, err := authentication.CreateCredentialsFromJson(data)
		if err != nil {
			s.Emit("error", "Unexpected credentials.")
			return
		}
		user, err := authentication.Authenticate(credentials)
		if err != nil {
			s.Emit("error", err.Error())
			return
		}
		s.SetContext(user.UserId)
		s.Emit("success", "OK")
	})

	server.OnEvent("/", "bb", func(s socketio.Conn, data string) {
		if s.Context() == nil {
			s.Emit("error", "Authenticate first.")
			return
		}
		onBBUpdate(s.Context().(int), s, geography.CreateBoundingBoxFromJson(data))
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("closed", reason)
		onDisconnect(s.Context().(int))
	})

	go server.Serve()
	defer server.Close()

	http.Handle("/socket.io/", server)

	log.Println("Serving at localhost:8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func EmitAircrafts(aircrafts []aircraft.Aircraft, clients map[int]Client) {
	for _, client := range clients {
		var visibleAircrafts []aircraft.Aircraft
		for _, aircraft := range aircrafts {
			if geography.IsInBounds(client.Bb, aircraft.LatLng) {
				visibleAircrafts = append(visibleAircrafts, aircraft)
			}
		}

		json, err := json.Marshal(visibleAircrafts)
		if err != nil {
			fmt.Println(err)
			continue
		}
		client.Conn.Emit("aircrafts", string(json))
	}
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
