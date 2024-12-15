// Package usecases 包含应用程序的用例层。
// 这是用例层，负责定义应用程序的业务逻辑。
// 它协调领域服务来完成用例。
// 上一层：应用服务层
// 下一层：领域服务层

package usecases

import (
	"project-layout/internal/domain/model"
	"project-layout/internal/domain/service"
)

type UserUseCase struct {
	domainService *service.UserService
}

func NewUserUseCase(domainService *service.UserService) *UserUseCase {
	return &UserUseCase{domainService: domainService}
}

func (u *UserUseCase) FindByID(id string) (*model.User, error) {
	return u.domainService.FindByID(id)
}

func (u *UserUseCase) Create(user model.User) error {
	return u.domainService.Save(user)
}

func (u *UserUseCase) Update(user model.User) error {
	return u.domainService.Update(user)
}

func (u *UserUseCase) Delete(id string) error {
	return u.domainService.Delete(id)
}

type UserUseCaseFactory struct {
	domainService *service.UserService
}

func NewUserUseCaseFactory(domainService *service.UserService) *UserUseCaseFactory {
	return &UserUseCaseFactory{domainService: domainService}
}

func (f *UserUseCaseFactory) CreateUserUseCase() *UserUseCase {
	return NewUserUseCase(f.domainService)
}
