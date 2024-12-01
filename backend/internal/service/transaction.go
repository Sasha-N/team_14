package service

import (
	"personal-finance-app/internal/models"
	"personal-finance-app/internal/repository"
	"time"
)

type TransactionService interface {
	CreateTransaction(userID uint, amount int64, categoryID *uint, transactionDate time.Time, transactionType string) error
	GetTransactions(userID uint, filter *repository.TransactionFilter) ([]*models.Transaction, error)
	GetTransactionByID(transactionID uint) (*models.Transaction, error)
	UpdateTransaction(transactionID uint, amount int64, categoryID *uint, transactionDate time.Time, transactionType string) error
	DeleteTransaction(transactionID uint) error
}

type transactionService struct {
	repo repository.TransactionRepository
}

func NewTransactionService(repo repository.TransactionRepository) TransactionService {
	return &transactionService{repo: repo}
}

func (s *transactionService) CreateTransaction(userID uint, amount int64, categoryID *uint, transactionDate time.Time, transactionType string) error {
	transaction := &models.Transaction{
		UserID:          userID,
		Amount:          amount,
		CategoryID:      categoryID,
		TransactionDate: transactionDate,
		Type:            transactionType,
	}
	return s.repo.CreateTransaction(transaction)
}

func (s *transactionService) GetTransactions(userID uint, filter *repository.TransactionFilter) ([]*models.Transaction, error) {
	return s.repo.GetTransactionsByUser(userID, filter)
}

func (s *transactionService) GetTransactionByID(transactionID uint) (*models.Transaction, error) {
	return s.repo.GetTransactionByID(transactionID)
}

func (s *transactionService) UpdateTransaction(transactionID uint, amount int64, categoryID *uint, transactionDate time.Time, transactionType string) error {
	transaction, err := s.repo.GetTransactionByID(transactionID)
	if err != nil {
		return err
	}
	transaction.Amount = amount
	transaction.CategoryID = categoryID
	transaction.TransactionDate = transactionDate
	transaction.Type = transactionType
	return s.repo.UpdateTransaction(transaction)
}

func (s *transactionService) DeleteTransaction(transactionID uint) error {
	return s.repo.DeleteTransaction(transactionID)
}
