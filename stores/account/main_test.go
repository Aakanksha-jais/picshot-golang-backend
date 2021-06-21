package account

import (
	"os"
	"testing"

	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/datastore"

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
	db, _ := datastore.GetNewSQLClient(testLogger, testConfigs)

	a = &app.App{Logger: testLogger, Config: testConfigs, DataStore: datastore.DataStore{SQL: db}}

	os.Exit(m.Run())
}
