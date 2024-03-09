// middlewares/LoggingMiddleware.go

package middlewares

import (
	"net/http"
)

type LoggingMiddleware struct {
	next http.Handler
}
