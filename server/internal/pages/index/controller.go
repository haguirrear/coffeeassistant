package index

import (
	"net/http"

	"github.com/haguirrear/coffeeassistant/server/pkg/inertia"
	"github.com/haguirrear/coffeeassistant/server/pkg/logger"
)

func handler(w http.ResponseWriter, r *http.Request) {
	r = r.WithContext(inertia.Manager.WithViewData(r.Context(), "title", "Coffe Assistant"))
	if err := inertia.Manager.Render(w, r, "Index", nil); err != nil {
		logger.Logger.Error().Err(err).Msg("")

	}

}

func GetController() (string, http.Handler) {
	return "/", http.HandlerFunc(handler)
}
