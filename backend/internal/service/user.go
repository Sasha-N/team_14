package service

import (
	"errors"
	"personal-finance-app/internal/models"
	"personal-finance-app/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	RegisterUser(name, password string) error
	CheckPassword(name, password string) (uint, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) RegisterUser(name, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &models.User{Name: name, Password: string(hashedPassword)}
	err = s.repo.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (s *userService) CheckPassword(name, password string) (uint, error) {
	user, err := s.repo.GetUserByName(name)
	if err != nil {
		return 0, err
	}
	if user == nil {
		return 0, errors.New("invalid username or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return 0, errors.New("invalid username or password")
	}
	return user.ID, nil
}
