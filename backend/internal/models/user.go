package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           uint           `gorm:"primaryKey"`               // Auto-increment primary key
	FirstName    string         `gorm:"size:100;not null"`        // First name, max size 100
	MiddleName   string         `gorm:"size:100"`                 // Middle name, optional
	LastName     string         `gorm:"size:100;not null"`        // Last name, max size 100
	Email        string         `gorm:"size:100;unique;not null"` // Email ID, unique and mandatory
	PasswordHash string         `gorm:"not null"`                 // Hashed password, mandatory
	CreatedBy    string         `gorm:"size:50"`                  // Who created the record
	UpdatedBy    string         `gorm:"size:50"`                  // Who last updated the record
	CreatedAt    time.Time      // Automatically managed creation timestamp
	UpdatedAt    time.Time      // Automatically managed update timestamp
	DeletedAt    gorm.DeletedAt `gorm:"index"` // Soft delete timestamp
}
