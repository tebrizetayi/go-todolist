package main

import (
	"fmt"
	"go-todolist/cmd"
	"go-todolist/todolist"
	"time"
)

func main() {
	config, db, err := cmd.InitConfigDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	service := todolist.NewService(config, db)
	user, errUser := service.CreateUser(fmt.Sprintf("johndoe%d", time.Now().UnixNano()), "John", "Atayi", "123456")

	if errUser != nil {
		panic(err)
	}
	fmt.Print(user)

}
