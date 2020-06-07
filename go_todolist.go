package main

import (
	"fmt"
	"go-todolist/cmd"
	"go-todolist/models"
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

	user, err := service.CreateUser(fmt.Sprintf("johndoe%d", time.Now().UnixNano()), "John", "Atayi", "123456")

	if err != nil {
		panic(err)
	}
	fmt.Print(user)
	var todoitem *models.TodoItem

	todoitem, err = service.CreateTodoItem(user, "Call to Doctor", "Call to doctor Sternberg and tell her about Murads situation")
	if err != nil {
		panic(err)
	}
	fmt.Print(todoitem)

	todoitem, err = service.CreateTodoItem(user, "Call to Commerzbank", "Call to doctor Commerzbank and complain about online banking")
	if err != nil {
		panic(todoitem)
	}
	fmt.Print(todoitem)

}
