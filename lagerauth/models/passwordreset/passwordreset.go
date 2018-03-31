package passwordreset

import (
	"lagerauth/models"
	"lagerauth/models/core"

	"github.com/satori/go.uuid"
)

type PasswordResetCode struct {
	models.Base
	Code uuid.UUID `gorm:"type:varchar(36)"`

	UserID uint
	User   core.User
}
