package oauth

import (
	"lagerauth/models"
	"lagerauth/models/core"

	"github.com/satori/go.uuid"
)

type OAuthToken struct {
	models.Base

	Token uuid.UUID `gorm:"type:varchar(36)"`

	ApplicationID uint
	Application   core.Application

	UserID uint
	User   core.User

	Enabled bool
}
