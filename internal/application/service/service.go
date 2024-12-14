// Package service 包含应用程序的服务层。
// 这是应用服务层，负责处理外部接口（如 HTTP 请求）与用例层之间的交互。
// 它将外部请求转换为用例层的调用，并将用例层的结果返回给外部接口。
// 上一层：接口层
// 下一层：用例层

// 在 DDD 中，应用服务层负责处理外部接口与用例层之间的交互。
// 它将外部请求转换为用例层的调用，并将用例层的结果返回给外部接口。
// 该层的核心定义是应用服务。
// ***这是防腐层的一部分，用于隔离外部接口的变化。***

// Diagram:
// 接口层 -> 应用服务层 -> 用例层

package service

import (
    "project-layout/internal/domain/model"
    "project-layout/internal/application/usecases"
)

type ApplicationService struct {
    useCase *usecases.UseCase
}

func NewApplicationService(useCase *usecases.UseCase) *ApplicationService {
    return &ApplicationService{useCase: useCase}
}

func (s *ApplicationService) ExecuteFindByID(id string) (*model.Entity, error) {
    return s.useCase.FindByID(id)
}

func (s *ApplicationService) ExecuteCreate(entity model.Entity) error {
    return s.useCase.Create(entity)
}

func (s *ApplicationService) ExecuteUpdate(entity model.Entity) error {
    return s.useCase.Update(entity)
}

func (s *ApplicationService) ExecuteDelete(id string) error {
    return s.useCase.Delete(id)
}
