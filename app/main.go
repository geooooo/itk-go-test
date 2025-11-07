package main

import (
	"os"

	"github.com/geooooo/itk-go-test/internal/server"
)

const configPath = "config.env"

func main() {
	server.RunServer(configPath, os.Stdout)
}
