package todolist

import "go-todolist/models"

type ServiceInterface interface {
	CreateUser(username, firstname, lastname, password string) (*models.User, error)
	CreateTodoItem(user *models.User, title, description string) (*models.TodoItem, error)
	GetAllUsers() ([]*models.User, error)
	UserExists(username string) bool
	FindUserByUsername(username string) (*models.User, error)
	GetTodoItemsByUser(user models.User) ([]*models.TodoItem, error)
	GetAllTodoItems() ([]*models.TodoItem, error)
	DeleteUser(user *models.User) error
	DeleteTodoItem(todoItem *models.TodoItem) error
	UpdateUser(user *models.User) (*models.User, error)
	UpdateTodoItem(todoItem *models.TodoItem) (*models.TodoItem, error)
}
