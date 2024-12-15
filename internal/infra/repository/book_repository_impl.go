// Package repository 包含具体的仓储实现。
// 这是具体的仓储实现层，负责数据的实际存储和检索。
// 它实现了领域仓储接口。
// 上一层：领域仓储接口
// 下一层：数据库或其他存储介质

package repository

import (
    "database/sql"
    "project-layout/internal/domain/model"

    _ "github.com/mattn/go-sqlite3"
)

type BookRepositoryImpl struct {
    db *sql.DB
}

func NewBookRepositoryImpl(db *sql.DB) *BookRepositoryImpl {
    return &BookRepositoryImpl{db: db}
}

func (r *BookRepositoryImpl) Initialize() error {
    _, err := r.db.Exec(`CREATE TABLE IF NOT EXISTS books (
        id TEXT PRIMARY KEY,
        title TEXT,
        author TEXT
    )`)
    return err
}

func (r *BookRepositoryImpl) Save(book model.Book) error {
    _, err := r.db.Exec("INSERT INTO books (id, title, author) VALUES (?, ?, ?)", book.ID, book.Title, book.Author)
    return err
}

func (r *BookRepositoryImpl) FindByID(id string) (*model.Book, error) {
    row := r.db.QueryRow("SELECT id, title, author FROM books WHERE id = ?", id)
    var book model.Book
    err := row.Scan(&book.ID, &book.Title, &book.Author)
    if err != nil {
        return nil, err
    }
    return &book, nil
}

func (r *BookRepositoryImpl) Update(book model.Book) error {
    _, err := r.db.Exec("UPDATE books SET title = ?, author = ? WHERE id = ?", book.Title, book.Author, book.ID)
    return err
}

func (r *BookRepositoryImpl) Delete(id string) error {
    _, err := r.db.Exec("DELETE FROM books WHERE id = ?", id)
    return err
}

func (r *BookRepositoryImpl) DeleteAllBooks() error {
    _, err := r.db.Exec("DELETE FROM books")
    return err
}

func (r *BookRepositoryImpl) Shutdown() error {
    return r.db.Close()
}
