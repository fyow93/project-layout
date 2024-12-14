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

func (u *UseCase) Create(entity model.Entity) error {
	// use case logic
	return u.domainService.Save(entity)
}

func (u *UseCase) FindByID(id string) (*model.Entity, error) {
	// use case logic
	return u.domainService.FindByID(id)
}
