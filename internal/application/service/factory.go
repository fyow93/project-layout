package service

import (
	"project-layout/internal/application/usecases"
)

type ServiceFactory struct {
	useCaseFactory *usecases.UseCaseFactory
}

func NewServiceFactory(useCaseFactory *usecases.UseCaseFactory) *ServiceFactory {
	return &ServiceFactory{useCaseFactory: useCaseFactory}
}

func (f *ServiceFactory) CreateApplicationService() *ApplicationService {
	return NewApplicationService(f.useCaseFactory.CreateUseCase())
}

// 你可以在这里添加更多的服务初始化方法
