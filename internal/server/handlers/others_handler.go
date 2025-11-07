package handlers

import (
	"net/http"

	"github.com/geooooo/itk-go-test/internal/logger"
)

func HandleOthers(logger logger.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	}
}
