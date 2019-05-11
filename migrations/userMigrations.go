package main

import (
	"github.com/Maximo-Miranda/go-alvaro-userservice/utils"
	"github.com/Maximo-Miranda/go-alvaro-userservice/internal/user"
)

// main
func main(){

	user := user.UserModel{}

	conn, _ := utils.Connect()

	conn.LogMode(true)

	conn.AutoMigrate(&user)

	conn.Close()

}