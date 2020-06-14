package todolist

import (
	"errors"
	"go-todolist/models"
)

func (s *Service) CreateTodoItem(user *models.User, title, description string) (*models.TodoItem, error) {
	todoitem := models.NewToDoItem(user, title, description)

	if err := s.db.Create(todoitem).Error; err != nil {
		return nil, err
	}
	return todoitem, nil
}

func (s *Service) GetTodoItemsByUser(user models.User) ([]*models.TodoItem, error) {
	var todoItems []*models.TodoItem

	if err := s.db.Where(&models.TodoItem{UserID: user.ID}).Find(&todoItems).Error; err != nil {
		return nil, err
	}
	return todoItems, nil
}

func (s *Service) GetAllTodoItems() ([]*models.TodoItem, error) {
	var todoItems []*models.TodoItem

	if err := s.db.Find(&todoItems).Error; err != nil {
		return nil, err
	}
	return todoItems, nil
}

func (s *Service) UpdateTodoItem(todoItem *models.TodoItem) (*models.TodoItem, error) {

	if err := s.db.Save(&todoItem).Error; err != nil {
		return nil, err
	}
	return todoItem, nil
}

func (s *Service) DeleteTodoItem(todoItem *models.TodoItem) error {
	if err := s.db.Delete(&todoItem).Error; err != nil {
		return err
	}
	return nil
}

func (s *Service) FindById(id int) (*models.TodoItem, error) {
	todoitem := new(models.TodoItem)
	notFound := s.db.Where("ID=?", id).First(todoitem).RecordNotFound()

	if notFound {
		return nil, errors.New("TodoItem is not found!")
	}

	return todoitem, nil
}
