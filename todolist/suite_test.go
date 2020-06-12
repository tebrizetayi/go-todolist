package todolist

import (
	"go-todolist/config"
	"go-todolist/database"
	"go-todolist/models"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/suite"
)

type TodoListTestSuite struct {
	suite.Suite
	cnf     *config.Config
	db      *gorm.DB
	service *Service
}

// The SetupSuite method will be run by testify once, at the very
// start of the testing suite, before any tests are run.
func (suite *TodoListTestSuite) SetupSuite() {
	// Initialise the config
	suite.cnf = config.NewTestConfig()

	// Create the test database
	db, err := database.NewDatabase(suite.cnf)
	if err != nil {
		panic(err)
	}
	suite.db = db

	// Initialise the service
	suite.service = NewService(suite.cnf, suite.db)
}

// The TearDownSuite method will be run by testify once, at the very
// end of the testing suite, after all tests have been run.
func (suite *TodoListTestSuite) TearDownSuite() {
	suite.db.DropTable(new(models.User))
	suite.db.DropTable(new(models.TodoItem))
}

// The SetupTest method will be run before every test in the suite.
func (suite *TodoListTestSuite) SetupTest() {
	//
}

// The TearDownTest method will be run after every test in the suite.
func (suite *TodoListTestSuite) TearDownTest() {
	suite.db.Unscoped().Delete(new(models.User))
	suite.db.Unscoped().Delete(new(models.TodoItem))
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestTodoListTestSuite(t *testing.T) {
	suite.Run(t, new(TodoListTestSuite))
}
