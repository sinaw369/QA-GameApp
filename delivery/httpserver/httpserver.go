package httpserver

import (
	"Q/A-GameApp/config"
	"Q/A-GameApp/service/authservice"
	"Q/A-GameApp/service/userservise"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	config  config.Config
	authSvc authservice.Service
	userSvc userservise.Service
}

func New(config config.Config, authSvc authservice.Service, userSvc userservise.Service) Server {
	return Server{
		config:  config,
		authSvc: authSvc,
		userSvc: userSvc,
	}
}
func (s Server) Server() {
	e := echo.New()
	//middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	//routes
	e.GET("/health-check", s.healthCheck)

	userGroups := e.Group("/users")
	userGroups.POST("/register", s.userRegister)
	userGroups.POST("/login", s.userLogin)
	userGroups.GET("/profile", s.userProfile)
	// start server
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", s.config.HTTPServer.Port)))
}
