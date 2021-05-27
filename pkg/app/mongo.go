package app

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/configs"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB interface {
	DB() *mongo.Database
}

// MongoConfig holds the configurations for Mongo Connectivity
type MongoConfig struct {
	HostName    string
	Port        string
	Username    string
	Password    string
	Database    string
	SSL         bool
	RetryWrites bool
}

type mongoDB struct {
	*mongo.Database
	config *MongoConfig
	logger log.Logger
}

func (db mongoDB) DB() *mongo.Database {
	return db.Database
}

func getMongoConnectionString(config *MongoConfig) string {
	if config.Username == "" || config.Password == "" {
		return fmt.Sprintf("mongodb://%s:%v/", config.HostName, config.Port)
	}

	return fmt.Sprintf("mongodb://%v:%v@%v:%v/?ssl=%v&retrywrites=%v", config.Username, config.Password, config.HostName, config.Port, config.SSL, config.RetryWrites)
}

func getBoolEnv(config configs.Config, varName string) (bool, error) {
	val := config.Get(varName)
	if val == "" {
		return false, nil
	}

	return strconv.ParseBool(val)
}

func getMongoConfigFromEnv(logger log.Logger, config configs.Config) *MongoConfig {
	enableSSL, err := getBoolEnv(config, "MONGO_DB_ENABLE_SSL")
	if err != nil {
		logger.Warnf("error in reading bool value for mongo enable ssl: %vs", err.Error())
	}

	retryWrites, err := getBoolEnv(config, "MONGO_DB_RETRY_WRITES")
	if err != nil {
		logger.Warnf("error in reading bool value for mongo retry writes: %vs", err.Error())
	}

	return &MongoConfig{
		Username:    config.Get("MONGO_DB_USER"),
		Password:    config.Get("MONGO_DB_PASS"),
		HostName:    config.GetOrDefault("MONGO_DB_HOST", "localhost"),
		Port:        config.GetOrDefault("MONGO_DB_PORT", "27017"),
		Database:    config.Get("MONGO_DB_NAME"),
		SSL:         enableSSL,
		RetryWrites: retryWrites,
	}
}

func GetNewMongoDB(logger log.Logger, config configs.Config) (MongoDB, error) {
	mongoConfig := getMongoConfigFromEnv(logger, config)

	connectionString := getMongoConnectionString(mongoConfig)

	clientOptions := options.Client().ApplyURI(connectionString)

	const defaultMongoTimeout = 10
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(defaultMongoTimeout)*time.Second)

	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		logger.Fatalf("cannot connect to mongo: %v", err)
		return mongoDB{}, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		logger.Fatalf("error in pinging mongo client: %v", err)
		return mongoDB{}, err
	}

	logger.Infof("connected to mongo: [%v@%v at port %v]", mongoConfig.Username, mongoConfig.HostName, mongoConfig.Port)

	return mongoDB{Database: client.Database(mongoConfig.Database), config: mongoConfig}, nil
}
