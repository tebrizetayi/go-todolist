package todolist

import (
	"go-todolist/models"

	"github.com/stretchr/testify/assert"
)

func (suite *TodoListTestSuite) TestCreateTodoItem() {
	var (
		user     *models.User
		todoItem *models.TodoItem
	)

	user, _ = suite.service.CreateUser(
		"johndoe", // username
		"John",    // firstname
		"Doe",     // lastname
		"123456",  // password
	)

	// We try to insert a todoitem

	todoItem, _ = suite.service.CreateTodoItem(
		user,          // user
		"Title",       // Title
		"Description", // Description
	)

	// User object should not be nil
	assert.NotNil(suite.T(), todoItem)
}

func (suite *TodoListTestSuite) TestGet() {
	var (
		user           *models.User
		todoItemsCount int = 4
		todoItems      []*models.TodoItem
	)

	user, _ = suite.service.CreateUser(
		"johndoe", // username
		"John",    // firstname
		"Doe",     // lastname
		"123456",  // password
	)

	// We try to insert a todoitem

	for i := 0; i < todoItemsCount; i++ {
		_, _ = suite.service.CreateTodoItem(
			user,          // user
			"Title",       // Title
			"Description", // Description
		)
	}

	todoItems, _ = suite.service.GetTodoItemsByUser(*user)
	// User object should not be nil
	assert.Equal(suite.T(), todoItemsCount, len(todoItems))
}
