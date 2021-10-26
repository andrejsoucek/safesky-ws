package config

import (
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	SafeSkyApiKey string
	SafeSkyApiUrl string
}

func GetConfig(name string) Config {
	err := godotenv.Load(name)
	if err != nil {
		log.Fatal(err)
	}

	return Config{
		SafeSkyApiUrl: os.Getenv("SAFESKY_API_URL"),
		SafeSkyApiKey: os.Getenv("SAFESKY_API_KEY"),
	}
}
