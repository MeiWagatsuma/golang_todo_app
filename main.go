package main

import (
	"fmt"
	"log"

	"./app/models"
)

func main() {
	fmt.Println(models.Db)

	// controllers.StartMainServer()
	user, _ := models.GetUserByEmail("test@example.com")
	fmt.Println(user)

	session, err := user.CreateSession()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(session)

	valid, err := session.CheckSession()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(valid)
}
