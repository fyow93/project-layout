// Package service 包含领域服务层。
// 这是领域服务层，负责处理领域逻辑。
// 它与领域模型和仓储接口交互。
// 上一层：用例层
// 下一层：仓储层

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

func (s *UserService) Save(user model.User) error {
	return s.repo.Save(user)
}

func (s *UserService) FindByID(id string) (*model.User, error) {
	return s.repo.FindByID(id)
}

func (s *UserService) Update(user model.User) error {
	return s.repo.Update(user)
}

func (s *UserService) Delete(id string) error {
	return s.repo.Delete(id)
}
