package app

type DataStore struct {
	Mongo MongoDB
	SQL   SQLClient
}
