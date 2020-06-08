package todolist

import "go-todolist/models"

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
