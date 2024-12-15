package service

import (
	"project-layout/internal/application/usecases"
)

type ServiceFactory struct {
	bookUseCaseFactory *usecases.BookUseCaseFactory
	userUseCaseFactory *usecases.UserUseCaseFactory
}

func NewServiceFactory(bookUseCaseFactory *usecases.BookUseCaseFactory, userUseCaseFactory *usecases.UserUseCaseFactory) *ServiceFactory {
	return &ServiceFactory{
		bookUseCaseFactory: bookUseCaseFactory,
		userUseCaseFactory: userUseCaseFactory,
	}
}

func (f *ServiceFactory) CreateBookService() *BookService {
	return NewBookService(f.bookUseCaseFactory.CreateBookUseCase())
}

func (f *ServiceFactory) CreateUserService() *UserService {
	return NewUserService(f.userUseCaseFactory.CreateUserUseCase())
}
