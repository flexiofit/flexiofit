
package models

import (
)

type FitService struct {
	BaseModel
	Name      string `gorm:"column:name;size:100;not null"`
	FitAllies []FitAllie `gorm:"many2many:fit_allie_services;"` // Many-to-many relationship with FitAllie
}
