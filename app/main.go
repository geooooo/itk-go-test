package main

import (
	"os"

	"github.com/geooooo/itk-go-test/internal/server"
)

const configPath = "config.env"

// TODO: накидать тестов
// TODO: обернуть в docker и compose
func main() {
	server.RunServer(configPath, os.Stdout)
}
