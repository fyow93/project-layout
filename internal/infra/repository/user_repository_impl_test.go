package repository

import (
	"database/sql"
	"project-layout/internal/domain/model"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func setupTestUserDB(t *testing.T) *UserRepositoryImpl {
	db, err := sql.Open("sqlite3", "file::memory:?cache=shared")
	if err != nil {
		t.Fatalf("Failed to initialize database: %v", err)
	}

	repo := NewUserRepositoryImpl(db)
	err = repo.Initialize()
	if err != nil {
		t.Fatalf("Failed to create table: %v", err)
	}

	// Register cleanup function
	t.Cleanup(func() {
		repo.DeleteAllUsers()
		repo.Shutdown()
	})

	return repo
}

func TestUserRepositoryImpl_Save(t *testing.T) {
	repo := setupTestUserDB(t)

	user := model.User{ID: "1", Name: "Test User", Email: "test@example.com"}
	err := repo.Save(user)
	assert.NoError(t, err)

	savedUser, err := repo.FindByID("1")
	assert.NoError(t, err)
	assert.Equal(t, user, *savedUser)
}

func TestUserRepositoryImpl_FindByID(t *testing.T) {
	repo := setupTestUserDB(t)

	user := model.User{ID: "1", Name: "Test User", Email: "test@example.com"}
	err := repo.Save(user)
	assert.NoError(t, err)

	foundUser, err := repo.FindByID("1")
	assert.NoError(t, err)
	assert.Equal(t, user, *foundUser)
}

func TestUserRepositoryImpl_Update(t *testing.T) {
	repo := setupTestUserDB(t)

	user := model.User{ID: "1", Name: "Test User", Email: "test@example.com"}
	err := repo.Save(user)
	assert.NoError(t, err)

	user.Name = "Updated User"
	user.Email = "updated@example.com"
	err = repo.Update(user)
	assert.NoError(t, err)

	updatedUser, err := repo.FindByID("1")
	assert.NoError(t, err)
	assert.Equal(t, user, *updatedUser)
}

func TestUserRepositoryImpl_Delete(t *testing.T) {
	repo := setupTestUserDB(t)

	user := model.User{ID: "1", Name: "Test User", Email: "test@example.com"}
	err := repo.Save(user)
	assert.NoError(t, err)

	err = repo.Delete("1")
	assert.NoError(t, err)

	deletedUser, err := repo.FindByID("1")
	assert.Error(t, err)
	assert.Nil(t, deletedUser)
}
