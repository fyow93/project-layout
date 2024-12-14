package usecases

import (
	"project-layout/internal/domain/service"
)

type UseCaseFactory struct {
	domainService *service.DomainService
}

func NewUseCaseFactory(domainService *service.DomainService) *UseCaseFactory {
	return &UseCaseFactory{domainService: domainService}
}

func (f *UseCaseFactory) CreateUseCase() *UseCase {
	return NewUseCase(f.domainService)
}

// 你可以在这里添加更多的用例初始化方法
