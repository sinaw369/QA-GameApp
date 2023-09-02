package main

import (
	"Q/A-GameApp/repository/mysql"
	"Q/A-GameApp/service/userservise"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const (
	JwtSignKey = "jwt_secret"
)

func main() {
	http.HandleFunc("/users/register", userRegisterHandler)
	http.HandleFunc("/users/health-check", healthCheckHandler)
	http.HandleFunc("/users/login", userLoginHandler)
	http.HandleFunc("/users/profile", userProfileHandler)
	log.Println("server is listening on port :8080")
	http.ListenAndServe(":8080", nil)

}
func healthCheckHandler(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(writer, `{"message":"ok"}`)
}

type UserRegisterHandler struct {
}

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
	mysqlRepo := mysql.New()
	userSvc := userservise.New(mysqlRepo, JwtSignKey)
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
	mysqlRepo := mysql.New()
	userSvc := userservise.New(mysqlRepo, JwtSignKey)
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

	// validate jwt token and

	preq := userservise.ProfileRequest{UserID: 0}
	data, err := io.ReadAll(req.Body)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err.Error())))

		return
	}

	err = json.Unmarshal(data, &preq)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err.Error())))

		return
	}

	mysqlRepo := mysql.New()
	userSvc := userservise.New(mysqlRepo, JwtSignKey)
	resp, err := userSvc.Profile(preq)
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
