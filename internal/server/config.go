package server

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/geooooo/itk-go-test/internal/logger"
)

type config struct {
	host       string
	port       string
	apiVersion string

	logger *logger.Logger
}

func newConfig(logger *logger.Logger) *config {
	return &config{
		logger: logger,
	}
}

func (c *config) readFromFile(path string) {
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
			c.apiVersion = value
		default:
			c.logger.Error(fmt.Errorf("unexpected config line '%s'", line))
		}
	}
}

func (c *config) addr() string {
	return fmt.Sprintf("%s:%s", c.host, c.port)
}
