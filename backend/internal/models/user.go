package models

import (
	. "backend/internal/resources/constants"
)

type User struct {
	BaseModel
	FirstName    string `gorm:"size:100;not null"`        
	MiddleName   string `gorm:"size:100"`                 
	LastName     string `gorm:"size:100;not null"`        
	Email        string `gorm:"size:100;unique;not null"` 
	Mobile       string `gorm:"size:20; unique; not null"`
	UserType     USERROLE `gorm:"not null;default:-1"`
	PasswordHash string `gorm:"not null"`
	CreatedBy    int `gorm:"size:50"`
	UpdatedBy    int `gorm:"size:50"`

	// GymPartner   GymPartner `gorm:"foreignkey:GymPartnerID"` // This refers to the GymPartner model, and GORM will use the field name `GymPartnerID` in the User table
	GymPartner []GymPartner `gorm:"foreignKey:UserID"`

}
