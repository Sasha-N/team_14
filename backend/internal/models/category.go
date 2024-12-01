package models

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	UserID    uint           `gorm:"not null"`
	Name      string         `gorm:"not null;uniqueIndex:idx_user_category_name"`
}
