package todolist

import "go-todolist/models"

func (s *Service) CreateTodoItem(user *models.User, title, description string) (*models.TodoItem, error) {
	todoitem := models.NewToDoItem(user, title, description)

	if err := s.db.Create(todoitem).Error; err != nil {
		return nil, err
	}
	return todoitem, nil
}
