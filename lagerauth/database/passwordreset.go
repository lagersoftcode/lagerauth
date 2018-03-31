package database

import (
	passwordResetModels "lagerauth/models/passwordreset"

	"github.com/satori/go.uuid"
)

type PasswordReset struct {
	db *Database
}

func (d *Database) PasswordReset() *PasswordReset {
	return &PasswordReset{db: d}
}

func (p *PasswordReset) DeleteCode(code uuid.UUID) {
	var passwordResetCode passwordResetModels.PasswordResetCode
	res := p.db.Where(passwordResetModels.PasswordResetCode{Code: code}).First(&passwordResetCode)
	if res.Error != nil {
		return
	}
	p.db.Delete(&passwordResetCode)
}

func (p *PasswordReset) GetUserIDFromCode(code uuid.UUID) (uint, error) {
	var passwordResetCode passwordResetModels.PasswordResetCode
	res := p.db.Where(passwordResetModels.PasswordResetCode{Code: code}).First(&passwordResetCode)
	return passwordResetCode.UserID, res.Error
}

func (p *PasswordReset) AddCode(userId uint, code uuid.UUID) error {
	p.db.Where(passwordResetModels.PasswordResetCode{UserID: userId}).Delete(passwordResetModels.PasswordResetCode{})

	passwordResetCode := passwordResetModels.PasswordResetCode{
		UserID: userId,
		Code:   code,
	}

	res := p.db.Create(&passwordResetCode)
	return res.Error
}
