package core

import (
	"lagerauth/models"
)

type Role struct {
	models.Base

	Name        string
	Description string
	Extra       string

	Permissions []Permission `json:"-"`

	Users []User `json:"-" gorm:"many2many:users_roles;"`

	Application   Application `json:"-"`
	ApplicationID uint
}
