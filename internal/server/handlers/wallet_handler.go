package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/geooooo/itk-go-test/internal/db"
	"github.com/geooooo/itk-go-test/internal/logger"
	"github.com/geooooo/itk-go-test/internal/server/handlers/models"
	"github.com/geooooo/itk-go-test/internal/server/helpers"
)

func HandleWallet(logger logger.ILogger, db db.IDb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		formattedRequestUri := helpers.FormatRequestUri(r)

		defer r.Body.Close()
		buf, err := io.ReadAll(r.Body)
		if err != nil {
			logger.Error(fmt.Errorf("%s - error on read request body: %w", formattedRequestUri, err))
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		requestData := string(buf)

		wr := &models.WalletRequest{}
		if err := json.Unmarshal(buf, wr); err != nil {
			logger.Error(fmt.Errorf("%s - error on parse request body '%s': %w", formattedRequestUri, requestData, err))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if !wr.IsValid() {
			logger.Error(fmt.Errorf("%s - error on validation request data '%s'", formattedRequestUri, requestData))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		logger.Log(fmt.Sprintf("%s - got data '%s'", formattedRequestUri, requestData))

		if err := db.UpdateWalletBalance(wr.Id, wr.Amount, wr.Operation); errors.Is(err, sql.ErrNoRows) {
			logger.Error(fmt.Errorf("%s - error on update balance in db, not exists uuid '%s': %w", formattedRequestUri, wr.Id, err))
			w.WriteHeader(http.StatusBadRequest)
			return
		} else if err != nil {
			logger.Error(fmt.Errorf("%s - error on update balance in db: %w", formattedRequestUri, err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		logger.Log(fmt.Sprintf("%s - ok response", formattedRequestUri))

		w.WriteHeader(http.StatusOK)
	}
}
