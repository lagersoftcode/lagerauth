package core

import (
	"lagerauth/models"
)

type Permission struct {
	models.Base

	Method     string
	Controller string
	Action     string
	Extra      string

	Role   Role `json:"-"`
	RoleID uint
}
