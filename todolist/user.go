package todolist

import (
	"errors"
	"go-todolist/models"
	pass "go-todolist/util/password"
	"strings"
)

func (s *Service) CreateUser(username, firstname, lastname, password string) (*models.User, error) {
	user := models.NewUser(username, firstname, lastname, password)
	passwordHashed, err := pass.HashPassword(password)
	if err != nil {
		return nil, err
	}
	user.Password = string(passwordHashed)

	if s.UserExists(username) {
		return nil, errors.New("Username taken")
	}
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

func (s *Service) FindUserByUsername(username string) (*models.User, error) {
	user := new(models.User)
	notFound := s.db.Where("username = LOWER(?)", strings.ToLower(username)).First(user).RecordNotFound()

	// Not found
	if notFound {
		return nil, errors.New("User not found")
	}

	return user, nil
}

func (s *Service) UserExists(username string) bool {
	_, err := s.FindUserByUsername(username)
	return err == nil
}

func (s *Service) UpdateUser(user *models.User) (*models.User, error) {

	if err := s.db.Save(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) DeleteUser(user *models.User) error {
	if err := s.db.Delete(&user).Error; err != nil {
		return err
	}
	return nil
}
