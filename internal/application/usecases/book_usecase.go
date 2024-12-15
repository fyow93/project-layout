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

	"github.com/go-playground/validator/v10"
)

type BookUseCase struct {
	bookService *service.BookService
}

func NewBookUseCase(domainService *service.BookService) *BookUseCase {
	return &BookUseCase{bookService: domainService}
}

func (u *BookUseCase) FindByID(id string) (*model.Book, error) {
	// 业务逻辑：检查ID格式
	if err := validateID(id); err != nil {
		return nil, err
	}
	book, err := u.bookService.FindByID(id)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (u *BookUseCase) Create(book model.Book) error {
	// 业务逻辑：验证实体
	if err := model.ValidateBook(&book); err != nil {
		return err
	}
	// 业务逻辑：检查实体是否已存在
	existingBook, _ := u.bookService.FindByID(book.ID)
	if existingBook != nil {
		return errors.New("实体已存在")
	}
	return u.bookService.Save(book)
}

func (u *BookUseCase) Update(book model.Book) error {
	// 业务逻辑：验证实体
	if err := model.ValidateBook(&book); err != nil {
		return err
	}
	// 业务逻辑：检查实体是否存在
	existingBook, err := u.bookService.FindByID(book.ID)
	if err != nil {
		return err
	}
	if existingBook == nil {
		return errors.New("实体不存在")
	}
	return u.bookService.Update(book)
}

func (u *BookUseCase) Delete(id string) error {
	// 业务逻辑：检查ID格式
	if err := validateID(id); err != nil {
		return err
	}
	// 业务逻辑：检查实体是否存在
	existingBook, err := u.bookService.FindByID(id)
	if err != nil {
		return err
	}
	if existingBook == nil {
		return errors.New("实体不存在")
	}
	return u.bookService.Delete(id)
}

func validateID(id string) error {
	validate := validator.New()
	return validate.Var(id, "required,uuid4")
}

type BookUseCaseFactory struct {
	domainService *service.BookService
}

func NewBookUseCaseFactory(domainService *service.BookService) *BookUseCaseFactory {
	return &BookUseCaseFactory{domainService: domainService}
}

func (f *BookUseCaseFactory) CreateBookUseCase() *BookUseCase {
	return NewBookUseCase(f.domainService)
}
