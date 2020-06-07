package todolist

import "go-todolist/models"

func (s *Service) CreateUser(username, firstname, lastname, password string) (*models.User, error) {
	user := models.NewUser(username, firstname, lastname, password)

	if err := s.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
