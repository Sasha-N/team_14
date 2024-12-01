package handlers

import (
	"net/http"
	"personal-finance-app/internal/models"
	"personal-finance-app/internal/repository"
	"personal-finance-app/internal/service"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	transactionService service.TransactionService
}

func NewTransactionHandler(transactionService service.TransactionService) *TransactionHandler {
	return &TransactionHandler{transactionService: transactionService}
}

func (h *TransactionHandler) CreateTransaction(c *gin.Context) {
	value, exists := c.Get("user_id")
	if !exists {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "user_id not found in context"})
		return
	}
	userID, ok := value.(uint)
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "invalid user_id type"})
		return
	}

	var transaction struct {
		Amount          int64  `json:"amount" binding:"required"`
		CategoryID      uint   `json:"category_id" binding:"required"`
		TransactionDate string `json:"transaction_date" binding:"required"`
		Type            string `json:"type" binding:"required"`
	}

	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transactionDate, err := time.Parse("2006-01-02", transaction.TransactionDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	categoryID := transaction.CategoryID
	categoryIDPtr := &categoryID

	newTransaction := models.Transaction{
		UserID:          userID,
		Amount:          transaction.Amount,
		CategoryID:      categoryIDPtr,
		TransactionDate: transactionDate,
		Type:            transaction.Type,
	}

	if err := h.transactionService.CreateTransaction(&newTransaction); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Transaction created successfully"})
}

func (h *TransactionHandler) GetTransactions(c *gin.Context) {
	value, exists := c.Get("user_id")
	if !exists {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "user_id not found in context"})
		return
	}
	userID, ok := value.(uint)
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "invalid user_id type"})
		return
	}

	filter := &repository.TransactionFilter{}
	if startDate, err := time.Parse(time.RFC3339, c.Query("startDate")); err == nil {
		filter.StartDate = startDate
	}
	if endDate, err := time.Parse(time.RFC3339, c.Query("endDate")); err == nil {
		filter.EndDate = endDate
	}
	if categoryID, err := strconv.Atoi(c.Query("categoryID")); err == nil {
		catID := uint(categoryID)
		filter.CategoryID = &catID
	}

	transactions, err := h.transactionService.GetTransactions(userID, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, transactions)
}

func (h *TransactionHandler) GetTransaction(c *gin.Context) {
	transactionID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid transaction ID"})
		return
	}

	transaction, err := h.transactionService.GetTransactionByID(uint(transactionID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, transaction)
}

func (h *TransactionHandler) UpdateTransaction(c *gin.Context) {
	transactionID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid transaction ID"})
		return
	}

	var transactionUpdate models.Transaction
	if err := c.ShouldBindJSON(&transactionUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Проверка, что пользователь может редактировать транзакцию (проверка userID)
	value, exists := c.Get("user_id")
	if !exists {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "user_id not found in context"})
		return
	}
	userID, ok := value.(uint)
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "invalid user_id type"})
		return
	}
	transaction, err := h.transactionService.GetTransactionByID(uint(transactionID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if transaction.UserID != userID {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "You are not allowed to update this transaction"})
		return
	}

	err = h.transactionService.UpdateTransaction(uint(transactionID), transactionUpdate.Amount, transactionUpdate.CategoryID, transactionUpdate.TransactionDate, transactionUpdate.Type)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Transaction updated successfully"})
}

func (h *TransactionHandler) DeleteTransaction(c *gin.Context) {
	transactionID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid transaction ID"})
		return
	}

	//Проверка, что пользователь может удалить транзакцию (проверка userID)
	value, exists := c.Get("user_id")
	if !exists {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "user_id not found in context"})
		return
	}
	userID, ok := value.(uint)
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "invalid user_id type"})
		return
	}
	transaction, err := h.transactionService.GetTransactionByID(uint(transactionID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if transaction.UserID != userID {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "You are not allowed to delete this transaction"})
		return
	}

	err = h.transactionService.DeleteTransaction(uint(transactionID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Transaction deleted successfully"})
}
