package service

import (
	"project-layout/internal/domain/model"
	"project-layout/internal/domain/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user model.User) model.User {
	return s.repo.CreateUser(user)
}

func (s *UserService) GetUser(id string) (model.User, error) {
	return s.repo.GetUser(id)
}

func (s *UserService) UpdateUser(id string, updatedUser model.User) (model.User, error) {
	return s.repo.UpdateUser(id, updatedUser)
}

func (s *UserService) DeleteUser(id string) error {
	return s.repo.DeleteUser(id)
}
