package configs

import (
	"fmt"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/errors"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/log"
	"os"

	"github.com/joho/godotenv"
)

type ConfigLoader struct {
}

func NewConfigLoader(confLocation string) (ConfigLoader, error) {
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
		return ConfigLoader{}, errors.Error{Message: fmt.Sprintf("cannot load configs from folder: %s", confLocation)}
	}

	return ConfigLoader{}, nil
}

func (c ConfigLoader) Get(key string) string {
	return os.Getenv(key)
}

func (c ConfigLoader) GetOrDefault(key, defaultVal string) string {
	val := os.Getenv(key)
	if val != "" {
		return val
	}

	return defaultVal
}
