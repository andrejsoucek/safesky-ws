package websocket

import (
	"fmt"
	"net/http"

	"github.com/andrejsoucek/safesky-ws/authentication"
	"github.com/andrejsoucek/safesky-ws/geography"
	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/polling"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
	log "github.com/sirupsen/logrus"
)

func Listen(clients *Clients) {
	server := createServer()
	server.OnEvent("/", "auth", func(conn socketio.Conn, data string) {
		credentials, err := authentication.CreateCredentialsFromJson(data)
		if err != nil {
			conn.Emit("error", "Unexpected credentials.")
			return
		}
		user, err := authentication.Authenticate(credentials)
		if err != nil {
			conn.Emit("error", err.Error())
			return
		}
		if user.Premium != true {
			conn.Emit("error", "No premium")
		}
		conn.SetContext(user)
		conn.Emit("success", "OK")
	})

	server.OnEvent("/", "bb", func(conn socketio.Conn, data string) {
		if conn.Context() == nil {
			conn.Emit("error", "Authenticate first.")
			return
		}
		log.Info("BoundingBox Received", data)
		clients.SetBoundingBox(conn, geography.CreateBoundingBoxFromJson(data))
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(conn socketio.Conn, reason string) {
		clients.Remove(conn)
	})

	go server.Serve()
	defer server.Close()

	http.Handle("/socket.io/", server)

	log.Info("Serving at localhost:8000...")
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
