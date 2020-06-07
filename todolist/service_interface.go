package todolist

import "go-todolist/models"

type ServiceInterface interface {
	CreateUser(username, firstname, lastname, password string) (*models.User, error)
}
