package handler

import (
	"github.com/CUBS-sources-code/CUBS-coin/service"
	"github.com/gofiber/fiber/v2"
)

type transactionHandler struct {
	transactionService service.TransactionService
}

func NewTransactionHandler(transactionService service.TransactionService) transactionHandler {
	return transactionHandler{transactionService: transactionService}
}

// GetTransactionsBySender(string) ([]TransactionResponse, error)
// GetTransactionsByReceiver(string) ([]TransactionResponse, error)

func (h transactionHandler) GetTransactions(c *fiber.Ctx) error {
	
	transactions, err := h.transactionService.GetTransactions()
	if err != nil {
		return handlerError(c, err)
	}

	return c.JSON(fiber.Map{
		"transactions": transactions,
	})
}

func (h transactionHandler) GetTransaction(c *fiber.Ctx) error {
	
	id, err := c.ParamsInt("id")
    if err != nil {
        return fiber.ErrBadRequest
    }

	transaction, err := h.transactionService.GetTransaction(id)
	if err != nil {
		return handlerError(c, err)
	}
	return c.JSON(transaction)
}

func (h transactionHandler) CreateTransaction(c *fiber.Ctx) error {
	
	var request service.TransactionRequest
    if err := c.BodyParser(&request); err != nil {
       return handlerError(c, err)
    }

	transaction, err := h.transactionService.CreateTransaction(request)
	if err != nil {
		return handlerError(c, err)
	}

	return c.JSON(transaction)
}