package inertia

import (
	"embed"
	"net/http"

	"github.com/haguirrear/coffeeassistant/server/pkg/config"
	inertiago "github.com/petaki/inertia-go"
)

//go:embed index.html
var templateFS embed.FS

const (
	rootTemplate = "index.html"
)

var Manager = NewInertiaManager()

func NewInertiaManager() *inertiago.Inertia {
	manager := inertiago.NewWithFS(config.Conf.BaseURL, rootTemplate, config.Conf.Version, templateFS)

	return manager
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := Manager.WithViewData(r.Context(), "is_prod", config.Conf.IsProd)
		ctx = Manager.WithViewData(ctx, "title", "app")

		if config.Conf.IsProd {
			manifest := readManifest()
			entry := getEntryManifest(manifest, config.InertiaEntry)
			ctx = Manager.WithViewData(ctx, "manifest", entry)
		}

		nextWithInertia := Manager.Middleware(next)
		nextWithInertia.ServeHTTP(w, r.WithContext(ctx))
	})
}
