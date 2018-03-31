package core

import (
	"lagerauth/models"

	"github.com/satori/go.uuid"
)

type Application struct {
	models.Base

	Name        string    `json:"name" gorm:"unique_index"`
	Description string    `json:"description"`
	ClientID    uuid.UUID `json:"clientId" gorm:"type:varchar(36)"`
	SecretKey   string    `json:"secretKey"`
	Enabled     bool      `json:"enabled"`

	Users []User `json:"-" gorm:"many2many:users_applications;"`
	Roles []Role `json:"-"`
}
