package datastore

type DataStore struct {
	Mongo MongoDB
	SQL   SQLClient
	S3    AWSS3
}
