package driver

import (
	"context"
	"fmt"
	"time"

	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/configs"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/log"

	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/mongo"
)

type MongoConfigs struct {
	Username string
	Password string
	HostName string
	Port     string
	Database string
}

func (c MongoConfigs) ConnectToMongo(logger log.Logger) (*mongo.Database, error) {
	var connectionString string

	if c.Username != "" && c.Password != "" {
		connectionString = fmt.Sprintf("mongodb://%v:%v@%v:%v/", c.Username, c.Password, c.HostName, c.Port)
	} else {
		connectionString = fmt.Sprintf("mongodb://%s:%v/", c.HostName, c.Port)
	}

	clientOptions := options.Client().ApplyURI(connectionString)

	const defaultMongoTimeout = 3
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(defaultMongoTimeout)*time.Second)

	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		logger.Fatalf("cannot connect to  mongo: %v", err)
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		logger.Fatalf("error in pinging mongo client: %v", err)
		return nil, err
	}

	logger.Infof("connected to mongo: %v@%v at port: %v", c.Username, c.HostName, c.Port)
	return client.Database(c.Database), nil
}

func NewMongoConfigs(config configs.Config) MongoConfigs {
	return MongoConfigs{
		Username: config.Get("MONGO_DB_USER"),
		Password: config.Get("MONGO_DB_PASS"),
		HostName: config.GetOrDefault("MONGO_DB_HOST", "localhost"),
		Port:     config.GetOrDefault("MONGO_DB_PORT", "27017"),
		Database: config.Get("MONGO_DB_NAME"),
	}
}
