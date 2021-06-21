package blog

import (
	"os"
	"testing"

	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/datastore"

	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/app"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/configs"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/log"
)

//nolint
var a *app.App

func TestMain(m *testing.M) {
	os.Setenv("ENV", "test")

	testLogger := log.NewLogger()
	testConfigs := configs.NewConfigLoader("../../configs")
	mongoDB, _ := datastore.GetNewMongoDB(testLogger, testConfigs)

	a = &app.App{Logger: testLogger, Config: testConfigs, DataStore: datastore.DataStore{Mongo: mongoDB}}

	os.Exit(m.Run())
}
