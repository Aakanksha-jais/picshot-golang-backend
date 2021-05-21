package configs

import (
	"os"

	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/log"

	"github.com/joho/godotenv"
)

type config struct {
	log          log.Logger
	confLocation string
}

func NewConfigLoader(confLocation string) config {
	log := log.NewLogger()

	defaultFile := confLocation + "/.env"

	env := os.Getenv("ENV")
	if env == "" {
		env = "local"
	}
	overrideFile := confLocation + "/." + env + ".env"

	err1 := godotenv.Load(overrideFile)
	if err1 == nil {
		log.Infof("loaded config from file: %s", overrideFile)
	}

	err2 := godotenv.Load(defaultFile)
	if err2 == nil {
		log.Infof("loaded config from file: %s", defaultFile)
	}

	if err1 != nil && err2 != nil {
		log.Fatalf("cannot load configs from folder: %s", confLocation)
	}

	return config{log: log, confLocation: confLocation}
}

func (c config) Get(key string) string {
	return os.Getenv(key)
}

func (c config) GetOrDefault(key, defaultVal string) string {
	val := os.Getenv(key)
	if val != "" {
		return val
	}

	return defaultVal
}
