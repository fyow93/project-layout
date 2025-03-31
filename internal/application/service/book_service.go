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
	"project-layout/internal/application/usecases"
	"project-layout/internal/domain/model"
)

// BookService 应用服务，协调用例执行并处理与外部交互
// 它作为领域层和接口层之间的中间层
type BookService struct {
	useCase *usecases.BookUseCase // 用例层引用
}

// NewBookService 创建新的书籍应用服务实例
// 参数:
//   - useCase: 书籍用例实例
// 返回:
//   - *BookService: 新创建的书籍应用服务
func NewBookService(useCase *usecases.BookUseCase) *BookService {
	return &BookService{useCase: useCase}
}

// ExecuteFindByID 执行根据ID查找书籍的用例
// 该方法作为接口层和用例层之间的桥梁
// 参数:
//   - id: 书籍的唯一标识符
// 返回:
//   - *model.Book: 找到的书籍模型指针
//   - error: 如果查找过程中发生错误
func (s *BookService) ExecuteFindByID(id string) (*model.Book, error) {
	return s.useCase.FindByID(id)
}

// ExecuteCreate 执行创建新书籍的用例
// 参数:
//   - book: 要创建的书籍模型
// 返回:
//   - error: 如果创建过程中发生错误
func (s *BookService) ExecuteCreate(book model.Book) error {
	return s.useCase.Create(book)
}

// ExecuteUpdate 执行更新书籍的用例
// 参数:
//   - book: 包含更新信息的书籍模型
// 返回:
//   - error: 如果更新过程中发生错误
func (s *BookService) ExecuteUpdate(book model.Book) error {
	return s.useCase.Update(book)
}

// ExecuteDelete 执行删除书籍的用例
// 参数:
//   - id: 要删除的书籍的唯一标识符
// 返回:
//   - error: 如果删除过程中发生错误
func (s *BookService) ExecuteDelete(id string) error {
	return s.useCase.Delete(id)
}
