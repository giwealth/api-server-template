package httpapi

import (
	"api-service-template/internal/option"
)

func (s *Server) router(opt *option.Options) {
	s.user(opt)
	// router end
}

func (s *Server) user(opt *option.Options) {
	c := newUserController(opt)
	s.g.POST("/users", wrapHandler(c.Create))
	s.g.DELETE("/users", wrapHandler(c.Delete))
	s.g.PUT("/users", wrapHandler(c.Update))
	s.g.GET("/users", wrapHandler(c.List))
	s.g.GET("/users/detail", wrapHandler(c.Detail))
}
