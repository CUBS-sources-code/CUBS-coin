package service

import (
	"fmt"

	"github.com/CUBS-sources-code/CUBS-coin/errs"
	"github.com/CUBS-sources-code/CUBS-coin/logs"
	"github.com/CUBS-sources-code/CUBS-coin/repository"
	"gorm.io/gorm"
)

type transactionService struct {
	transactionRepository repository.TransactionRepository
	userRepository repository.UserRepository
}

func NewTransactionService(transactionRepository repository.TransactionRepository, userRepository repository.UserRepository) transactionService {
	return transactionService{transactionRepository: transactionRepository,
		userRepository: userRepository}
}

func (s transactionService) GetTransactions() ([]TransactionResponse, error) {

	transactions, err := s.transactionRepository.GetAll()
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	transactionResponses := []TransactionResponse{}
	for _, transaction := range transactions {
		transactionResponse := TransactionResponse{
			TransactionId: transaction.ID,
			Sender: 	transaction.Sender,
			Receiver: transaction.Receiver,
			Amount: transaction.Amount,
			CreatedAt: transaction.CreatedAt.String(),
		}
		transactionResponses = append(transactionResponses, transactionResponse)
	}

	return transactionResponses, nil
}

func (s transactionService) GetTransaction(id int) (*TransactionResponse, error) {
	transaction, err := s.transactionRepository.GetById(id)
	if err != nil {

		if err == gorm.ErrRecordNotFound {
			fmt.Println("err")
			logs.Error(err)
			return nil, errs.NewNotFoundError("transaction not found")
		}

		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	transactionResponse := TransactionResponse{
		TransactionId: transaction.ID,
		Sender: 	transaction.Sender,
		Receiver: transaction.Receiver,
		Amount: transaction.Amount,
		CreatedAt: transaction.CreatedAt.String(),
	}
	return &transactionResponse, nil
}

func (s transactionService) GetTransactionsBySender(sender_id string) ([]TransactionResponse, error) {
	transactions, err := s.transactionRepository.GetBySender(sender_id)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	transactionResponses := []TransactionResponse{}
	for _, transaction := range transactions {
		transactionResponse := TransactionResponse{
			TransactionId: transaction.ID,
			Sender:        transaction.Sender,
			Receiver:      transaction.Receiver,
			Amount:        transaction.Amount,
			CreatedAt:     transaction.CreatedAt.String(),
		}
		transactionResponses = append(transactionResponses, transactionResponse)
	}

	return transactionResponses, nil
}

func (s transactionService) GetTransactionsByReceiver(receiver_id string) ([]TransactionResponse, error) {
	transactions, err := s.transactionRepository.GetByReceiver(receiver_id)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	transactionResponses := []TransactionResponse{}
	for _, transaction := range transactions {
		transactionResponse := TransactionResponse{
			TransactionId: transaction.ID,
			Sender:        transaction.Sender,
			Receiver:      transaction.Receiver,
			Amount:        transaction.Amount,
			CreatedAt:     transaction.CreatedAt.String(),
		}
		transactionResponses = append(transactionResponses, transactionResponse)
	}

	return transactionResponses, nil
}

func (s transactionService) CreateTransaction(transactionRequest TransactionRequest) (*TransactionResponse, error) {

	sender_id := transactionRequest.Sender
	receiver_id := transactionRequest.Receiver
	amount := transactionRequest.Amount

	if sender_id == "" || receiver_id == "" || amount <= 0 {
		return nil, errs.NewBadRequestError("invalid transaction request")
	}

	// check is sender exist
	sender, serr := s.userRepository.GetById(sender_id)
	if serr != nil {
		logs.Error(serr)
		return nil, errs.NewNotFoundError("sender doesn't existed")
	}

	// check is receiver exist
	_, rerr := s.userRepository.GetById(receiver_id)
	if rerr != nil {
		logs.Error(rerr)
		return nil, errs.NewNotFoundError("receiver doesn't existed")
	}

	// check if sender has enough balance
	if sender.Balance < amount {
		logs.Error("insufficient balance")
		return nil, errs.NewForbiddenError("insufficient balance")
	}
	
	// create transaction
	transaction, err := s.transactionRepository.Create(sender_id, receiver_id, amount)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	// update sender balance
	_, err = s.userRepository.SubtractBalance(sender_id, amount)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	// update reciever balance
	_, err = s.userRepository.AddBalance(receiver_id, amount)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	transactionResponse := &TransactionResponse{
		TransactionId: transaction.ID,
		Sender:        transaction.Sender,
		Receiver:      transaction.Receiver,
		Amount:        transaction.Amount,
		CreatedAt:     transaction.CreatedAt.String(),
	}

	return transactionResponse, nil
}