package repository

import (
	"time"
)

type User struct {
	ID      string `gorm:"primaryKey;size:10"`
	Name    string
	Balance int
	Password string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserRepository interface {
	GetAll() ([]User, error)
	GetById(string) (*User, error)
	Create(string, string, string) (*User, error)
	AddBalance(string, int) (*User, error)
	SubtractBalance(string, int) (*User, error)
}
