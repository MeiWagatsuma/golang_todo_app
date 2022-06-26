package models

import (
	"log"
	"time"
)

type User struct {
	ID int
	UUID string
	Name string
	Email string
	Password string
	Create_at time.Time
}

func (u *User) CreateUser() (err error) {
	cmd := `INSERT INTO users (
		uuid,
		name, 
		email, 
		password, 
		created_at) VALUES (?, ?, ?, ?, ?)`

	_, err = Db.Exec(cmd, 
		createUUID(), 
		u.Name, 
		u.Email, 
		Encrypt(u.Password), 
		time.Now())

	if err != nil {
		log.Fatal(err)
	}

	return err
}