package core

import (
	"lagerauth/models"
)

type Permission struct {
	models.Base

	Method     string `json:"method"`
	Controller string `json:"controller"`
	Action     string `json:"action"`
	Extra      string `json:"-"`

	Role   Role `json:"-"`
	RoleID uint `json:"-"`
}
