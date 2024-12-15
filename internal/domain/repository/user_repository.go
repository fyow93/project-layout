// Package repository 包含领域仓储接口。
// 这是仓储层，负责数据持久化。
// 它定义了数据存储和检索的接口。
// 上一层：领域服务层
// 下一层：具体的仓储实现

package repository

import "project-layout/internal/domain/model"

type UserRepository interface {
	Save(user model.User) error
	FindByID(id string) (*model.User, error)
	Update(user model.User) error
	Delete(id string) error
	Initialize() error
}
