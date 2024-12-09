
package models

type FitCrew struct {
    BaseModel
    AllieID         int    `gorm:"column:allie_id"`
    ManagerFirstName  string `gorm:"column:manager_first_name;size:100;not null"`
    ManagerMiddleName string `gorm:"column:manager_middle_name;size:100"`
    ManagerLastName   string `gorm:"column:manager_last_name;size:100"`
    Email             string `gorm:"column:email;size:100;not null"`
    Mobile            string `gorm:"column:mobile;size:20;not null"`
    AlternateMobile   string `gorm:"column:alternate_mobile;size:20"`
    IsActive          bool   `gorm:"column:is_active; default:true"`
    GymName           string `gorm:"column:gym_name;size:50"`
    Address           string `gorm:"column:address;size:255"`
    City              string `gorm:"column:city;size:100"`
		Lat               string `gorm:"column:lat;size:30"`
		Long              string `gorm:"column:long;size:30"`
    State             string `gorm:"column:state;size:100"`
    PinCode           string `gorm:"column:pin_code;size:20"`
    Capacity          int    `gorm:"column:capacity"`
    CreatedBy         int    `gorm:"column:created_by"`
    UpdatedBy         int    `gorm:"column:updated_by"`

    FitAllie FitAllie `gorm:"foreignKey:AllieID"`  // Each GymBranch belongs to one GymPartner
    Customers  []Customer `gorm:"foreignKey:CrewID"`   // Each GymBranch has many Customers
		TrainerProfiles []TrainerProfile `gorm:"foreignKey:CrewID"`
}
