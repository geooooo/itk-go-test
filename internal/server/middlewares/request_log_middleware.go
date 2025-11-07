package middlewares

import (
	"fmt"
	"net/http"

	"github.com/geooooo/itk-go-test/internal/logger"
)

func RequestLogMiddleware(next http.HandlerFunc, logger *logger.Logger) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Log(fmt.Sprintf("request api - %s %s", r.Method, r.RequestURI))

		next.ServeHTTP(w, r)
	})
}
