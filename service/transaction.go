package service

type TransactionResponse struct {
	TransactionId int `json:"transaction_id"`
	Sender        string `json:"sender"`
	Receiver      string `json:"receiver"`
	Amount        int    `json:"amount"`
	CreatedAt     string `json:"created_at"`
}

type TransactionRequest struct {
	Sender        string `json:"sender"`
	Receiver      string `json:"receiver"`
	Amount        int    `json:"amount"`
}

type TransactionService interface {
	GetTransactions() ([]TransactionResponse, error)
	GetTransaction(int) (*TransactionResponse, error)
	GetTransactionsBySender(string) ([]TransactionResponse, error)
	GetTransactionsByReceiver(string) ([]TransactionResponse, error)
	CreateTransaction(TransactionRequest) (*TransactionResponse, error)
}
