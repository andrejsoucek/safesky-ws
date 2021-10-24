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

func Listen(clients map[socketio.Conn]geography.BoundingBox) {
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
		if user.Premium != true {
			s.Emit("error", "No premium")
		}
		s.SetContext(user)
		s.Emit("success", "OK")
	})

	server.OnEvent("/", "bb", func(s socketio.Conn, data string) {
		if s.Context() == nil {
			s.Emit("error", "Authenticate first.")
			return
		}
		user := s.Context().(authentication.User)
		if user.Premium != true {
			s.Emit("error", "User is not premium.")
			return
		}
		clients[s] = geography.CreateBoundingBoxFromJson(data)
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		delete(clients, s)
	})

	go server.Serve()
	defer server.Close()

	http.Handle("/socket.io/", server)

	log.Println("Serving at localhost:8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func EmitAircrafts(aircrafts []aircraft.Aircraft, clients map[socketio.Conn]geography.BoundingBox) {
	for conn, bb := range clients {
		var visibleAircrafts []aircraft.Aircraft
		for _, aircraft := range aircrafts {
			if geography.IsInBounds(bb, aircraft.LatLng) {
				visibleAircrafts = append(visibleAircrafts, aircraft)
			}
		}

		json, err := json.Marshal(visibleAircrafts)
		if err != nil {
			fmt.Println(err)
			continue
		}
		conn.Emit("aircrafts", string(json))
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
