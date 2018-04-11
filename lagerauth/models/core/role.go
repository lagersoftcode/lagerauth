package core

import (
	"lagerauth/models"
)

type Role struct {
	models.Base

	Name        string `json:"name"`
	Description string `json:"description"`
	Extra       string `json:"extra"`

	Permissions []Permission `json:"permissions"`

	Users []User `json:"-" gorm:"many2many:users_roles;"`

	Application   Application `json:"application"`
	ApplicationID uint        `json:"-"`
}
