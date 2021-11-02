package storage

import (
	"github.com/Rau9/library/internal/models"
	"gorm.io/gorm"
)

func RunMigrations(dbClient *gorm.DB) error {
	for _, model := range models.DbModels {
		if err := dbClient.AutoMigrate(model); err != nil {
			return err
		}
	}
	// TODO: seed the database with default users
	// NOTE: this would not be production code and should be deleted
	if _, err := models.CreateUserIfNotExists(dbClient, "admin", "admin"); err != nil {
		return err
	}
	// TODO: seed books categories
	return nil
}
