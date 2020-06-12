package database

import (
	"go-todolist/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDatabase(t *testing.T) {
	cnf := config.NewTestConfig()
	// Database
	_, err := NewDatabase(cnf)
	assert.Nil(t, err)
}
