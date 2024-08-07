package server

import (
	"net/http"

	"github.com/haguirrear/coffeeassistant/server/internal/server/middleware"
	"github.com/haguirrear/coffeeassistant/server/pkg/config"
	"github.com/haguirrear/coffeeassistant/server/pkg/inertia"
	"github.com/haguirrear/coffeeassistant/server/pkg/logger"
	"github.com/rs/zerolog"
)

type Server struct {
	mux   *http.ServeMux
	debug bool
}

func NewServer() *Server {
	return &Server{
		mux:   http.NewServeMux(),
		debug: config.Conf.IsDev,
	}
}

func (s *Server) Setup() {
	s.mux.HandleFunc("/health", healthHandler)
	s.mux.HandleFunc("/favicon.svg", faviconHandler)

	inertia.FileServer(s.mux)
	s.Pages()
}

func (s *Server) RegisterHandler(pattern string, handler http.Handler) {
	handler = s.globalMiddleware(handler)

	s.mux.Handle(pattern, handler)
}

func (s *Server) RegisterPage(pattern string, handler http.Handler) {
	handler = s.globalMiddleware(handler)
	handler = Use(handler, inertia.Middleware)

	s.mux.Handle(pattern, handler)
}

func (s *Server) globalMiddleware(handler http.Handler) http.Handler {
	return Use(handler, middleware.LogMiddleware, middleware.RecoverMiddleware)
}

func (s *Server) Start(addr string) error {
	s.Setup()

	// change to read env
	zerolog.SetGlobalLevel(zerolog.TraceLevel)

	if addr == "" {
		addr = "localhost:8000"
	}

	logger.Logger.Info().Msgf("Server started at %s", addr)

	return http.ListenAndServe(addr, s.mux)
}

func Use(handler http.Handler, middlewareFuncs ...func(http.Handler) http.Handler) http.Handler {
	for _, mid := range middlewareFuncs {
		handler = mid(handler)
	}

	return handler
}
