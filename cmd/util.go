package cmd

import (
	"go-todolist/config"
	"go-todolist/database"
	"go-todolist/models"

	"github.com/jinzhu/gorm"
)

// initConfigDB loads the configuration and connects to the database
func InitConfigDB() (*config.Config, *gorm.DB, error) {
	// Config
	cnf := config.NewConfig()

	// Database
	db, err := database.NewDatabase(cnf)
	if err != nil {
		return nil, nil, err
	}

	//Creating Tables
	if !db.HasTable(&models.User{}) {
		db.CreateTable(&models.User{})
	}
	if !db.HasTable(&models.TodoItem{}) {
		db.CreateTable(&models.TodoItem{})
	}

	return cnf, db, nil
}
