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

func (s *ApplicationService) Execute(entity model.Entity) error {
	// application logic
	return s.useCase.PerformAction(entity)
}
