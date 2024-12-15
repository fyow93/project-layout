// Package repository 包含领域仓储接口。
// 这是仓储层，负责数据持久化。
// 它定义了数据存储和检索的接口。
// 上一层：领域服务层
// 下一层：具体的仓储实现

// 在 DDD 中，仓储接口定义了领域对象的存储和检索方法。
// 它是领域层与数据存储层之间的桥梁。
// 该层的核心定义是仓储接口。
// ***这是防腐层的一部分，用于隔离领域层与外部数据存储的变化。***

// Diagram:
// 领域服务层 -> 仓储接口 -> 具体的仓储实现

package repository

import "project-layout/internal/domain/model"

type BookRepository interface {
	Save(book model.Book) error
	FindByID(id string) (*model.Book, error)
	Update(book model.Book) error
	Delete(id string) error
}
