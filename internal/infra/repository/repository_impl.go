package repository

import "project-layout/internal/domain/model"

type RepositoryImpl struct {
	// database connection
}

func (r *RepositoryImpl) Save(entity model.Entity) error {
	// save entity to database
	return nil
}

func (r *RepositoryImpl) FindByID(id string) (*model.Entity, error) {
	// find entity by ID
	return &model.Entity{ID: id, Name: "Example"}, nil
}

func (r *RepositoryImpl) Update(entity model.Entity) error {
	// update entity in database
	return nil
}
func (r *RepositoryImpl) Delete(id string) error {
	// delete entity from database
	return nil
}
