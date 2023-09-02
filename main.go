package main

import (
	"Q/A-GameApp/config"
	"Q/A-GameApp/delivery/httpserver"
	"Q/A-GameApp/repository/mysql"
	"Q/A-GameApp/service/authservice"
	"Q/A-GameApp/service/userservise"
	"time"
)

const (
	JwtSignKey                 = "jwt_secret"
	AccessTokenSubject         = "at"
	RefreshTokenSubject        = "rt"
	AccessTokenExpireDuration  = time.Hour * 24
	RefreshTokenExpireDuration = time.Hour * 24 * 7
)

func setupServices(cfg config.Config) (authservice.Service, userservise.Service) {
	authSvc := authservice.New(cfg.Auth)
	mysqlRepo := mysql.New(cfg.Mysql)
	userSvc := userservise.New(authSvc, mysqlRepo)
	return authSvc, userSvc
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

	authSvc, userSvc := setupServices(cfg)
	server := httpserver.New(cfg, authSvc, userSvc)
	server.Server()
}

/*
func userRegisterHandler(writer http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		fmt.Fprintf(writer, `{"error":"invalid request"}`)

		return
	}
	data, err := io.ReadAll(req.Body)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err.Error())))

		return
	}
	var uReq userservise.RegisterRequest
	err = json.Unmarshal(data, &uReq)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err.Error())))

		return
	}
	authSvc := authservice.New(JwtSignKey, AccessTokenSubject, RefreshTokenSubject,
		RefreshTokenExpireDuration, AccessTokenExpireDuration)
	mysqlRepo := mysql.New()
	userSvc := userservise.New(authSvc, mysqlRepo)
	_, err = userSvc.Register(uReq)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err.Error())))

		return
	}
	writer.Write([]byte(`{"message":"user created"}`))
}
func userLoginHandler(writer http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		fmt.Fprintf(writer, `{"error":"invalid request"}`)

		return
	}
	data, err := io.ReadAll(req.Body)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err.Error())))

		return
	}
	var lreq userservise.LoginRequest
	err = json.Unmarshal(data, &lreq)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err.Error())))

		return
	}
	authSvc := authservice.New(JwtSignKey, AccessTokenSubject, RefreshTokenSubject,
		RefreshTokenExpireDuration, AccessTokenExpireDuration)
	mysqlRepo := mysql.New()
	userSvc := userservise.New(authSvc, mysqlRepo)
	resp, err := userSvc.Login(lreq)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err.Error())))

		return
	}
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err.Error())))

		return
	}
	data, err = json.Marshal(resp)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err.Error())))

		return

	}
	writer.Write(data)
}
func userProfileHandler(writer http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		fmt.Fprintf(writer, `{"error":"invalid request"}`)

		return
	}
	authSvc := authservice.New(JwtSignKey, AccessTokenSubject, RefreshTokenSubject,
		RefreshTokenExpireDuration, AccessTokenExpireDuration)
	authToken := req.Header.Get("Authorization")
	//fmt.Println("auth:", auth)
	//fmt.Println(req.Header)
	claims, err := authSvc.VarifyToken(authToken)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err.Error())))
	}
	mysqlRepo := mysql.New()
	userSvc := userservise.New(authSvc, mysqlRepo)
	resp, err := userSvc.Profile(userservise.ProfileRequest{UserID: claims.UserID})
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err.Error())))

		return
	}
	//data, err := json.MarshalIndent(resp, "", "\t")
	data, err := json.Marshal(resp)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err.Error())))

		return

	}
	writer.Write(data)
}
*/
