package core

import "lagerauth/models"

type User struct {
	models.Base

	Email       string  `json:"email" gorm:"unique_index"`
	Password    string  `json:"-"`
	Name        *string `json:"name"`
	Department  *string `json:"department"`
	LockedOut   bool    `json:"isLockedout"`
	Description *string `json:"description"`
	Extra       *string `json:"-"`

	Applications []Application `json:"-" gorm:"many2many:users_applications;" `
	Roles        []Role        `json:"-" gorm:"many2many:users_roles;" `
}
