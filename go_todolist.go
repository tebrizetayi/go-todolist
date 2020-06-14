package main

import (
	"fmt"
	"go-todolist/cmd"
	"go-todolist/models"
	"go-todolist/todolist"
)

func main() {
	config, db, err := cmd.InitConfigDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	service := todolist.NewService(config, db)

	users, err := service.GetAllUsers()

	for i := 0; i < len(users); i++ {
		fmt.Println(users[i])
	}

	user := models.User{MyGormModel: models.MyGormModel{ID: 38}}

	todoItems, err := service.GetTodoItemsByUser(user)
	for i := 0; i < len(todoItems); i++ {
		fmt.Println(todoItems[i])
	}
}
