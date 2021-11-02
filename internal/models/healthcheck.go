package models

type Healthcheck struct {
	Base
	Status bool `gorm:"type:bool"`
}

func NewHealthcheck() *Healthcheck {
	return &Healthcheck{}
}
