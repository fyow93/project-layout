// Package service 包含领域服务层。
// 这是领域服务层，负责处理领域逻辑。
// 它与领域模型和仓储接口交互。
// 上一层：用例层
// 下一层：仓储层

// 在 DDD 中，领域服务层负责处理复杂的领域逻辑。
// 它封装了领域模型的操作，并通过仓储接口与数据存储层交互。
// 该层的核心定义是领域服务。
//
// **领域层（Domain Layer）**：
//   - **领域服务（Domain Service）**：领域服务包含核心业务逻辑和规则。它们直接操作领域对象（实体、值对象等），并确保业务规则的一致性。
//   - **实体（Entity）**：实体是具有唯一标识的对象，包含业务逻辑和属性。
//   - **值对象（Value Object）**：值对象是不可变的对象，通常用于描述领域中的某些属性。

// Diagram:
// 用例层 -> 领域服务层 -> 仓储层

package service

import (
	"project-layout/internal/domain/model"
	"project-layout/internal/domain/repository"
)

// BookService 领域服务，处理书籍相关的核心业务逻辑
type BookService struct {
	repo repository.BookRepository
}

func NewBookService(repo repository.BookRepository) *BookService {
	return &BookService{repo: repo}
}

// Save 保存一本新书到数据库
// 参数:
//   - book: 要保存的书籍领域模型
//
// 返回:
//   - error: 如果保存过程中发生错误，返回相应错误；否则返回 nil
func (s *BookService) Save(book model.Book) error {
	return s.repo.Save(book)
}

// FindByID 根据ID查找书籍
// 参数:
//   - id: 书籍的唯一标识符
//
// 返回:
//   - *model.Book: 找到的书籍模型指针，如果未找到则为 nil
//   - error: 如果查找过程中发生错误，返回相应错误；否则返回 nil
func (s *BookService) FindByID(id string) (*model.Book, error) {
	return s.repo.FindByID(id)
}

// Update 更新现有书籍的信息
// 参数:
//   - book: 包含更新信息的书籍领域模型
//
// 返回:
//   - error: 如果更新过程中发生错误，返回相应错误；否则返回 nil
func (s *BookService) Update(book model.Book) error {
	return s.repo.Update(book)
}

// Delete 根据ID删除书籍
// 参数:
//   - id: 要删除的书籍的唯一标识符
//
// 返回:
//   - error: 如果删除过程中发生错误，返回相应错误；否则返回 nil
func (s *BookService) Delete(id string) error {
	return s.repo.Delete(id)
}
