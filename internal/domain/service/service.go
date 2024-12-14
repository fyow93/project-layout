// Package service 包含领域服务层。
// 这是领域服务层，负责处理领域逻辑。
// 它与领域模型和仓储接口交互。
// 上一层：用例层
// 下一层：仓储层

// 在 DDD 中，领域服务层负责处理复杂的领域逻辑。
// 它封装了领域模型的操作，并通过仓储接口与数据存储层交互。
// 该层的核心定义是领域服务。

// Diagram:
// 用例层 -> 领域服务层 -> 仓储层

package service

import (
	"project-layout/internal/domain/model"
	"project-layout/internal/domain/repository"
)

type DomainService struct {
	repo repository.Repository
}

func NewDomainService(repo repository.Repository) *DomainService {
	return &DomainService{repo: repo}
}

func (s *DomainService) Save(entity model.Entity) error {
	return s.repo.Save(entity)
}

func (s *DomainService) FindByID(id string) (*model.Entity, error) {
	return s.repo.FindByID(id)
}

func (s *DomainService) Update(entity model.Entity) error {
	return s.repo.Update(entity)
}

func (s *DomainService) Delete(id string) error {
	return s.repo.Delete(id)
}
