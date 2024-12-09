package models
import (
	"time"
)

type TrainerProfile struct {
    BaseModel
    CrewID          int       `gorm:"column:crew_id"`
    FullName         string    `gorm:"column:full_name;size:50;not null"`
    Email            string    `gorm:"column:email;size:100"`
    Mobile           string    `gorm:"column:mobile;size:20;not null"`
    AlternateMobile  string    `gorm:"column:alternate_mobile;size:20"`
    IsActive         bool      `gorm:"column:is_active;default:true"`
    ExpStartedFrom   time.Time `gorm:"column:exp_started_from;type:date;not null"`
    FitCrew         FitCrew  `gorm:"foreignKey:CrewID"` // Each TrainerProfile belongs to one FitCrew
}
