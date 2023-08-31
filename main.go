package main

import (
	"Q/A-GameApp/entity"
	"Q/A-GameApp/repository/mysql"
	"Q/A-GameApp/service/userservise"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/users/register", userRegisterHandler)
	log.Println("server is listening on port :8080")
	http.ListenAndServe(":8080", nil)

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
	userSvc := userservise.New(mysqlRepo)
	_, err = userSvc.Register(uReq)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err.Error())))

		return
	}
	writer.Write([]byte(`{"message":"user created"}`))
}
func testUserMysqlRepo() {
	mysqlrepo := mysql.New()
	createdUser, err := mysqlrepo.Register(entity.User{ID: 0, PhoneNumber: "09125843", Name: "John"})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(createdUser)
	}
	isUniqueUser, err := mysqlrepo.IsPhoneNumberUnique(createdUser.PhoneNumber)
	if err != nil {
		fmt.Println("unique err:", err)
	}
	fmt.Println("is unique:", isUniqueUser)

}
