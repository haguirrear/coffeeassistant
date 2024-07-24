package server

import (
	"github.com/haguirrear/coffeeassistant/server/internal/pages/app"
	"github.com/haguirrear/coffeeassistant/server/internal/pages/index"
)

func (s *Server) Pages() {

	s.RegisterPage(app.GetController())
	s.RegisterPage(index.GetController())
}
