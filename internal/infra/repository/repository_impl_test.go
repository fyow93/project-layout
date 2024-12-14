package repository

import (
	"project-layout/internal/domain/model"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func setupTestDB(t *testing.T) *RepositoryImpl {
	db, err := NewRepositoryImpl("file::memory:?cache=shared")
	if err != nil {
		t.Fatalf("Failed to initialize database: %v", err)
	}
	err = db.Initialize()
	if err != nil {
		t.Fatalf("Failed to create table: %v", err)
	}

	// Register cleanup function
	t.Cleanup(func() {
		db.db.Exec("DELETE FROM entities")
		db.db.Close()
	})

	return db
}

func TestRepositoryImpl_Save(t *testing.T) {
	repo := setupTestDB(t)

	entity := model.Entity{ID: "1", Name: "Test Entity"}
	err := repo.Save(entity)
	assert.NoError(t, err)

	savedEntity, err := repo.FindByID("1")
	assert.NoError(t, err)
	assert.Equal(t, entity, *savedEntity)
}

func TestRepositoryImpl_FindByID(t *testing.T) {
	repo := setupTestDB(t)

	entity := model.Entity{ID: "1", Name: "Test Entity"}
	err := repo.Save(entity)
	assert.NoError(t, err)

	foundEntity, err := repo.FindByID("1")
	assert.NoError(t, err)
	assert.Equal(t, entity, *foundEntity)
}

func TestRepositoryImpl_Update(t *testing.T) {
	repo := setupTestDB(t)

	entity := model.Entity{ID: "1", Name: "Test Entity"}
	err := repo.Save(entity)
	assert.NoError(t, err)

	entity.Name = "Updated Entity"
	err = repo.Update(entity)
	assert.NoError(t, err)

	updatedEntity, err := repo.FindByID("1")
	assert.NoError(t, err)
	assert.Equal(t, entity, *updatedEntity)
}

func TestRepositoryImpl_Delete(t *testing.T) {
	repo := setupTestDB(t)

	entity := model.Entity{ID: "1", Name: "Test Entity"}
	err := repo.Save(entity)
	assert.NoError(t, err)

	err = repo.Delete("1")
	assert.NoError(t, err)

	deletedEntity, err := repo.FindByID("1")
	assert.Error(t, err)
	assert.Nil(t, deletedEntity)
}
