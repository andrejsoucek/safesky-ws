package main

import (
	"flag"
	"time"

	"github.com/andrejsoucek/safesky-ws/aircraft"
	"github.com/andrejsoucek/safesky-ws/config"
	"github.com/andrejsoucek/safesky-ws/geography"
	"github.com/andrejsoucek/safesky-ws/safesky"
	"github.com/andrejsoucek/safesky-ws/websocket"
	socketio "github.com/googollee/go-socket.io"
	log "github.com/sirupsen/logrus"
)

type Options struct {
	ConfigFile string
}

func doEvery(
	d time.Duration,
	f func(cfg config.Config) ([]aircraft.Aircraft, error),
	aircrafts []aircraft.Aircraft,
	clients *websocket.Clients,
	cfg config.Config,
) {
	for range time.Tick(d) {
		if len(clients.Connections) == 0 {
			continue
		}
		data, err := f(cfg)
		if err == nil {
			aircrafts = data
		} else {
			log.Error("Fetch airplanes", err)
		}

		clients.EmitAircrafts(aircrafts)
	}
}

func parseArgs() Options {
	var opt Options

	flag.StringVar(&opt.ConfigFile, "env-file", "config.env", ".env file with configuration")
	flag.Parse()
	return opt
}

func main() {
	options := parseArgs()
	cfg := config.GetConfig(options.ConfigFile)
	aircrafts := []aircraft.Aircraft{}
	clients := &websocket.Clients{
		Connections: map[socketio.Conn]geography.BoundingBox{},
	}
	go websocket.Listen(clients)
	doEvery(cfg.SafeSkyUpdateInterval*time.Millisecond, safesky.GetAircrafts, aircrafts, clients, cfg)
}
