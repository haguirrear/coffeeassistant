package inertia

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	dist "github.com/haguirrear/coffeeassistant"
	"github.com/haguirrear/coffeeassistant/server/internal/server/middleware"
	"github.com/haguirrear/coffeeassistant/server/pkg/config"
	"github.com/haguirrear/coffeeassistant/server/pkg/logger"
)

type viteManifest map[string]ManifestItem

type ManifestItem struct {
	File           string   `json:"file"`
	Name           string   `json:"name"`
	Src            string   `json:"src"`
	IsEntry        bool     `json:"isEntry"`
	IsDynamicEntry bool     `json:"isDynamicEntry"`
	CSS            []string `json:"css"`
	Assets         []string `json:"assets"`
	Imports        []string `json:"imports"`
	DynamicImports []string `json:"dynamicImports"`
}

func readManifest() viteManifest {
	manifest := viteManifest{}
	manBytes, err := dist.DistFS.ReadFile("dist/manifest.json")
	if err != nil {
		logger.Logger.Err(err).Msg("error trying to read manifest.json")

		return manifest
	}

	if err := json.Unmarshal(manBytes, &manifest); err != nil {

		logger.Logger.Err(err).Msg("error trying to unmarshal manifest.json")

		return manifest
	}

	return manifest
}

func getEntryManifest(manifest viteManifest, name string) map[string]any {
	item, found := manifest[name]
	if !found {
		return map[string]any{}
	}

	chunks := getChunks(manifest, item.Imports)
	chunks = append(chunks, getChunks(manifest, item.DynamicImports)...)

	var chunksCSS, chunksFiles []string

	for _, c := range chunks {
		if len(c.CSS) > 0 {
			chunksCSS = append(chunksCSS, c.CSS...)
		}

		if c.File != "" {
			chunksFiles = append(chunksFiles, c.File)
		}

	}

	return map[string]any{
		"css":         addPrefixAll(item.CSS),
		"js":          addPrefix(item.File),
		"chunk_css":   addPrefixAll(chunksCSS),
		"chunk_files": addPrefixAll(chunksFiles),
	}
}

func addPrefix(file string) string {
	return "public/" + file
}

func addPrefixAll(files []string) []string {
	for i, s := range files {
		files[i] = addPrefix(s)
	}

	return files
}

func getChunks(manifest viteManifest, chunkNames []string) []ManifestItem {
	var importedChunks []ManifestItem
	for _, chunkName := range chunkNames {
		chunk, found := manifest[chunkName]
		if !found {
			logger.Logger.Warn().Msgf("chunk %q not found in manifest", chunkName)

			continue
		}

		importedChunks = append(importedChunks, chunk)
	}

	return importedChunks
}

func FileServer(mux *http.ServeMux) {
	if config.Conf.IsDev {
		publicHandler := middleware.LogMiddleware(http.StripPrefix("/public", http.FileServer(http.Dir("./public"))))
		frontendHandler := middleware.LogMiddleware(http.StripPrefix("/public/frontend", http.FileServer(http.Dir("./frontend"))))
		mux.Handle("/public/", publicHandler)
		mux.Handle("/public/frontend/", frontendHandler)
	} else {
		handler := middleware.LogMiddleware(ReplacePrefix("/public", "/dist", http.FileServerFS(dist.DistFS)))
		mux.Handle("/public/", handler)

	}

}

func ReplacePrefix(prefix string, replace string, h http.Handler) http.Handler {
	if prefix == "" {
		return h
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p := strings.TrimPrefix(r.URL.Path, prefix)
		rp := strings.TrimPrefix(r.URL.RawPath, prefix)

		p = replace + p
		rp = replace + rp
		if len(p) < len(r.URL.Path) && (r.URL.RawPath == "" || len(rp) < len(r.URL.RawPath)) {
			r2 := new(http.Request)
			*r2 = *r
			r2.URL = new(url.URL)
			*r2.URL = *r.URL
			r2.URL.Path = p
			r2.URL.RawPath = rp
			h.ServeHTTP(w, r2)
		} else {
			http.NotFound(w, r)
		}
	})
}

func ServePublicAsset(w http.ResponseWriter, r *http.Request, filename string) {
	if config.Conf.IsDev {
		http.ServeFile(w, r, fmt.Sprintf("./public/%s", filename))
	} else {
		http.ServeFileFS(w, r, dist.DistFS, fmt.Sprintf("./dist/%s", filename))
	}
}
