package driver

import (
	"context"
	"fmt"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/configs"
	"time"

	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/mongo"
)

type MongoConfigs struct {
	Host     string
	Port     string
	Database string
}

func (c MongoConfigs) ConnectToMongo() (*mongo.Database, error) {
	connectionString := fmt.Sprintf("mongodb://%s:%v/", c.Host, c.Port)

	clientOptions := options.Client().ApplyURI(connectionString)

	const defaultMongoTimeout = 3
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(defaultMongoTimeout)*time.Second)

	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return client.Database(c.Database), nil
}

func NewMongoConfigs(config configs.ConfigLoader) MongoConfigs {
	return MongoConfigs{
		Host:     config.Get("MONGO_HOST"),
		Port:     config.Get("MONGO_PORT"),
		Database: config.GetOrDefault("MONGO_DATABASE", "picshot"),
	}
}
