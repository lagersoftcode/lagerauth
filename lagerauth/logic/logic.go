package logic

import (
	"lagerauth/database"
	"lagerauth/email"

	l "github.com/ozkar99/logger"
)

/* Database */
var db *database.Database

func SetDB(d *database.Database) {
	db = d
}

func DB() *database.Database {
	return db
}

/* JWT Key Helpers */
var jwtKey []byte

func SetJWTSecret(s string) {
	jwtKey = []byte(s)
}
func JWTSecretKey() []byte {
	return jwtKey
}

/* Logger Helper */
var logger l.Logger

func SetLogger(l l.Logger) {
	logger = l
}

func Logger() l.Logger {
	return logger
}

/* Email Sender */
var emailSender *email.EmailSender

func SetEmailSender(e *email.EmailSender) {
	emailSender = e
}

func EmailSender() *email.EmailSender {
	return emailSender
}
