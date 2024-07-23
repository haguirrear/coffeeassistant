package server

import (
	"net/http"

	"github.com/haguirrear/coffeeassistant/server/pkg/logger"
)

func RecoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				logger.Logger.Error().Any("panic", r).Msg("panic error")
			}
		}()

		next.ServeHTTP(w, r)
	})
}
