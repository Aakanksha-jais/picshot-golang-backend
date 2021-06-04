package account

import (
	"os"
	"testing"

	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/app"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/configs"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/log"
)

var a *app.App

func TestMain(m *testing.M) {
	os.Setenv("ENV", "test")

	testLogger := log.NewLogger()
	testConfigs := configs.NewConfigLoader("../../configs")
	db, _ := app.GetNewSQLClient(testLogger, testConfigs)

	a = &app.App{Logger: testLogger, Config: testConfigs, DataStore: app.DataStore{SQL: db}}

	os.Exit(m.Run())
}
