package repository

import (
	"project-layout/internal/domain/model"
)

type UserRepository interface {
	CreateUser(user model.User) model.User
	GetUser(id string) (model.User, error)
	UpdateUser(id string, updatedUser model.User) (model.User, error)
	DeleteUser(id string) error
}
