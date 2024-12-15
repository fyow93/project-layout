// Package service 包含应用程序的服务层。
// 这是应用服务层，负责处理外部接口（如 HTTP 请求）与用例层之间的交互。
// 它将外部请求转换为用例层的调用，并将用例层的结果返回给外部接口。
// 上一层：接口层
// 下一层：用例层

// 在 DDD 中，应用服务层负责处理外部接口与用例层之间的交互。
// 它将外部请求转换为用例层的调用，并将用例层的结果返回给外部接口。
// 该层的核心定义是应用服务。
// ***这是防腐层的一部分，用于隔离外部接口的变化。***
// **应用服务层（Application Service Layer）**：
//   - **应用服务（Application Service）**：应用服务层负责处理外部接口（如HTTP请求）与用例层之间的交互。它将外部请求转换为用例层的调用，并将用例层的结果返回给外部接口。
//   - **应用逻辑**：应用服务层的逻辑通常包括权限检查、日志记录、数据转换等。这些逻辑与具体的业务规则无关，但对于应用程序的运行是必要的。

// Diagram:
// 接口层 -> 应用服务层 -> 用例层

package service

import (
	"errors"
	"log"
	"project-layout/internal/application/usecases"
	"project-layout/internal/domain/model"
	"strings"
)

type ApplicationService struct {
	useCase *usecases.UseCase
}

func NewApplicationService(useCase *usecases.UseCase) *ApplicationService {
	return &ApplicationService{useCase: useCase}
}

func (s *ApplicationService) ExecuteFindByID(id string) (*model.Entity, error) {
	// 应用逻辑：记录日志
	log.Printf("执行查找操作，ID: %s", id)
	entity, err := s.useCase.FindByID(id)
	if err != nil {
		return nil, err
	}
	// 应用逻辑：转换实体以供响应
	return transformEntityForResponse(entity), nil
}

func (s *ApplicationService) ExecuteCreate(entity model.Entity) error {
	// 应用逻辑：记录日志
	log.Printf("执行创建操作，实体: %+v", entity)
	// 应用逻辑：检查权限
	if !hasCreatePermission(entity) {
		return errors.New("没有创建权限")
	}
	return s.useCase.Create(entity)
}

func (s *ApplicationService) ExecuteUpdate(entity model.Entity) error {
	// 应用逻辑：记录日志
	log.Printf("执行更新操作，实体: %+v", entity)
	// 应用逻辑：检查权限
	if !hasUpdatePermission(entity) {
		return errors.New("没有更新权限")
	}
	return s.useCase.Update(entity)
}

func (s *ApplicationService) ExecuteDelete(id string) error {
	// 应用逻辑：记录日志
	log.Printf("执行删除操作，ID: %s", id)
	// 应用逻辑：检查权限
	if !hasDeletePermission(id) {
		return errors.New("没有删除权限")
	}
	return s.useCase.Delete(id)
}

// 辅助函数：转换实体以供响应
func transformEntityForResponse(entity *model.Entity) *model.Entity {
	// 示例转换逻辑
	entity.Name = strings.ToUpper(entity.Name)
	return entity
}

// 辅助函数：检查创建权限
func hasCreatePermission(entity model.Entity) bool {
	// 示例权限检查逻辑
	return true
}

// 辅助函数：检查更新权限
func hasUpdatePermission(entity model.Entity) bool {
	// 示例权限检查逻辑
	return true
}

// 辅助函数：检查删除权限
func hasDeletePermission(id string) bool {
	// 示例权限检查逻辑
	return true
}
