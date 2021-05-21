package driver

import (
	"database/sql"
	"fmt"

	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/configs"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/log"

	_ "github.com/go-sql-driver/mysql"
)

type SQLConfigs struct {
	HostName string
	Username string
	Password string
	Port     string
	Database string
}

func (c SQLConfigs) ConnectToSQL(logger log.Logger) (*sql.DB, error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local", c.Username, c.Password, c.HostName, c.Port, c.Database)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		logger.Fatalf("cannot connect to sql %v", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		logger.Fatalf("error in pinging mysql client: %v", err)
		return nil, err
	}

	logger.Infof("connected to mysql: %v@%v at port: %v", c.Username, c.HostName, c.Port)
	return db, nil
}

func NewSQLConfigs(config configs.Config) SQLConfigs {
	return SQLConfigs{
		HostName: config.GetOrDefault("DB_HOST", "localhost"),
		Username: config.GetOrDefault("DB_USER", "root"),
		Password: config.Get("DB_PASSWORD"),
		Port:     config.GetOrDefault("DB_PORT", "3306"),
		Database: config.Get("DB_NAME"),
	}
}
