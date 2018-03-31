package database

import (
	"errors"
	"lagerauth/models/core"
)

type User struct {
	db *Database
}

func (d *Database) User() *User {
	return &User{db: d}
}

func (u *User) GetUsers() ([]core.User, error) {
	var users []core.User
	res := u.db.Find(&users)
	return users, res.Error
}

func (u *User) EmailExists(email string) error {

	var user core.User
	res := u.db.Where("email = ?", email).First(&user)

	if res.Error != nil || user.Email == "" {
		return errors.New("Email not found")
	}

	return nil
}

func (u *User) GetHashedPasswordFromEmail(email string) ([]byte, error) {

	var user core.User
	res := u.db.Where("email = ? ", email).First(&user)

	if res.Error == nil {
		return []byte(user.Password), nil
	}

	return nil, errors.New("Login is not valid")
}

func (u *User) GetEmailFromID(userID uint) (string, error) {
	var user core.User
	res := u.db.First(&user, userID)
	return user.Email, res.Error
}

func (u *User) GetIDFromEmail(email string) (uint, error) {
	var user core.User
	res := u.db.Where("email = ?", email).First(&user)

	if res.Error != nil {
		return 0, res.Error
	}

	return user.ID, nil
}

func (u *User) ResetPassword(userID uint, password string) error {
	var user core.User
	res := u.db.First(&user, userID)
	if res.Error != nil {
		return res.Error
	}

	user.Password = password
	res = u.db.Save(&user)
	return res.Error
}
