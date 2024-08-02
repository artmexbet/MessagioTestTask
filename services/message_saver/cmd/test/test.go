package main

import (
	"MessagioTestTask/pkg/db"
	"MessagioTestTask/pkg/models"
	"fmt"
)

func main() {
	db, err := db.New(&db.Config{
		Host:     "localhost",
		Port:     "5434",
		User:     "baseuser",
		Password: "basepassword",
		DbName:   "testtask",
	})
	fmt.Println(err)
	i, err := db.AddMessage(models.Message{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i)
}
