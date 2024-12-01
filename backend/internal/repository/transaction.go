package repository

import (
	"personal-finance-app/internal/models"
	"time"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreateTransaction(transaction *models.Transaction) error
	GetTransactionsByUser(userID uint, filter *TransactionFilter) ([]*models.Transaction, error)
	GetTransactionByID(transactionID uint) (*models.Transaction, error)
	UpdateTransaction(transaction *models.Transaction) error
	DeleteTransaction(transactionID uint) error
}

// TransactionFilter - структура для фильтрации транзакций
type TransactionFilter struct {
	StartDate  time.Time
	EndDate    time.Time
	CategoryID *uint
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db: db}
}

func (r *transactionRepository) CreateTransaction(transaction *models.Transaction) error {
	return r.db.Create(transaction).Error
}

func (r *transactionRepository) GetTransactionsByUser(userID uint, filter *TransactionFilter) ([]*models.Transaction, error) {
	query := r.db.Where("user_id = ?", userID)
	if filter != nil {
		if !filter.StartDate.IsZero() {
			query = query.Where("transaction_date >= ?", filter.StartDate)
		}
		if !filter.EndDate.IsZero() {
			query = query.Where("transaction_date <= ?", filter.EndDate)
		}
		if filter.CategoryID != nil {
			query = query.Where("category_id = ?", *filter.CategoryID)
		}
	}
	var transactions []*models.Transaction
	err := query.Find(&transactions).Error
	return transactions, err
}

func (r *transactionRepository) GetTransactionByID(transactionID uint) (*models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.First(&transaction, transactionID).Error
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

func (r *transactionRepository) UpdateTransaction(transaction *models.Transaction) error {
	return r.db.Save(transaction).Error
}

func (r *transactionRepository) DeleteTransaction(transactionID uint) error {
	return r.db.Delete(&models.Transaction{}, transactionID).Error
}
