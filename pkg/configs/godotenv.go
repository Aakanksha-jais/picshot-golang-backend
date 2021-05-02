package configs

import (
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/log"
	"os"

	"github.com/joho/godotenv"
)

type ConfigLoader struct {
	configFolderPath string
	logger           log.Logger
}

func NewConfigLoader(configFolderPath string, logger log.Logger) ConfigLoader {
	configFile := configFolderPath + "/.env"

	err := godotenv.Load(configFile)

	if err != nil {
		logger.Errorf("cannot load config from %s, error: %s", configFile, err)
		return ConfigLoader{}
	}

	logger.Infof("loading configs from  file %s", configFile)
	return ConfigLoader{configFolderPath: configFolderPath, logger: logger}
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
