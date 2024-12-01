package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID              uint `gorm:"primaryKey"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
	UserID          uint           `gorm:"not null"`
	Amount          int64          `gorm:"not null"`
	CategoryID      *uint          `gorm:"null"`
	TransactionDate time.Time      `gorm:"not null;default:CURRENT_DATE"`
	Type            string         `gorm:"not null;check:type IN ('income', 'expense')"`
}
