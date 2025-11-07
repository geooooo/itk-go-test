package config

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/geooooo/itk-go-test/internal/logger"
)

type Config struct {
	host string
	port string

	ApiVersion string

	DbReset    bool
	dbName     string
	dbUser     string
	dbPassword string

	logger logger.ILogger
}

func NewConfig(logger logger.ILogger) *Config {
	return &Config{
		logger: logger,
	}
}

func (c *Config) ReadFromFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		c.logger.Error(err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			c.logger.Error(err)
			os.Exit(1)
		}

		line := scanner.Text()
		parts := strings.Split(line, "=")
		key, value := parts[0], parts[1]

		switch key {
		case "host":
			c.host = value
		case "port":
			c.port = value
		case "apiVersion":
			c.ApiVersion = value
		case "dbUser":
			c.dbUser = value
		case "dbPassword":
			c.dbPassword = value
		case "dbName":
			c.dbName = value
		case "dbReset":
			c.DbReset = value == "yes"
		default:
			c.logger.Error(fmt.Errorf("unexpected config line '%s'", line))
		}
	}
}

func (c *Config) Addr() string {
	return fmt.Sprintf("%s:%s", c.host, c.port)
}

func (c *Config) ConnStr() string {
	return fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", c.dbUser, c.dbPassword, c.dbName)
}
