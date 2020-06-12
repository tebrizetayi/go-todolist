package todolist

import (
	"go-todolist/models"

	"github.com/stretchr/testify/assert"
)

func (suite *TodoListTestSuite) TestCreateUser() {
	var (
		user *models.User
		err  error
	)

	// We try to insert a unique user
	user, err = suite.service.CreateUser(
		"johndoe", // username
		"John",    // firstname
		"Doe",     // lastname
		"123456",  // password
	)

	// User object should not be nil
	assert.NotNil(suite.T(), user)

	// Correct user object should be returned
	if assert.NotNil(suite.T(), user) {
		assert.Equal(suite.T(), "johndoe", user.Username)
	}

	// We try to insert a non unique user
	user, err = suite.service.CreateUser(
		"johndoe", // username
		"John",    // firstname
		"Doe",     // lastname
		"123456",  // password
	)

	// Error should not be nil
	assert.NotNil(suite.T(), err)

	// Test username case insensitivity
	user, err = suite.service.CreateUser(
		"Johndoe", // username
		"John",    // firstname
		"Doe",     // lastname
		"123456",  // password
	)

	// Error should be nil
	assert.Nil(suite.T(), err)

	// Correct user object should be returned
	if assert.NotNil(suite.T(), user) {
		assert.Equal(suite.T(), "Johndoe", user.Username)
	}
}
