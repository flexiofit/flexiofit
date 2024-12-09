package models

import (
    . "backend/internal/resources/constants"
)

type User struct {
    BaseModel
    FirstName    string       `gorm:"column:first_name;size:25;not null"`        
    MiddleName   string       `gorm:"column:middle_name;size:25"`                 
    LastName     string       `gorm:"column:last_name;size:25;not null"`        
    Email        string       `gorm:"column:email;size:100;unique;not null"` 
    IsActive          bool    `gorm:"column:is_active;default:true"`
    Username     string       `gorm:"column:username;size:100;unique"` 
    Mobile       string       `gorm:"column:mobile;size:20;unique;not null"`
    UserType     USERROLE `gorm:"column:user_type;not null;default:-1"`
    PasswordHash string       `gorm:"column:password_hash;not null"`
    CreatedBy    int          `gorm:"column:created_by"`
    UpdatedBy    int          `gorm:"column:updated_by"`
}
