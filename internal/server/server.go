package server

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/geooooo/itk-go-test/internal/logger"
	"github.com/geooooo/itk-go-test/internal/server/api"
	"github.com/geooooo/itk-go-test/internal/server/handlers"
	"github.com/geooooo/itk-go-test/internal/server/middlewares"
)

func RunServer(configPath string, logOutput *os.File) {
	logger := logger.NewLogger(logOutput)

	config := newConfig(logger)
	config.readFromFile(configPath)

	initHandlers(config.apiVersion, logger)

	addr := config.addr()
	logger.Log(fmt.Sprintf("starting server %s", addr))

	err := http.ListenAndServe(addr, nil)
	if errors.Is(err, http.ErrServerClosed) {
		logger.Log("server stopped")
	} else {
		logger.Error(err)
		os.Exit(1)
	}
}

func initHandlers(apiVersion string, logger *logger.Logger) {
	http.HandleFunc(
		api.ConfigureEndpoint(api.Wallet, apiVersion),
		middlewares.RequestLogMiddleware(
			handlers.HandleWallet(logger),
			logger,
		),
	)

	http.HandleFunc(
		api.ConfigureEndpoint(api.Wallets, apiVersion),
		middlewares.RequestLogMiddleware(
			handlers.HandleWallets(logger),
			logger,
		),
	)
}
