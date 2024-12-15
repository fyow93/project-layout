package model

import (
    "github.com/go-playground/validator/v10"
)

// Book 结构体表示一本书
type Book struct {
    ID     string `json:"id" validate:"required,uuid4"`
    Title  string `json:"title" validate:"required"`
    Author string `json:"author" validate:"required"`
}

// ValidateBook 验证 Book 实例
func ValidateBook(book *Book) error {
    validate := validator.New()
    return validate.Struct(book)
}
