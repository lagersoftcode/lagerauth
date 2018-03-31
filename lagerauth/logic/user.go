package logic

import "lagerauth/models/core"

func GetUsers() []core.User {
	var users []core.User
	db.Find(&users)
	return users
}
