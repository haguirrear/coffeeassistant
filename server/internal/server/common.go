package server

import (
	"net/http"

	"github.com/haguirrear/coffeeassistant/server/internal/server/srverr"
	"github.com/haguirrear/coffeeassistant/server/internal/server/srvwrite"
	"github.com/haguirrear/coffeeassistant/server/pkg/inertia"
)

var healthHandler = http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
	if err := srvwrite.JSON(w, http.StatusOK, map[string]any{"status": "OK"}); err != nil {
		srverr.HandleErr(w, err)
	}
})

var faviconHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	inertia.ServePublicAsset(w, r, "favicon.svg")
})
