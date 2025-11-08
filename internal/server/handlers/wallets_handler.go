package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/geooooo/itk-go-test/internal/db"
	"github.com/geooooo/itk-go-test/internal/logger"
	"github.com/geooooo/itk-go-test/internal/server/handlers/models"
	"github.com/geooooo/itk-go-test/internal/server/helpers"
)

func HandleWallets(endpoint string, logger logger.ILogger, db db.IDb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		formattedRequestUri := helpers.FormatRequestUri(r)
		parts := strings.Split(r.URL.Path, endpoint)
		uuidParam := parts[1]

		if len(uuidParam) == 0 || strings.Contains(uuidParam, "/") {
			logger.Error(fmt.Errorf("%s - error on request format", formattedRequestUri))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		logger.Log(fmt.Sprintf("%s - got data '%s'", formattedRequestUri, uuidParam))

		sum, err := db.GetWalletBalance(uuidParam)
		if errors.Is(err, sql.ErrNoRows) {
			logger.Error(fmt.Errorf("%s - error on get balance from db, not exists uuid '%s': %w", formattedRequestUri, uuidParam, err))
			w.WriteHeader(http.StatusBadRequest)
			return
		} else if err != nil {
			logger.Error(fmt.Errorf("%s - error on get balance from db: %w", formattedRequestUri, err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		wr := &models.WalletsResponse{
			Sum: sum,
		}
		buf, err := json.Marshal(wr)
		if err != nil {
			logger.Error(fmt.Errorf("%s - error on serialize response: %w", formattedRequestUri, err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		logger.Log(fmt.Sprintf("%s - ok response data '%s'", formattedRequestUri, string(buf)))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(buf)
	}
}
