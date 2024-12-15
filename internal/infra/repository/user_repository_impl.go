package repository

import (
	"errors"
	"sync"

	"project-layout/internal/domain/model"
)

type UserRepositoryImpl struct {
	mu    sync.Mutex
	users map[string]model.User
}

func NewUserRepository() *UserRepositoryImpl {
	return &UserRepositoryImpl{
		users: make(map[string]model.User),
	}
}

func (r *UserRepositoryImpl) CreateUser(user model.User) model.User {
	r.mu.Lock()
	defer r.mu.Unlock()
	user.ID = generateID()
	r.users[user.ID] = user
	return user
}

func (r *UserRepositoryImpl) GetUser(id string) (model.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	user, exists := r.users[id]
	if !exists {
		return model.User{}, errors.New("user not found")
	}
	return user, nil
}

func (r *UserRepositoryImpl) UpdateUser(id string, updatedUser model.User) (model.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	user, exists := r.users[id]
	if !exists {
		return model.User{}, errors.New("user not found")
	}
	if updatedUser.Name != "" {
		user.Name = updatedUser.Name
	}
	if updatedUser.Email != "" {
		user.Email = updatedUser.Email
	}
	r.users[id] = user
	return user, nil
}

func (r *UserRepositoryImpl) DeleteUser(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	_, exists := r.users[id]
	if !exists {
		return errors.New("user not found")
	}
	delete(r.users, id)
	return nil
}

func generateID() string {
	// Implement a function to generate a unique ID
	return "some-unique-id"
}
