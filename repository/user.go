package repository

import (
	"time"
)

type User struct {
	ID      string `gorm:"primaryKey;size:10"`
	Name    string
	Balance int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserRepository interface {
	GetAll() ([]User, error)
	GetById(string) (*User, error)
}
