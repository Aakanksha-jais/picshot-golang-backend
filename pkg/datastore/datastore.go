package datastore

import "github.com/Aakanksha-jais/picshot-golang-backend/pkg/app"

type DataStore struct {
	Mongo app.MongoDB
	SQL   SQLClient
	S3    app.AWSS3
}
