package helpers

import (
	"fmt"
	"net/http"
)

func FormatRequestUri(r *http.Request) string {
	return fmt.Sprintf("request api - %s %s", r.Method, r.RequestURI)
}
