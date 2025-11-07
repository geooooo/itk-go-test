package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/geooooo/itk-go-test/internal/logger"
	"github.com/geooooo/itk-go-test/internal/server/handlers/models"
	"github.com/geooooo/itk-go-test/internal/server/helpers"
)

func HandleWallet(logger logger.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		formattedRequestUri := helpers.FormatRequestUri(r)

		defer r.Body.Close()
		buf, err := io.ReadAll(r.Body)
		if err != nil {
			logger.Error(fmt.Errorf("%s error on read request body: %w", formattedRequestUri, err))
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		requestData := string(buf)

		wr := &models.WalletRequest{}
		if err := json.Unmarshal(buf, wr); err != nil {
			logger.Error(fmt.Errorf("%s error on parse request body '%s': %w", formattedRequestUri, requestData, err))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if !wr.IsValid() {
			logger.Error(fmt.Errorf("%s error on validation request data '%s'", formattedRequestUri, requestData))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		logger.Log(fmt.Sprintf("%s got data '%s'", formattedRequestUri, requestData))

		// todo

		w.WriteHeader(http.StatusOK)
	}
}
