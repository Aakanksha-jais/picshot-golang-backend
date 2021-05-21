package app

import (
	"os"

	"github.com/Aakanksha-jais/picshot-golang-backend/driver"

	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/configs"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/log"
)

type App struct {
	server *server
	Logger log.Logger
	Config configs.Config
	DataStore
}

func New() *App {
	app := &App{}

	// initialize configs
	app.readConfig()

	// initialize Logger
	app.Logger = log.NewLogger()

	app.initializeStores(app.Config)

	// initialize server
	app.server = NewServer(app)

	return app
}

// GET adds a Handler for http GET method for a route pattern.
func (a *App) GET(pattern string, handler Handler) {
	a.add("GET", pattern, handler)
}

// PUT adds a Handler for http PUT method for a route pattern.
func (a *App) PUT(pattern string, handler Handler) {
	a.add("PUT", pattern, handler)
}

// POST adds a Handler for http POST method for a route pattern.
func (a *App) POST(pattern string, handler Handler) {
	a.add("POST", pattern, handler)
}

// DELETE adds a Handler for http DELETE method for a route pattern.
func (a *App) DELETE(pattern string, handler Handler) {
	a.add("DELETE", pattern, handler)
}

func (a *App) add(method, pattern string, h Handler) {
	a.server.Router.Add(method, pattern, h)
}

func (a App) Start() {
	a.server.Start(a.Logger)
}

// readConfig reads the configuration from the default location.
func (a *App) readConfig() {
	var configLocation string
	if _, err := os.Stat("./configs"); err == nil {
		configLocation = "./configs"
	}

	a.Config = configs.NewConfigLoader(configLocation)
}

func (a *App) initializeStores(config configs.Config) {
	mongoDB, err := driver.NewMongoConfigs(config).ConnectToMongo(a.Logger)
	if err != nil {
		a.Logger.Errorf("cannot connect to mongo: %s", err.Error())
	}

	sqlDB, err := driver.NewSQLConfigs(config).ConnectToSQL(a.Logger)
	if err != nil {
		a.Logger.Errorf("cannot connect to mysql: %s", err.Error())
	}

	a.DataStore = DataStore{MongoDB: mongoDB, SQLDB: sqlDB}
}
