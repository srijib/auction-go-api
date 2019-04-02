package config

import (
	"strconv"
)

type DatabaseConfig struct {
	dbhost         string
	dbname         string
	dbport         int
	connectionPool int
	appServerPort  int
}

func GetDBConfig() *DatabaseConfig {
	return &DatabaseConfig{
		dbhost:         "localhost",
		dbname:         "auction_db",
		dbport:         27017,
		connectionPool: 10,
		appServerPort:  8000,
	}
}

func (cfg *DatabaseConfig) GetDatabaseHostname() string {
	return cfg.dbhost
}
func (cfg *DatabaseConfig) GetDatabaseName() string {
	return cfg.dbname
}
func (cfg *DatabaseConfig) GetDatabasePort() string {
	return strconv.Itoa(cfg.dbport)
}
func (cfg *DatabaseConfig) GetConnectionPool() int {
	return cfg.connectionPool
}
func (cfg *DatabaseConfig) GetAppServerPort() string {
	return strconv.Itoa(cfg.appServerPort)
}
