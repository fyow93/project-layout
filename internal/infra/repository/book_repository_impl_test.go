package repository

import (
    "database/sql"
    "project-layout/internal/domain/model"
    "testing"

    _ "github.com/mattn/go-sqlite3"
    "github.com/stretchr/testify/assert"
)

func setupTestBookDB(t *testing.T) *BookRepositoryImpl {
    db, err := sql.Open("sqlite3", "file::memory:?cache=shared")
    if err != nil {
        t.Fatalf("Failed to initialize database: %v", err)
    }

    repo := NewBookRepositoryImpl(db)
    err = repo.Initialize()
    if err != nil {
        t.Fatalf("Failed to create table: %v", err)
    }

    // Register cleanup function
    t.Cleanup(func() {
        repo.DeleteAllBooks()
        repo.Shutdown()
    })

    return repo
}

func TestBookRepositoryImpl_Save(t *testing.T) {
    repo := setupTestBookDB(t)

    book := model.Book{ID: "1", Title: "Test Book", Author: "Test Author"}
    err := repo.Save(book)
    assert.NoError(t, err)

    savedBook, err := repo.FindByID("1")
    assert.NoError(t, err)
    assert.Equal(t, book, *savedBook)
}

func TestBookRepositoryImpl_FindByID(t *testing.T) {
    repo := setupTestBookDB(t)

    book := model.Book{ID: "1", Title: "Test Book", Author: "Test Author"}
    err := repo.Save(book)
    assert.NoError(t, err)

    foundBook, err := repo.FindByID("1")
    assert.NoError(t, err)
    assert.Equal(t, book, *foundBook)
}

func TestBookRepositoryImpl_Update(t *testing.T) {
    repo := setupTestBookDB(t)

    book := model.Book{ID: "1", Title: "Test Book", Author: "Test Author"}
    err := repo.Save(book)
    assert.NoError(t, err)

    book.Title = "Updated Book"
    book.Author = "Updated Author"
    err = repo.Update(book)
    assert.NoError(t, err)

    updatedBook, err := repo.FindByID("1")
    assert.NoError(t, err)
    assert.Equal(t, book, *updatedBook)
}

func TestBookRepositoryImpl_Delete(t *testing.T) {
    repo := setupTestBookDB(t)

    book := model.Book{ID: "1", Title: "Test Book", Author: "Test Author"}
    err := repo.Save(book)
    assert.NoError(t, err)

    err = repo.Delete("1")
    assert.NoError(t, err)

    deletedBook, err := repo.FindByID("1")
    assert.Error(t, err)
    assert.Nil(t, deletedBook)
}
