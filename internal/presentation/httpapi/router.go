package httpapi

import (
	"api-service-template/internal/app"
	"api-service-template/internal/option"
)

func (s *Server) router(opt *option.Options) {
	application := app.NewApplication(opt)
	s.user(application)
	// router end
}

func (s *Server) user(application *app.Application) {
	c := newUserController(application)
	s.g.POST("/users", wrapHandler(c.Create))
	s.g.DELETE("/users", wrapHandler(c.Delete))
	s.g.PUT("/users", wrapHandler(c.Update))
	s.g.GET("/users", wrapHandler(c.List))
	s.g.GET("/users/detail", wrapHandler(c.Detail))
}
