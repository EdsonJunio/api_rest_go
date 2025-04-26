package service

import (
	"api_rest_go/internal/model"
	"api_rest_go/internal/repository"
)

type UserService interface {
	ListUsers() ([]model.User, error)
	GetUser(id int) (model.User, error)
	CreateUser(input CreateUserInput) (model.User, error)
	UpdateUser(id int, input UpdateUserInput) (model.User, error)
	DeleteUser(id int) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) ListUsers() ([]model.User, error) {
	return s.repo.GetAll()
}

func (s *userService) GetUser(id int) (model.User, error) {
	return s.repo.GetByID(id)
}

func (s *userService) CreateUser(input CreateUserInput) (model.User, error) {
	user := model.User{
		Name:  input.Name,
		Email: input.Email,
	}
	return s.repo.Create(user)
}

func (s *userService) UpdateUser(id int, input UpdateUserInput) (model.User, error) {
	if _, err := s.repo.GetByID(id); err != nil {
		return model.User{}, err
	}
	user := model.User{ID: id, Name: input.Name, Email: input.Email}
	return s.repo.Update(user)
}

func (s *userService) DeleteUser(id int) error {
	return s.repo.Delete(id)
}
