// middlewares/LoggingMiddleware.go

package middlewares

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

var CartMux = http.NewServeMux()

type LoggingMiddleware struct {
	Next http.Handler
}

func (lm LoggingMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if lm.Next == nil {
		lm.Next = CartMux
	}

	slog.Info(fmt.Sprintf("Received %v request on route: %v", r.Method, r.URL.Path))
	now := time.Now()

	lm.Next.ServeHTTP(w, r)

	slog.Info(fmt.Sprintf("Response generated for %v request on route: %v. Duration: %v", r.Method, r.URL.Path, time.Since(now)))
}
