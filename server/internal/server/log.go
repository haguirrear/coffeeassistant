package server

import (
	"net/http"
	"time"

	"github.com/haguirrear/coffeeassistant/server/pkg/logger"
)

func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		rw := newResponseWrapper(w)

		defer func() {
			l := logger.Logger.Info().Str("Method", r.Method).Str("Path", r.URL.Path).Str("exec_time", time.Since(start).String())

			if rw.statusCode != 0 {
				l = l.Int("code", rw.statusCode)
			}

			l.Msg("")
		}()
		next.ServeHTTP(rw, r)
	})
}

type responseWrapper struct {
	w          http.ResponseWriter
	statusCode int
}

func newResponseWrapper(w http.ResponseWriter) *responseWrapper {
	return &responseWrapper{w: w}
}

func (rw *responseWrapper) WriteHeader(status int) {
	rw.statusCode = status
	rw.w.WriteHeader(status)
}

func (rw *responseWrapper) Header() http.Header {
	return rw.w.Header()
}

func (rw *responseWrapper) Write(b []byte) (int, error) {
	return rw.w.Write(b)
}
