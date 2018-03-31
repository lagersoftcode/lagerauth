package logic

import (
	coreModels "lagerauth/models/core"
	oauthModels "lagerauth/models/oauth"

	"github.com/satori/go.uuid"
)

func GetApplicationPermissionsFromToken(token uuid.UUID) ([]coreModels.Permission, error) {
	var tokenModel oauthModels.OAuthToken
	res := db.First(&tokenModel, oauthModels.OAuthToken{Token: token})
	if res.Error != nil {
		return nil, res.Error
	}

	var user coreModels.User
	res = db.First(&user, tokenModel.UserID).Related(&user.Roles, "Roles")
	if res.Error != nil {
		return nil, res.Error
	}

	var application coreModels.Application
	res = db.First(&application, tokenModel.ApplicationID).Related(&application.Roles, "Roles")

	var permissions []coreModels.Permission

	for _, userRole := range user.Roles {
		for _, role := range application.Roles {

			// If user has role add permissions:
			if role.ID == userRole.ID {
				var rolePermissions []coreModels.Permission
				db.Find(&rolePermissions, coreModels.Permission{RoleID: role.ID})

				for _, p := range rolePermissions {
					permissions = append(permissions, p)
				}
			}

		}
	}

	return permissions, nil
}

func DisableToken(token uuid.UUID) error {
	var oAuthToken oauthModels.OAuthToken
	res := db.First(&oAuthToken, oauthModels.OAuthToken{Token: token})
	if res.Error != nil {
		return res.Error
	}

	oAuthToken.Enabled = false
	db.Save(&oAuthToken)
	return nil
}

func GetUserFromToken(token uuid.UUID) (*coreModels.User, error) {
	var tokenModel oauthModels.OAuthToken
	res := db.First(&tokenModel, oauthModels.OAuthToken{Token: token})
	if res.Error != nil {
		return nil, res.Error
	}

	var user coreModels.User
	res = db.Find(&user, tokenModel.UserID)
	if res.Error != nil {
		return nil, res.Error
	}

	return &user, nil
}
