package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/geooooo/itk-go-test/internal/logger"
	"github.com/geooooo/itk-go-test/internal/server/helpers"
)

func HandleWallets(endpoint string, logger logger.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		formattedRequestUri := helpers.FormatRequestUri(r)
		parts := strings.Split(r.URL.Path, endpoint)
		uuidParam := parts[1]

		if len(uuidParam) == 0. || strings.Contains(uuidParam, "/") {
			logger.Error(fmt.Errorf("%s error on request format", formattedRequestUri))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		logger.Log(fmt.Sprintf("%s got data '%s'", formattedRequestUri, uuidParam))

		// todo

		w.WriteHeader(http.StatusOK)
	}
}
