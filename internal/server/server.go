package server

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/geooooo/itk-go-test/internal/config"
	"github.com/geooooo/itk-go-test/internal/db"
	"github.com/geooooo/itk-go-test/internal/logger"
	"github.com/geooooo/itk-go-test/internal/server/api"
	"github.com/geooooo/itk-go-test/internal/server/handlers"
	"github.com/geooooo/itk-go-test/internal/server/middlewares"
)

func RunServer(configPath string, logOutput *os.File) {
	logger := logger.NewLogger(logOutput)
	config := config.NewConfig()

	db, err := db.NewDb(config)
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}

	initHandlers(config.ApiVersion, logger, db)

	addr := config.Addr()
	logger.Log(fmt.Sprintf("starting server %s", addr))

	// TODO: (упрощение) обработка сигналов опущена
	err = http.ListenAndServe(addr, nil)
	if errors.Is(err, http.ErrServerClosed) {
		logger.Log("server stopped")
	} else {
		logger.Error(err)
		os.Exit(1)
	}
}

// TODO: (упрощение) по-хорошему, при bad request надо отдавать доп информацию о произошедшей ошибке клиету
func initHandlers(apiVersion string, logger logger.ILogger, db db.IDb) {
	http.HandleFunc(
		"/",
		middlewares.RequestLogMiddleware(
			handlers.HandleOthers(logger),
			logger,
		),
	)

	http.HandleFunc(
		api.ConfigureEndpoint(api.Wallet, apiVersion),
		middlewares.RequestLogMiddleware(
			handlers.HandleWallet(logger, db),
			logger,
		),
	)

	walletsEndpoint := api.ConfigureEndpoint(api.Wallets, apiVersion)
	http.HandleFunc(
		walletsEndpoint,
		middlewares.RequestLogMiddleware(
			handlers.HandleWallets(walletsEndpoint, logger, db),
			logger,
		),
	)
}
