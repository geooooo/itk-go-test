package config

import (
	"fmt"
	"os"
)

type Config struct {
	host string
	port string

	ApiVersion string

	DbReset    bool
	dbName     string
	dbUser     string
	dbPassword string
}

func NewConfig() *Config {
	return &Config{
		host:       os.Getenv("host"),
		port:       os.Getenv("port"),
		ApiVersion: os.Getenv("apiVersion"),
		DbReset:    os.Getenv("dbReset") == "yes",
		dbName:     os.Getenv("dbName"),
		dbUser:     os.Getenv("dbUser"),
		dbPassword: os.Getenv("dbPassword"),
	}
}

func (c *Config) Addr() string {
	return fmt.Sprintf("%s:%s", c.host, c.port)
}

func (c *Config) ConnStr() string {
	return fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", c.dbUser, c.dbPassword, c.dbName)
}
