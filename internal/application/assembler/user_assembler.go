package assembler

import (
	"project-layout/internal/application/dto"
	"project-layout/internal/domain/model"
)

// UserAssembler 用于在 DTO 和领域模型之间转换
type UserAssembler struct{}

// ToDTO 将领域模型转换为 DTO
func (a *UserAssembler) ToDTO(user *model.User) *dto.UserDTO {
	if user == nil {
		return nil
	}
	return &dto.UserDTO{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}

// ToDTOList 将领域模型列表转换为 DTO 列表
func (a *UserAssembler) ToDTOList(users []*model.User) []*dto.UserDTO {
	dtos := make([]*dto.UserDTO, len(users))
	for i, user := range users {
		dtos[i] = a.ToDTO(user)
	}
	return dtos
}

// ToModel 将创建 DTO 转换为领域模型
func (a *UserAssembler) ToModel(createDTO *dto.CreateUserDTO) *model.User {
	return &model.User{
		Name:  createDTO.Name,
		Email: createDTO.Email,
	}
}

// UpdateModel 使用更新 DTO 更新领域模型
func (a *UserAssembler) UpdateModel(user *model.User, updateDTO *dto.UpdateUserDTO) {
	user.Name = updateDTO.Name
	user.Email = updateDTO.Email
} 