package models

import "github.com/google/uuid"

type Author struct {
	Base
	Name        string `gorm:"size:50;not null;index"`
	Nick        string `gorm:"size:50;index"`
	DateOfBirth string `gorm:"size:50;not null;index"`
	DateOfDeath string `gorm:"size:50;index"`
	BookID      uuid.UUID
}

func NewAuthor() Author {
	return Author{}
}
