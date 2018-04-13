package logic

import (
	"lagerauth/models/core"
	"time"
)

func GetRoles() ([]core.Role, error) {
	var roles []core.Role
	res := db.Preload("Application").Preload("Permissions").Find(&roles)
	return roles, res.Error
}

func GetRole(id uint) (core.Role, error) {
	var role core.Role
	res := db.Preload("Application").Preload("Permissions").Find(&role, id)
	return role, res.Error
}

func DeleteRole(id uint) error {
	res := db.Where(core.Permission{RoleID: id}).Delete(core.Permission{})
	res = db.Delete(core.Role{}, id)
	return res.Error
}

func CreateRole(role core.Role) error {
	role.CreatedAt = time.Now()
	role.UpdatedAt = time.Now()
	res := db.Create(&role)
	return res.Error
}

func UpdateRole(role core.Role, id uint) error {
	var dbRole core.Role
	res := db.Find(&dbRole, id)
	if res.Error != nil {
		return res.Error
	}

	dbRole.Name = role.Name
	dbRole.Description = role.Description
	dbRole.ApplicationID = role.Application.ID
	dbRole.UpdatedAt = time.Now()

	// delete all and re-populate
	res = db.Where(core.Permission{RoleID: id}).Delete(core.Permission{})
	for _, permission := range role.Permissions {
		permission.RoleID = id
		res = db.Create(&permission)
	}

	res = db.Save(&dbRole)
	return res.Error
}
