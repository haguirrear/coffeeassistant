package server

import "github.com/haguirrear/coffeeassistant/server/internal/pages/app"

func (s *Server) Pages() {

	s.RegisterPage(app.GetController())
}
