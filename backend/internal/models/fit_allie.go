
package models

type FitAllie struct {
    BaseModel
    UserID             int    `gorm:"column:user_id"`
    OwnerFirstName     string `gorm:"column:owner_first_name;size:100;not null"`
    OwnerMiddleName    string `gorm:"column:owner_middle_name;size:100"`
    OwnerLastName      string `gorm:"column:owner_last_name;size:100"`
    CoOwnerFirstName   string `gorm:"column:co_owner_first_name;size:100"`
    CoOwnerMiddleName  string `gorm:"column:co_owner_middle_name;size:100"`
    CoOwnerLastName    string `gorm:"column:co_owner_last_name;size:100"`
    Email              string `gorm:"column:email;size:100;not null"`
    Mobile             string `gorm:"column:mobile;size:20;not null"`
    AlternateMobile    string `gorm:"column:alternate_mobile;size:20"`
    CommissionRate     int    `gorm:"column:commission_rate"`
		NoOfBranch         int    `gorm:"column:no_of_branch;default:1"`
    IsActive           bool   `gorm:"column:is_active"`
    BusinessName            string `gorm:"column:gym_name;size:50;not null"`
    Address            string `gorm:"column:address;size:255"`
    City               string `gorm:"column:city;size:100"`
    State              string `gorm:"column:state;size:100"`
    PinCode            string `gorm:"column:pin_code;size:20"`
    CreatedBy          int    `gorm:"column:created_by"`
    UpdatedBy          int    `gorm:"column:updated_by"`

    FitCrews []FitCrew `gorm:"foreignKey:AllieID"`   // One GymPartner has many GymBranches
		FitService []FitService  `gorm:"many2many:fit_allie_services;"`
}
