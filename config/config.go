package config

import (
	"Q/A-GameApp/repository/mysql"
	"Q/A-GameApp/service/authservice"
)

type HTTPServer struct {
	Port int
}
type Config struct {
	HTTPServer HTTPServer
	Auth       authservice.Config
	Mysql      mysql.Config
}
