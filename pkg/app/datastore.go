package app

import (
	"database/sql"

	"go.mongodb.org/mongo-driver/mongo"
)

type DataStore struct {
	MongoDB *mongo.Database
	SQLDB   *sql.DB
}
