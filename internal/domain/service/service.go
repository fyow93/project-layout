package service

import (
	"project-layout/internal/domain/model"
	"project-layout/internal/domain/repository"
)

type DomainService struct {
	repo repository.Repository
}

func NewDomainService(repo repository.Repository) *DomainService {
	return &DomainService{repo: repo}
}

func (s *DomainService) Save(entity model.Entity) error {
	// business logic
	return s.repo.Save(entity)
}

func (s *DomainService) FindByID(id string) (*model.Entity, error) {
	// business logic
	return s.repo.FindByID(id)
}
