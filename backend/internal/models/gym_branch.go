
package models

type GymBranch struct {
    BaseModel
    PartnerID         int    `gorm:"column:partner_id"`
    ManagerFirstName  string `gorm:"column:manager_first_name;size:100"`
    ManagerMiddleName string `gorm:"column:manager_middle_name;size:100"`
    ManagerLastName   string `gorm:"column:manager_last_name;size:100"`
    Email             string `gorm:"column:email;size:100"`
    Mobile            string `gorm:"column:mobile;size:20"`
    AlternateMobile   string `gorm:"column:alternate_mobile;size:20"`
    IsActive          bool   `gorm:"column:is_active"`
    GymName           string `gorm:"column:gym_name;size:50"`
    Address           string `gorm:"column:address;size:255"`
    City              string `gorm:"column:city;size:100"`
    State             string `gorm:"column:state;size:100"`
    PinCode           string `gorm:"column:pin_code;size:20"`
    Capacity          int    `gorm:"column:capacity"`
    CreatedBy         int    `gorm:"column:created_by"`
    UpdatedBy         int    `gorm:"column:updated_by"`

    GymPartner GymPartner `gorm:"foreignKey:PartnerID"`  // Each GymBranch belongs to one GymPartner
    Customers  []Customer `gorm:"foreignKey:BranchID"`   // Each GymBranch has many Customers
}
