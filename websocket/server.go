package websocket

import (
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

func Listen(clients *Clients, certPath string, keyPath string) {
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
			conn.Close()
			return
		}
		if user.Premium != true {
			conn.Emit("error", "No premium")
			conn.Close()
			return
		}
		log.Info("client successfully authenticated, ID: ", conn.ID())

		conn.SetContext(user)
		conn.Emit("success", "OK")
	})

	server.OnEvent("/", "bb", func(conn socketio.Conn, data string) {
		if conn.Context() == nil {
			conn.Emit("error", "Authenticate first.")
			conn.Close()
			return
		}
		clients.SetBoundingBox(conn, geography.CreateBoundingBoxFromJson(data))
	})

	server.OnError("/", func(conn socketio.Conn, e error) {
		log.Info("error:", e)
	})

	server.OnConnect("/", func(conn socketio.Conn) error {
		log.Info("new client connected, ID: ", conn.ID())
		return nil
	})

	server.OnDisconnect("/", func(conn socketio.Conn, reason string) {
		conn.Close()
		clients.Remove(conn)
		log.Info("client disconnected, ID: ", conn.ID(), ", current clients: ", clients)
	})

	go server.Serve()
	defer server.Close()

	http.Handle("/socket.io/", server)

	log.Info("Serving at localhost:4433...")
	log.Fatal(http.ListenAndServeTLS(
		":4433",
		certPath,
		keyPath,
		nil,
	))
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
