package config

import (
	"strconv"
)

type AppConfig struct {
	dialect        string
	dbhost         string
	dbname         string
	username       string
	password       string
	charset        string
	dbport         int
	connectionPool int
	appServerPort  int
	AppSecret      string
}

func GetAppConfig() *AppConfig {
	return &AppConfig{
		dialect:        "mysql",
		dbhost:         "localhost",
		dbname:         "auction_db",
		username:       "root",
		password:       "changeme",
		charset:        "utf8",
		dbport:         3306,
		connectionPool: 10,
		appServerPort:  8090,
		AppSecret:      "Secret Key",
	}
}

func (cfg *AppConfig) GetDatabaseDialect() string {
	return cfg.dialect
}
func (cfg *AppConfig) GetDatabaseHostname() string {
	return cfg.dbhost
}
func (cfg *AppConfig) GetDatabaseName() string {
	return cfg.dbname
}
func (cfg *AppConfig) GetDatabaseUsername() string {
	return cfg.username
}
func (cfg *AppConfig) GetDatabasePassword() string {
	return cfg.password
}
func (cfg *AppConfig) GetDatabaseCharset() string {
	return cfg.charset
}
func (cfg *AppConfig) GetDatabasePort() string {
	return strconv.Itoa(cfg.dbport)
}
func (cfg *AppConfig) GetConnectionPool() int {
	return cfg.connectionPool
}
func (cfg *AppConfig) GetAppServerPort() string {
	return strconv.Itoa(cfg.appServerPort)
}

func (cfg *AppConfig) GetAppSecret() string {
	return cfg.AppSecret
}
