package todolist

import "go-todolist/models"

type ServiceInterface interface {
	CreateUser(username, firstname, lastname, password string) (*models.User, error)
	CreateTodoItem(user *models.User, title, description string) (*models.TodoItem, error)
	GetAllUsers() ([]*model.User, error)
}
