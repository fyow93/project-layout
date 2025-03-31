// Package service 包含应用程序的服务层。
// 这是应用服务层，负责处理外部接口（如 HTTP 请求）与用例层之间的交互。
// 它将外部请求转换为用例层的调用，并将用例层的结果返回给外部接口。
// 上一层：接口层
// 下一层：用例层

package service

import (
	"project-layout/internal/application/assembler"
	"project-layout/internal/application/dto"
	"project-layout/internal/application/usecases"
)

type UserService struct {
	useCase    *usecases.UserUseCase
	assembler  *assembler.UserAssembler
}

func NewUserService(useCase *usecases.UserUseCase) *UserService {
	return &UserService{
		useCase:    useCase,
		assembler:  &assembler.UserAssembler{},
	}
}

func (s *UserService) ExecuteFindByID(id string) (*dto.UserDTO, error) {
	user, err := s.useCase.FindByID(id)
	if err != nil {
		return nil, err
	}
	return s.assembler.ToDTO(user), nil
}

func (s *UserService) ExecuteCreate(createDTO dto.CreateUserDTO) error {
	user := s.assembler.ToModel(&createDTO)
	return s.useCase.Create(*user)
}

func (s *UserService) ExecuteUpdate(id string, updateDTO dto.UpdateUserDTO) error {
	user, err := s.useCase.FindByID(id)
	if err != nil {
		return err
	}
	s.assembler.UpdateModel(user, &updateDTO)
	return s.useCase.Update(*user)
}

func (s *UserService) ExecuteDelete(id string) error {
	return s.useCase.Delete(id)
}
