package repository

import (
	"database/sql"
	"project-layout/internal/domain/model"

	_ "github.com/mattn/go-sqlite3"
)

type RepositoryImpl struct {
	db *sql.DB
}

func NewRepositoryImpl(dataSourceName string) (*RepositoryImpl, error) {
	db, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return nil, err
	}
	return &RepositoryImpl{db: db}, nil
}

func (r *RepositoryImpl) Initialize() error {
	_, err := r.db.Exec(`CREATE TABLE IF NOT EXISTS entities (
		id TEXT PRIMARY KEY,
		name TEXT
	)`)
	return err
}

func (r *RepositoryImpl) Save(entity model.Entity) error {
	_, err := r.db.Exec("INSERT INTO entities (id, name) VALUES (?, ?)", entity.ID, entity.Name)
	return err
}

func (r *RepositoryImpl) FindByID(id string) (*model.Entity, error) {
	row := r.db.QueryRow("SELECT id, name FROM entities WHERE id = ?", id)
	var entity model.Entity
	err := row.Scan(&entity.ID, &entity.Name)
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *RepositoryImpl) Update(entity model.Entity) error {
	_, err := r.db.Exec("UPDATE entities SET name = ? WHERE id = ?", entity.Name, entity.ID)
	return err
}

func (r *RepositoryImpl) Delete(id string) error {
	_, err := r.db.Exec("DELETE FROM entities WHERE id = ?", id)
	return err
}

func (r *RepositoryImpl) DeleteAllEntities() error {
	_, err := r.db.Exec("DELETE FROM entities")
	return err
}

func (r *RepositoryImpl) Shutdown() error {
	return r.db.Close()
}
