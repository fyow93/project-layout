package service

import (
	"project-layout/internal/application/usecases"
	"project-layout/internal/domain/model"
)

type ApplicationService struct {
	useCase *usecases.UseCase
}

func NewApplicationService(useCase *usecases.UseCase) *ApplicationService {
	return &ApplicationService{useCase: useCase}
}

func (s *ApplicationService) ExecuteFindByID(id string) (*model.Entity, error) {
	return s.useCase.FindByID(id)
}

func (s *ApplicationService) ExecuteCreate(entity model.Entity) error {
	return s.useCase.Create(entity)
}
