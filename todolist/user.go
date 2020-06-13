package todolist

import (
	"go-todolist/models"
	pass "go-todolist/util/password"
)

func (s *Service) CreateUser(username, firstname, lastname, password string) (*models.User, error) {
	user := models.NewUser(username, firstname, lastname, password)
	passwordHashed, err := pass.HashPassword(password)
	if err != nil {
		return nil, err
	}
	user.Password = string(passwordHashed)

	if err := s.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Service) GetAllUsers() ([]*models.User, error) {
	var users []*models.User
	if err := s.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
