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

// BookRepository 定义了书籍领域模型的持久化接口
// 该接口遵循仓储模式，为领域层提供数据访问抽象
// 实现此接口的具体类负责处理数据存储细节
type BookRepository interface {
	// Save 保存新书籍到数据存储
	// 参数:
	//   - book: 要保存的书籍领域模型
	// 返回:
	//   - error: 如果保存过程中发生错误，返回相应错误；否则返回 nil
	Save(book model.Book) error

	// FindByID 根据ID从数据存储中查找书籍
	// 参数:
	//   - id: 书籍的唯一标识符
	// 返回:
	//   - *model.Book: 找到的书籍模型指针，如果未找到则为 nil
	//   - error: 如果查找过程中发生错误，返回相应错误；否则返回 nil
	FindByID(id string) (*model.Book, error)

	// Update 更新数据存储中现有书籍的信息
	// 参数:
	//   - book: 包含更新信息的书籍领域模型
	// 返回:
	//   - error: 如果更新过程中发生错误，返回相应错误；否则返回 nil
	Update(book model.Book) error

	// Delete 从数据存储中删除指定ID的书籍
	// 参数:
	//   - id: 要删除的书籍的唯一标识符
	// 返回:
	//   - error: 如果删除过程中发生错误，返回相应错误；否则返回 nil
	Delete(id string) error
}
