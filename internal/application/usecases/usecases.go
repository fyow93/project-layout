// Package usecases 包含应用程序的用例层。
// 这是用例层，负责定义应用程序的业务逻辑。
// 它协调领域服务来完成用例。
// 上一层：应用服务层
// 下一层：领域服务层

// 在 DDD 中，用例层负责定义应用程序的业务逻辑。
// 它协调领域服务来完成具体的业务操作。
// 该层的核心定义是用例。
// **用例层（Use Case Layer）**：
//   - **用例（Use Case）**：用例层负责定义应用程序的业务逻辑。它协调领域服务来完成具体的业务操作。用例层的逻辑通常是跨领域对象的操作，或者是多个领域服务的协调。
//   - **业务逻辑**：用例层的业务逻辑通常是高层次的业务流程，例如验证输入数据、调用多个领域服务、处理事务等。



// Diagram:
// 应用服务层 -> 用例层 -> 领域服务层

package usecases

import (
	"errors"
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
	// 业务逻辑：检查ID格式
	if !isValidID(id) {
		return nil, errors.New("无效的ID格式")
	}
	entity, err := u.domainService.FindByID(id)
	if err != nil {
		return nil, err
	}
	// 业务逻辑：检查实体是否活跃
	if !entity.IsActive {
		return nil, errors.New("实体未激活")
	}
	return entity, nil
}

func (u *UseCase) Create(entity model.Entity) error {
	// 业务逻辑：验证实体
	if err := validateEntity(entity); err != nil {
		return err
	}
	// 业务逻辑：检查实体是否已存在
	existingEntity, _ := u.domainService.FindByID(entity.ID)
	if existingEntity != nil {
		return errors.New("实体已存在")
	}
	return u.domainService.Save(entity)
}

func (u *UseCase) Update(entity model.Entity) error {
	// 业务逻辑：检查实体是否存在
	existingEntity, err := u.domainService.FindByID(entity.ID)
	if err != nil {
		return err
	}
	if existingEntity == nil {
		return errors.New("实体不存在")
	}
	// 业务逻辑：验证更新数据
	if err := validateEntity(entity); err != nil {
		return err
	}
	return u.domainService.Update(entity)
}

func (u *UseCase) Delete(id string) error {
	// 业务逻辑：检查实体是否存在
	entity, err := u.domainService.FindByID(id)
	if err != nil {
		return err
	}
	if entity == nil {
		return errors.New("实体不存在")
	}
	// 业务逻辑：检查实体是否可以删除
	if !entity.CanBeDeleted {
		return errors.New("实体不能删除")
	}
	return u.domainService.Delete(id)
}

// 辅助函数：验证实体
func validateEntity(entity model.Entity) error {
	if entity.Name == "" {
		return errors.New("实体名称不能为空")
	}
	// 添加更多验证规则
	return nil
}

// 辅助函数：检查ID格式
func isValidID(id string) bool {
	// 示例：检查ID是否为UUID格式
	return len(id) == 36
}
