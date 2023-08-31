package main

import (
	"Q/A-GameApp/entity"
	"Q/A-GameApp/repository/mysql"
	"fmt"
)

func main() {

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
