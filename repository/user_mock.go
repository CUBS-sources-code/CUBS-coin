package repository

import (
	"errors"
	"time"
)

type userRepositoryMock struct {
	users []User
}

func NewUserRepositoryMock() userRepositoryMock {
	users := []User{
		{
			ID:      "1234567890",
			Name:    "John",
			Balance: 100,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:      "0987654321",
			Name:    "Jane",
			Balance: 200,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	return userRepositoryMock{users: users}
}

func (r userRepositoryMock) GetAll() ([]User, error) {
	return r.users, nil
}

func (r userRepositoryMock) GetById(id string) (*User, error) {
	for _, user := range r.users {
		if user.ID == id {
			return &user, nil
		}
	}
	return nil, errors.New("user not found")
}