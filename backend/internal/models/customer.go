package models

import (
	"time"
)

type Customer struct {
    BaseModel
    UserID           uint      `gorm:"column:user_id"`              // Foreign key to User
    BranchID         int       `gorm:"column:branch_id"`            // Foreign key to GymBranch
    FirstName        string    `gorm:"column:first_name;size:100"`
    MiddleName       string    `gorm:"column:middle_name;size:100"`
    LastName         string    `gorm:"column:last_name;size:100"`
    Email            string    `gorm:"column:email;size:100"`
    Mobile           string    `gorm:"column:mobile;size:20"`
    AlternateMobile  string    `gorm:"column:alternate_mobile;size:20"`
    DateOfBirth      time.Time `gorm:"column:date_of_birth"`
    MembershipStart  time.Time `gorm:"column:membership_start"`     // Membership start date
    MembershipEnd    time.Time `gorm:"column:membership_end"`       // Membership end date
    IsActive         bool      `gorm:"column:is_active"`            // Whether the membership is active

    GymBranch GymBranch `gorm:"foreignKey:BranchID"` // Each Customer belongs to one GymBranch
    User      User      `gorm:"foreignKey:UserID"`   // Relationship to User
}
