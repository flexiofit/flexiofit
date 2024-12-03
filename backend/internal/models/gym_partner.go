
package models

type GymPartner struct {
    BaseModel
    UserID             int    `gorm:"column:user_id"`
    OwnerFirstName     string `gorm:"column:owner_first_name;size:100"`
    OwnerMiddleName    string `gorm:"column:owner_middle_name;size:100"`
    OwnerLastName      string `gorm:"column:owner_last_name;size:100"`
    CoOwnerFirstName   string `gorm:"column:co_owner_first_name;size:100"`
    CoOwnerMiddleName  string `gorm:"column:co_owner_middle_name;size:100"`
    CoOwnerLastName    string `gorm:"column:co_owner_last_name;size:100"`
    Email              string `gorm:"column:email;size:100"`
    Mobile             string `gorm:"column:mobile;size:20"`
    AlternateMobile    string `gorm:"column:alternate_mobile;size:20"`
    CommissionRate     int    `gorm:"column:commission_rate"`
    IsActive           bool   `gorm:"column:is_active"`
    GymName            string `gorm:"column:gym_name;size:255"`
    Address            string `gorm:"column:address;size:255"`
    City               string `gorm:"column:city;size:100"`
    State              string `gorm:"column:state;size:100"`
    PinCode            string `gorm:"column:pin_code;size:20"`
		CreatedBy          int    `gorm:"column:created_by"`
		UpdatedBy          int    `gorm:"column:updated_by"`

		User      User      `gorm:"foreignKey:UserID"`
		// One-to-many relationship with User
		// Users []User `gorm:"foreignkey:GymPartnerID"` // Let GORM handle the foreign key, it assumes "GymPartnerID" as the field in the User model
}
