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

func (s *DomainService) DoSomething(entity model.Entity) error {
	// business logic
	return s.repo.Save(entity)
}
