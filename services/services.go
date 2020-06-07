package services

import (
	"go-todolist/todolist"
)

func init() {

}

var (

	TodolistService Todolist.ServiceInterface
)

// UseTodolistService sets the Todolist service
func UseTodolistService(o todolist.ServiceInterface) {
	TodolistService = o
}


// Init starts up all services
func Init(cnf *config.Config, db *gorm.DB) error {

	if nil == reflect.TypeOf(TodolistService) {
		TodolistService = Todolist.NewService(cnf, db)
	}

	return nil
}

// Close closes any open services
func Close() {

}
