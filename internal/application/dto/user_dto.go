package dto

// UserDTO 用于在应用层传输用户数据
type UserDTO struct {
	ID    string `json:"id"`
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

// CreateUserDTO 用于创建用户的请求数据
type CreateUserDTO struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

// UpdateUserDTO 用于更新用户的请求数据
type UpdateUserDTO struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
} 