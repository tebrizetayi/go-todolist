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

func (suite *TodoListTestSuite) TestDeleteTodoItem() {
	var (
		user     *models.User
		todoItem *models.TodoItem
		err      error
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

	err = suite.service.DeleteTodoItem(todoItem)

	// Error should be nil
	assert.Nil(suite.T(), err)

	todoItem, err = suite.service.FindById(todoItem.ID)

	// Error should not be nil
	assert.NotNil(suite.T(), err)

	// Error should not be nil and todoItem should be nil
	if assert.NotNil(suite.T(), err) {
		assert.Nil(suite.T(), todoItem)
	}
}

func (suite *TodoListTestSuite) TestUpdateTodoItem() {
	var (
		user     *models.User
		todoItem *models.TodoItem
		err      error
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

	newTitle := "newTitle"
	todoItem.Title = newTitle
	suite.service.UpdateTodoItem(todoItem)

	todoItem, err = suite.service.FindById(todoItem.ID)

	// Error should be nil and todoItem's title should equal to  newTitle
	if assert.Nil(suite.T(), err) {
		assert.Equal(suite.T(), newTitle, todoItem.Title)
	}
}
