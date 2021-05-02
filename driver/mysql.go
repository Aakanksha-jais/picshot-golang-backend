package driver

import (
	"database/sql"
	"fmt"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/configs"

	_ "github.com/go-sql-driver/mysql"
)

type SQLConfigs struct {
	Host     string
	Username string
	Password string
	Port     string
	Database string
}

func (c SQLConfigs) ConnectToSQL() (*sql.DB, error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", c.Username, c.Password, c.Host, c.Port, c.Database)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func NewSQLConfigs(config configs.ConfigLoader) SQLConfigs {
	return SQLConfigs{
		Host:     config.Get("DB_HOST"),
		Username: config.Get("DB_USER"),
		Password: config.Get("DB_PASSWORD"),
		Port:     config.Get("DB_PORT"),
		Database: config.Get("DATABASE"),
	}
}
