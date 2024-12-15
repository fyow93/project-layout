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

type UserRepositoryImpl struct {
    db *sql.DB
}

func NewUserRepositoryImpl(db *sql.DB) *UserRepositoryImpl {
    return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) Initialize() error {
    _, err := r.db.Exec(`CREATE TABLE IF NOT EXISTS users (
        id TEXT PRIMARY KEY,
        name TEXT,
        email TEXT
    )`)
    return err
}

func (r *UserRepositoryImpl) Save(user model.User) error {
    _, err := r.db.Exec("INSERT INTO users (id, name, email) VALUES (?, ?, ?)", user.ID, user.Name, user.Email)
    return err
}

func (r *UserRepositoryImpl) FindByID(id string) (*model.User, error) {
    row := r.db.QueryRow("SELECT id, name, email FROM users WHERE id = ?", id)
    var user model.User
    err := row.Scan(&user.ID, &user.Name, &user.Email)
    if err != nil {
        return nil, err
    }
    return &user, nil
}

func (r *UserRepositoryImpl) Update(user model.User) error {
    _, err := r.db.Exec("UPDATE users SET name = ?, email = ? WHERE id = ?", user.Name, user.Email, user.ID)
    return err
}

func (r *UserRepositoryImpl) Delete(id string) error {
    _, err := r.db.Exec("DELETE FROM users WHERE id = ?", id)
    return err
}

func (r *UserRepositoryImpl) DeleteAllUsers() error {
    _, err := r.db.Exec("DELETE FROM users")
    return err
}

func (r *UserRepositoryImpl) Shutdown() error {
    return r.db.Close()
}
