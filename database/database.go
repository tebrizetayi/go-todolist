package database

import (
	"errors"
	"fmt"
	"go-todolist/config"
	"go-todolist/models"

	"github.com/jinzhu/gorm"

	// Drivers
	_ "github.com/lib/pq"
)

var (
	ErrDatabaseNotConnected = errors.New("There is no connection to database!")
)

func NewDatabase(cnf *config.Config) (*gorm.DB, error) {
	// Postgres
	if cnf.Database.Type == "postgres" {
		// Connection args
		// see https://godoc.org/github.com/lib/pq#hdr-Connection_String_Parameters
		args := fmt.Sprintf(
			"sslmode=disable host=%s port=%d user=%s password='%s' dbname=%s",
			cnf.Database.Host,
			cnf.Database.Port,
			cnf.Database.User,
			cnf.Database.Password,
			cnf.Database.DatabaseName,
		)

		db, err := gorm.Open(cnf.Database.Type, args)
		if err != nil {
			return db, err
		}

		// Max idle connections
		db.DB().SetMaxIdleConns(cnf.Database.MaxIdleConns)

		// Max open connections
		db.DB().SetMaxOpenConns(cnf.Database.MaxOpenConns)

		// Database logging
		db.LogMode(cnf.IsDevelopment)

		if err = CreateTables(db); err != nil {
			return nil, err
		}
		return db, nil
	}

	// Database type not supported
	return nil, fmt.Errorf("Database type %s not suppported", cnf.Database.Type)
}

func CreateTables(db *gorm.DB) error {

	if db == nil {
		return ErrDatabaseNotConnected
	}

	//Creating Tables
	if !db.HasTable(&models.User{}) {
		db.CreateTable(&models.User{})
	}
	if !db.HasTable(&models.TodoItem{}) {
		db.CreateTable(&models.TodoItem{})
	}

	return nil
}
