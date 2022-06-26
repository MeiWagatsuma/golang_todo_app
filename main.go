package main

import (
	"fmt"

	"./app/models"
)

func main() {
	fmt.Println(models.Db)
	// defer Db.Close()

	/*
	u := &models.User{}
	u.Name = "test"
	u.Email = "test@example.com"
	u.Password = "password"
	fmt.Println(u)

	u.CreateUser()
	*/

	u, _ := models.GetUser(1)

	fmt.Println(u)
}