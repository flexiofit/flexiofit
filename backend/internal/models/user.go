package models

import (
	. "backend/internal/resources/constants"
)

type User struct {
	BaseModel
	FirstName    string `gorm:"size:100;not null"`        // First name, max size 100
	MiddleName   string `gorm:"size:100"`                 // Middle name, optional
	LastName     string `gorm:"size:100;not null"`        // Last name, max size 100
	Email        string `gorm:"size:100;unique;not null"` // Email ID, unique and mandatory
	UserType     USERROLE `gorm:"not null;default:-1"`
	PasswordHash string `gorm:"not null"`                 // Hashed password, mandatory
	CreatedBy    int `gorm:"size:50"`                  // Who created the record
	UpdatedBy    int `gorm:"size:50"`                  // Who last updated the record

	// GymPartner   GymPartner `gorm:"foreignkey:GymPartnerID"` // This refers to the GymPartner model, and GORM will use the field name `GymPartnerID` in the User table

}
