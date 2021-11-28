package config

import (
	"os"
	"strconv"
	"time"

	"github.com/andrejsoucek/safesky-ws/geography"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	SafeSkyUpdateInterval time.Duration
	SafeSkyApiKey         string
	SafeSkyApiUrl         string
	SafeSkyMaxAlt         string
	SafeSkyBB             geography.BoundingBox
}

func GetConfig(name string) Config {
	err := godotenv.Load(name)
	if err != nil {
		log.Fatal(err)
	}

	updateInt, err := strconv.Atoi(os.Getenv("SAFESKY_UPDATE_INTERVAL_MS"))
	if err != nil {
		log.Fatal(err)
	}

	swLat, err := strconv.ParseFloat(os.Getenv("SAFESKY_BB_SW_LAT"), 64)
	if err != nil {
		log.Fatal(err)
	}
	swLon, err := strconv.ParseFloat(os.Getenv("SAFESKY_BB_SW_LON"), 64)
	if err != nil {
		log.Fatal(err)
	}
	neLat, err := strconv.ParseFloat(os.Getenv("SAFESKY_BB_NE_LAT"), 64)
	if err != nil {
		log.Fatal(err)
	}
	neLon, err := strconv.ParseFloat(os.Getenv("SAFESKY_BB_NE_LON"), 64)
	if err != nil {
		log.Fatal(err)
	}

	return Config{
		SafeSkyUpdateInterval: time.Duration(updateInt),
		SafeSkyApiUrl:         os.Getenv("SAFESKY_API_URL"),
		SafeSkyApiKey:         os.Getenv("SAFESKY_API_KEY"),
		SafeSkyMaxAlt:         os.Getenv("SAFESKY_MAX_ALT"),
		SafeSkyBB: geography.BoundingBox{
			SouthWest: geography.LatLng{
				Lat: swLat,
				Lng: swLon,
			},
			NorthEast: geography.LatLng{
				Lat: neLat,
				Lng: neLon,
			},
		},
	}
}
