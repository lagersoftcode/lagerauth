package logic

import (
	"errors"

	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

func LoginValid(email, password string) error {

	hashedPasswd, err := db.User().GetHashedPasswordFromEmail(email)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword(hashedPasswd, []byte(password))
	if err == nil {
		return nil
	}

	logger.Error("Login failed for %s", email)
	return errors.New("Login Failed")
}

func EmailExists(email string) error {
	err := db.User().EmailExists(email)
	return err
}

func AddPasswordResetCode(email string) uuid.UUID {
	passwordResetCode := uuid.NewV4()
	userId, _ := db.User().GetIDFromEmail(email)
	db.PasswordReset().AddCode(userId, passwordResetCode)
	return passwordResetCode
}

func SendPasswordResetEmail(email string, code uuid.UUID) {
	err := emailSender.SendResetCode(email, code)
	if err != nil {
		logger.Error(err.Error())
	}
}

func ResetPasswordCodeIsValid(code uuid.UUID) bool {
	_, err := db.PasswordReset().GetUserIDFromCode(code)
	return err == nil
}

func ResetPasswordFromCode(code uuid.UUID, password string) error {
	userID, _ := db.PasswordReset().GetUserIDFromCode(code)
	hashedPass, _ := bcrypt.GenerateFromPassword([]byte(password), 0)
	strPass := string(hashedPass)

	db.PasswordReset().DeleteCode(code)
	db.OAuth().DisableTokensForUserID(userID)
	return db.User().ResetPassword(userID, strPass)
}
