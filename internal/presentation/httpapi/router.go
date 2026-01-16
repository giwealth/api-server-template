package httpapi

import (
	"api-service-template/internal/option"
)

func (s *Server) router(opt *option.Options) {
	s.user(opt)
	s.administrator(opt)
	// router end
}

func (s *Server) user(opt *option.Options) {
	c := newUserController(opt)
	s.g.POST("/users", wrapHandler(c.Create))
	s.g.DELETE("/users/:id", wrapHandler(c.Delete))
	s.g.PUT("/users/:id", wrapHandler(c.Update))
	s.g.GET("/users/:id", wrapHandler(c.Find))
	s.g.GET("/users", wrapHandler(c.List))
}

func (s *Server) administrator(opt *option.Options) {
	c := newAdministratorController(opt)
	s.g.POST("/administrators", wrapHandler(c.Create))
	s.g.DELETE("/administrators/:id", wrapHandler(c.Delete))
	s.g.PUT("/administrators/:id", wrapHandler(c.Update))
	s.g.GET("/administrators/:id", wrapHandler(c.Find))
	s.g.GET("/administrators", wrapHandler(c.List))
}
