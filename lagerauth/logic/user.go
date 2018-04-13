package logic

import (
	"lagerauth/models/core"
	"time"
)

func GetUsers() []core.User {
	var users []core.User
	db.Find(&users)
	return users
}

func GetUser(id uint) (core.User, error) {
	var user core.User
	res := db.Preload("Applications").Preload("Roles").Preload("Roles.Application").Find(&user, id)
	return user, res.Error
}

func DeleteUser(id uint) error {
	var user core.User
	res := db.Find(&user, id)
	if res.Error != nil {
		return res.Error
	}

	// remove associations
	db.Model(&user).Association("Applications").Clear()
	db.Model(&user).Association("Roles").Clear()

	res = db.Delete(&user)
	return res.Error
}

func CreateUser(user core.User) error {
	user.CreatedAt = time.Now()
	res := db.Create(&user)
	return res.Error
}

func UpdateUser(user core.User, id uint) error {
	var dbUser core.User
	res := db.Find(&dbUser, id)
	if res.Error != nil {
		return res.Error
	}

	dbUser.Email = user.Email
	dbUser.Name = user.Name
	dbUser.Department = user.Department
	dbUser.Enabled = user.Enabled
	dbUser.Description = user.Description
	dbUser.UpdatedAt = time.Now()

	// delete and re-populate associations
	db.Model(&dbUser).Association("Applications").Replace(user.Applications)
	db.Model(&dbUser).Association("Roles").Replace(user.Roles)

	res = db.Save(&dbUser)
	return res.Error
}
