package repository

import (
	"gorm.io/gorm"
)

type userRepositoryDB struct {
	db *gorm.DB
}

func NewUserRepositoryDB(db *gorm.DB) userRepositoryDB {
	db.AutoMigrate(User{})
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

func (r userRepositoryDB) AddBalance(id string, amount int) (*User, error) {

	// Get data
	user := User{}
	tx := r.db.First(&user, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	
	// Update data
	user.Balance += amount
	tx = r.db.Save(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &user, nil
}

func (r userRepositoryDB) SubtractBalance(id string, amount int) (*User, error) {

	// Get data
	user := User{}
	tx := r.db.First(&user, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	
	// Update data
	user.Balance -= amount
	tx = r.db.Save(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &user, nil
}

func (r userRepositoryDB) Create(id string, name string, password string) (*User, error) {

	user := User{
		ID: id,
		Name: name,
		Balance: 0,
		Password: password,
		Role: "member",
	}

	tx := r.db.Create(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &user, nil
}

func (r userRepositoryDB) ChangeRoleToAdmin(id string) (*User, error) {

	// Get data
	user := User{}
	tx := r.db.First(&user, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	
	// Update data
	user.Role = "admin"
	tx = r.db.Save(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &user, nil
}

func (r userRepositoryDB) ChangeRoleToMember(id string) (*User, error) {

	// Get data
	user := User{}
	tx := r.db.First(&user, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	
	// Update data
	user.Role = "member"
	tx = r.db.Save(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &user, nil
}