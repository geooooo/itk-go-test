package handlers

import (
	"net/http"

	"github.com/geooooo/itk-go-test/internal/logger"
)

func HandleWallets(logger *logger.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Log("do something")
	}
}
