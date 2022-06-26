package main

import (
	"fmt"

	"./app/models"
)

func main() {
	fmt.Println(models.Db)
	// defer Db.Close()

	/*
	// CREATE
	u := &models.User{}
	u.Name = "test"
	u.Email = "test@example.com"
	u.Password = "password"
	fmt.Println(u)

	u.CreateUser()
	*/

	/*
	// READ
	u, _ := models.GetUser(1)
	fmt.Println(u)
	*/

	/*
	// UPDATE
	u, _ := models.GetUser(1)
	u.Name = "Test2"
	u.Email = "test2@example.com"
	u.UpdateUser()
	// ↓確認用の出力
	u, _ = models.GetUser(1)
	fmt.Println(u)
	*/

	// DELETE
	u, _ := models.GetUser(1)
	u.DeleteUser()
	fmt.Println(u)
}