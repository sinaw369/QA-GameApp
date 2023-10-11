package httpserver

import (
	"Q/A-GameApp/config"
	"Q/A-GameApp/delivery/httpserver/userhandler"
	"Q/A-GameApp/service/authservice"
	"Q/A-GameApp/service/userservise"
	"Q/A-GameApp/validator/uservalidator"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	config      config.Config
	userHandler userhandler.Handler
}

func New(config config.Config, authSvc authservice.Service, userSvc userservise.Service, uservalidator uservalidator.Validator) Server {
	return Server{
		config:      config,
		userHandler: userhandler.New(authSvc, userSvc, uservalidator),
	}
}
func (s Server) Server() {
	e := echo.New()
	//middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	//routes
	e.GET("/health-check", s.healthCheck)
	s.userHandler.SetUserRoute(e)
	// start server
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", s.config.HTTPServer.Port)))
}
