package repository

import (
	"gorm.io/gorm"
)

type transactionRepositoryDB struct {
	db *gorm.DB
}

func NewTransactionRepositoryDB(db *gorm.DB) transactionRepositoryDB {
	db.AutoMigrate(Transaction{})
	return transactionRepositoryDB{db: db}
}

func (r transactionRepositoryDB) GetAll() ([]Transaction, error) {

	transactions := []Transaction{}	
	tx := r.db.Find(&transactions)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return transactions, nil
}

func (r transactionRepositoryDB) GetById(id int) (*Transaction, error) {

	transaction := Transaction{}
	tx := r.db.First(&transaction, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &transaction, nil
}

func (r transactionRepositoryDB) GetBySender(sender_id string) ([]Transaction, error) {

	transactions := []Transaction{}	
	tx := r.db.Where("sender <> ?", sender_id).Find(&transactions)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return transactions, nil
}

func (r transactionRepositoryDB) GetByReceiver(reciever_id string) ([]Transaction, error) {
	
	transactions := []Transaction{}	
	tx := r.db.Where("reciever <> ?", reciever_id).Find(&transactions)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return transactions, nil
}

func (r transactionRepositoryDB) Create(sender_id string, reciever_id string, amount int) (*Transaction, error) {

	transaction := Transaction{
		Sender: sender_id,
		Receiver: reciever_id,
		Amount: amount,
	}
	tx := r.db.Create(&transaction)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &transaction, nil
}