package main

import (
	"Q/A-GameApp/config"
	"Q/A-GameApp/delivery/httpserver"
	"Q/A-GameApp/repository/mysql"
	"Q/A-GameApp/service/authservice"
	"Q/A-GameApp/service/userservise"
	"Q/A-GameApp/validator/uservalidator"
	"time"
)

const (
	JwtSignKey                 = "jwt_secret"
	AccessTokenSubject         = "at"
	RefreshTokenSubject        = "rt"
	AccessTokenExpireDuration  = time.Hour * 24
	RefreshTokenExpireDuration = time.Hour * 24 * 7
)

func setupServices(cfg config.Config) (authservice.Service, userservise.Service, uservalidator.Validator) {
	authSvc := authservice.New(cfg.Auth)
	mysqlRepo := mysql.New(cfg.Mysql)
	validateUserSvc := uservalidator.New(mysqlRepo)
	userSvc := userservise.New(authSvc, mysqlRepo)
	return authSvc, userSvc, validateUserSvc
}
func main() {
	cfg := config.Config{
		HTTPServer: config.HTTPServer{Port: 8080},
		Auth: authservice.Config{
			SignKey:               JwtSignKey,
			AccessExpirationTime:  AccessTokenExpireDuration,
			RefreshExpirationTime: RefreshTokenExpireDuration,
			RefreshSubject:        RefreshTokenSubject,
			AccessSubject:         AccessTokenSubject,
		},
		Mysql: mysql.Config{
			UserName: "gameapp",
			Password: "some_example",
			Host:     "localhost",
			Port:     3306,
			Scheme:   "gameapp_db",
		},
	}

	authSvc, userSvc, validateUsersvc := setupServices(cfg)

	server := httpserver.New(cfg, authSvc, userSvc, validateUsersvc)
	server.Server()
}
