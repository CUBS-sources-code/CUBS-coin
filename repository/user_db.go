package repository

import (
	"gorm.io/gorm"
)

type userRepositoryDB struct {
	db *gorm.DB
}

func NewUserRepositoryDB(db *gorm.DB) userRepositoryDB {
	return userRepositoryDB{db: db}
}

func (r userRepositoryDB) GetAll() ([]User, error) {

	users := []User{}
	
	// query
	tx := r.db.Find(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return users, nil
}

func (r userRepositoryDB) GetById(id string) (*User, error) {

	user := User{}
	tx := r.db.First(&user, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &user, nil
}