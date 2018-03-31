package database

import (
	"errors"
	"lagerauth/models/core"
	oauthModels "lagerauth/models/oauth"

	"github.com/satori/go.uuid"
)

type OAuth struct {
	db *Database
}

func (d *Database) OAuth() *OAuth {
	return &OAuth{db: d}
}

func (o *OAuth) AddOAuthCode(code, clientID uuid.UUID, email string) error {

	var application core.Application
	res := o.db.Where(core.Application{ClientID: clientID}).First(&application)
	if res.Error != nil {
		return errors.New("ClientId not found")
	}

	var user core.User
	res = o.db.Where(core.User{Email: email}).First(&user)
	if res.Error != nil {
		return errors.New("User with that email not found")
	}

	oAuthCode := oauthModels.OAuthCode{
		Code:          code,
		ApplicationID: application.ID,
		UserID:        user.ID,
	}

	res = o.db.Create(&oAuthCode)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (o *OAuth) AddTokenAndBurnCode(code, token uuid.UUID, applicationID, userID uint) error {

	oAuthToken := oauthModels.OAuthToken{
		Token:         token,
		ApplicationID: applicationID,
		UserID:        userID,
		Enabled:       true,
	}

	res := o.db.Create(&oAuthToken)
	if res.Error != nil {
		return res.Error
	}

	var oAuthCode oauthModels.OAuthCode
	res = o.db.Where(oauthModels.OAuthCode{Code: code}).First(&oAuthCode)
	if res.Error != nil {
		return res.Error
	}

	oAuthCode.Burned = true
	res = o.db.Save(&oAuthCode)
	return res.Error
}

func (o *OAuth) GetApplicationIDFromCodeAndClientSecret(code, clientID uuid.UUID, clientSecret string) (uint, error) {

	var applicationFromSecret core.Application
	res := o.db.Where(core.Application{ClientID: clientID, SecretKey: clientSecret}).First(&applicationFromSecret)
	if res.Error != nil {
		return 0, errors.New("Cannot find application from client_id")
	}

	var oAuthCode oauthModels.OAuthCode
	res = o.db.Where(oauthModels.OAuthCode{Code: code}).First(&oAuthCode)
	if res.Error != nil {
		return 0, errors.New("Cannot find application_id from code")
	}

	if oAuthCode.ApplicationID != applicationFromSecret.ID {
		return 0, errors.New("client_id doesnt match with the code")
	}

	return applicationFromSecret.ID, nil
}

func (o *OAuth) GetApplicationIDFromCode(code, clientID uuid.UUID) (uint, error) {

	var codeModel oauthModels.OAuthCode
	res := o.db.Where(oauthModels.OAuthCode{Code: code}).First(&codeModel)
	if res.Error != nil {
		return 0, errors.New("Cannot find application from code")
	}

	var application core.Application
	res = o.db.Where(core.Application{ClientID: clientID}).First(&application)
	if res.Error != nil {
		return 0, errors.New("Cannot find application id from client_id")
	}

	if codeModel.ApplicationID != application.ID {
		return 0, errors.New("Application doesnt match code and client_id values")
	}

	return application.ID, nil
}

func (o *OAuth) GetUserIDFromCode(code uuid.UUID) (uint, error) {
	var oAuthCode oauthModels.OAuthCode
	res := o.db.Where(oauthModels.OAuthCode{Burned: false, Code: code}).First(&oAuthCode)
	return oAuthCode.UserID, res.Error
}

func (o *OAuth) GetClientSecretFromToken(token uuid.UUID) (string, error) {
	applicationID, err := o.GetApplicationIDFromToken(token)
	if err != nil {
		return "", err
	}
	return o.db.Application().GetClientSecretFromApplicationID(applicationID)
}

func (o *OAuth) GetApplicationIDFromToken(token uuid.UUID) (uint, error) {
	var oAuthToken oauthModels.OAuthToken
	res := o.db.Where(oauthModels.OAuthToken{Enabled: true, Token: token}).First(&oAuthToken)
	return oAuthToken.ApplicationID, res.Error
}

func (o *OAuth) DisableTokensForUserID(userID uint) {
	o.db.Where(oauthModels.OAuthToken{UserID: userID}).Update(oauthModels.OAuthToken{Enabled: false})
}
