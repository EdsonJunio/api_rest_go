package repository

import "api_rest_go/internal/model"

type UserRepository interface {
	GetAll() ([]model.User, error)
	GetByID(id int) (model.User, error)
	Create(u model.User) (model.User, error)
	Update(u model.User) (model.User, error)
	Delete(id int) error
}
