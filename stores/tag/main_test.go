package tag

import (
	"os"
	"testing"

	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/app"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/configs"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/log"
)

// nolint:gochecknoglobals //global var needed for tests
var a *app.App

func TestMain(m *testing.M) {
	os.Setenv("ENV", "test")

	testLogger := log.NewLogger()
	testConfigs := configs.NewConfigLoader("../../configs")
	mongoDB, _ := app.GetNewMongoDB(testLogger, testConfigs)

	a = &app.App{Logger: testLogger, Config: testConfigs, DataStore: app.DataStore{Mongo: mongoDB}}

	os.Exit(m.Run())
}
