package cmd

import (
	"go-todolist/config"
	"go-todolist/database"

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

	return cnf, db, nil
}
