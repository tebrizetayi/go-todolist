package services

import (
	"go-todolist/config"
	"go-todolist/todolist"
	"reflect"

	"github.com/jinzhu/gorm"
)

func init() {

}

var (
	TodolistService todolist.ServiceInterface
)

// UseTodolistService sets the Todolist service
func UseTodolistService(o todolist.ServiceInterface) {
	TodolistService = o
}

// Init starts up all services
func Init(cnf *config.Config, db *gorm.DB) error {

	if nil == reflect.TypeOf(TodolistService) {
		TodolistService = todolist.NewService(cnf, db)
	}

	return nil
}

// Close closes any open services
func Close() {}
