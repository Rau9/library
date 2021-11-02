package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/plugin/soft_delete"
)

// Base is the skeleton of each model using UUID as primary key
// More info: https://gorm.io/docs/models.html
type Base struct {
	ID        uuid.UUID `gorm:"primarykey;type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt soft_delete.DeletedAt `gorm:"index"`
}

// DbModels registers all database models
var DbModels = map[string]interface{}{
	// append here all new models
	"healthcheck": Healthcheck{},
	"user":        User{},
	"book":        Book{},
	"author":      Author{},
}
