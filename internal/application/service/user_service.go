// Package service 包含应用程序的服务层。
// 这是应用服务层，负责处理外部接口（如 HTTP 请求）与用例层之间的交互。
// 它将外部请求转换为用例层的调用，并将用例层的结果返回给外部接口。
// 上一层：接口层
// 下一层：用例层

package service

import (
	"project-layout/internal/application/usecases"
	"project-layout/internal/domain/model"
)

type UserService struct {
	useCase *usecases.UserUseCase
}

func NewUserService(useCase *usecases.UserUseCase) *UserService {
	return &UserService{useCase: useCase}
}

func (s *UserService) ExecuteFindByID(id string) (*model.User, error) {
	return s.useCase.FindByID(id)
}

func (s *UserService) ExecuteCreate(user model.User) error {
	return s.useCase.Create(user)
}

func (s *UserService) ExecuteUpdate(user model.User) error {
	return s.useCase.Update(user)
}

func (s *UserService) ExecuteDelete(id string) error {
	return s.useCase.Delete(id)
}
