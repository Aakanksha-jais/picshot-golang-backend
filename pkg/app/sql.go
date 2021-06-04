package app

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/configs"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/log"
)

type SQLClient interface {
	GetDB() *sql.DB
}

type sqlClient struct {
	*sql.DB
	config *SQLConfig
}

func (s sqlClient) GetDB() *sql.DB {
	return s.DB
}

type SQLConfig struct {
	HostName string
	Username string
	Password string
	Port     string
	Database string
}

func getSQLConfigFromEnv(config configs.Config) *SQLConfig {
	return &SQLConfig{
		HostName: config.GetOrDefault("DB_HOST", "localhost"),
		Username: config.GetOrDefault("DB_USER", "root"),
		Password: config.Get("DB_PASSWORD"),
		Port:     config.GetOrDefault("DB_PORT", "3306"),
		Database: config.Get("DB_NAME"),
	}
}

func getSQLConnectionString(config *SQLConfig) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.Username, config.Password, config.HostName, config.Port, config.Database)
}

func GetNewSQLClient(logger log.Logger, config configs.Config) (SQLClient, error) {
	sqlConfig := getSQLConfigFromEnv(config)

	connectionString := getSQLConnectionString(sqlConfig)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		logger.Fatalf("cannot connect to sql %v", err)
		return sqlClient{}, err
	}

	err = db.Ping()
	if err != nil {
		logger.Fatalf("error in pinging mysql client: %v", err)
		return sqlClient{}, err
	}

	logger.Infof("connected to mysql: [%v@%v at port: %v]", sqlConfig.Username, sqlConfig.HostName, sqlConfig.Port)

	return sqlClient{DB: db, config: sqlConfig}, nil
}
