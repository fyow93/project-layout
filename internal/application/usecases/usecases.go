package usecases

import (
	"project-layout/internal/domain/model"
	"project-layout/internal/domain/service"
)

type UseCase struct {
	domainService *service.DomainService
}

func NewUseCase(domainService *service.DomainService) *UseCase {
	return &UseCase{domainService: domainService}
}

func (u *UseCase) PerformAction(entity model.Entity) error {
	// use case logic
	return u.domainService.DoSomething(entity)
}
