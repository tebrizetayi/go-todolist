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
		"Johndoe1", // username
		"John",     // firstname
		"Doe",      // lastname
		"123456",   // password
	)

	// Error should be nil
	assert.Nil(suite.T(), err)

	// Correct user object should be returned
	if assert.NotNil(suite.T(), user) {
		assert.Equal(suite.T(), "Johndoe1", user.Username)
	}
}

func (suite *TodoListTestSuite) TestGetAllUsers() {
	var (
		userCount int
		users     []*models.User
	)
	for i := 0; i < userCount; i++ {
		_, _ = suite.service.CreateUser("johndoe"+string(i), "john", "doe", "123456")
	}

	users, _ = suite.service.GetAllUsers()
	assert.Equal(suite.T(), userCount, len(users))
}

func (suite *TodoListTestSuite) TestDeleteUser() {
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

	err = suite.service.DeleteUser(user)

	// Error should be nil
	assert.Nil(suite.T(), err)

	user, err = suite.service.FindUserByUsername(user.Username)

	// Error should not be nil
	assert.NotNil(suite.T(), err)

	// Error should not be nil and user should be nil
	if assert.NotNil(suite.T(), err) {
		assert.Nil(suite.T(), user)
	}
}

func (suite *TodoListTestSuite) TestUpdateUser() {
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

	newLastName := "UpdatedDoe"
	user.Lastname = newLastName
	suite.service.UpdateUser(user)

	user, err = suite.service.FindUserByUsername(user.Username)

	// Error should not be nil and user should be nil
	if assert.Nil(suite.T(), err) {
		assert.Equal(suite.T(), newLastName, user.Lastname)
	}
}
