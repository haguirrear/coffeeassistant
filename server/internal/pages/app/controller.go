package app

import (
	"net/http"

	"github.com/haguirrear/coffeeassistant/server/pkg/inertia"
	"github.com/haguirrear/coffeeassistant/server/pkg/logger"
)

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := inertia.Manager.WithViewData(r.Context(), "title", "Counter")
	newReq := r.WithContext(ctx)
	if err := inertia.Manager.Render(w, newReq, "App", nil); err != nil {
		logger.Logger.Error().Err(err).Msg("")
	}
}

func GetController() (string, http.Handler) {
	return "/app", http.HandlerFunc(handler)
}
