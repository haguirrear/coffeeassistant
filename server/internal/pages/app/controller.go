package app

import (
	"net/http"

	"github.com/haguirrear/coffeeassistant/server/pkg/inertia"
	"github.com/haguirrear/coffeeassistant/server/pkg/logger"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if err := inertia.Manager.Render(w, r, "App", nil); err != nil {
		logger.Logger.Error().Err(err).Msg("")
	}
	// srverr.HandleErr(w, err)
}

func GetController() (string, http.Handler) {
	return "/app", http.HandlerFunc(handler)
}
