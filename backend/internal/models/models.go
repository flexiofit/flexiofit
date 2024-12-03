// internal/models/models.go

package models

import (
	"fmt"
	"gorm.io/gorm"
)

// RegisterModels contains all models that need GORM auto-migration
var RegisterModels = []interface{}{
	&User{},
	&GymPartner{},
	// Add new models here
	// &Role{},
	// &Permission{},
}

// AutoMigrateDB handles the auto-migration of all registered models
func AutoMigrateDB(db *gorm.DB) error {
	for _, model := range RegisterModels {
		if err := db.AutoMigrate(model); err != nil {
			return fmt.Errorf("failed to migrate %T: %w", model, err)
		}
	}
	return nil
}

// GetRegisteredModels returns the names of all registered models
func GetRegisteredModels(db *gorm.DB) []string {
	modelNames := make([]string, 0, len(RegisterModels))
	for _, model := range RegisterModels {
		modelNames = append(modelNames, db.Model(model).Name())
	}
	return modelNames
}
