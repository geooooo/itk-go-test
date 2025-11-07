package middlewares

import (
	"net/http"

	"github.com/geooooo/itk-go-test/internal/logger"
	"github.com/geooooo/itk-go-test/internal/server/helpers"
)

func RequestLogMiddleware(next http.HandlerFunc, logger logger.ILogger) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Log(helpers.FormatRequestUri(r))

		next.ServeHTTP(w, r)
	})
}
