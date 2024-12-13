package repository

import "project-layout/internal/domain/model"

type Repository interface {
	Save(entity model.Entity) error
	FindByID(id string) (*model.Entity, error)
}
