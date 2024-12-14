// Package usecases 包含应用程序的用例层。
// 这是用例层，负责定义应用程序的业务逻辑。
// 它协调领域服务来完成用例。
// 上一层：应用服务层
// 下一层：领域服务层

// 在 DDD 中，用例层负责定义应用程序的业务逻辑。
// 它协调领域服务来完成具体的业务操作。
// 该层的核心定义是用例。

// Diagram:
// 应用服务层 -> 用例层 -> 领域服务层

package usecases

import (
	"project-layout/internal/domain/model"
	"project-layout/internal/domain/service"
)

type UseCase struct {
	domainService *service.DomainService
}

func NewUseCase(domainService *service.DomainService) *UseCase {
	return &UseCase{domainService: domainService}
}

func (u *UseCase) FindByID(id string) (*model.Entity, error) {
	return u.domainService.FindByID(id)
}

func (u *UseCase) Create(entity model.Entity) error {
	return u.domainService.Save(entity)
}

func (u *UseCase) Update(entity model.Entity) error {
	return u.domainService.Update(entity)
}

func (u *UseCase) Delete(id string) error {
	return u.domainService.Delete(id)
}
