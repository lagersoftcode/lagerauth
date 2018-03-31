package database

import (
	"errors"
	"lagerauth/models/core"

	"github.com/satori/go.uuid"
)

type Application struct {
	db *Database
}

func (d *Database) Application() *Application {
	return &Application{db: d}
}

func (a *Application) GetClientSecretFromApplicationID(applicationID uint) (string, error) {

	var app core.Application
	a.db.First(&app, applicationID)

	if app.SecretKey == "" {
		return "", errors.New("ApplicationId Not Found")
	}

	return app.SecretKey, nil
}

func (a *Application) UserHasAccess(userID uint, clientID uuid.UUID) error {

	applicationID, err := a.GetIDFromClientID(clientID)
	if err != nil {
		return err
	}

	var user core.User
	res := a.db.First(&user, userID).Related(&user.Applications, "Applications")

	exists := false
	for _, app := range user.Applications {
		if app.ID == applicationID {
			exists = true
			break
		}
	}

	if !exists {
		return errors.New("Account has no access to application")
	}

	return res.Error
}

func (a *Application) GetIDFromClientID(clientID uuid.UUID) (uint, error) {

	var application core.Application
	res := a.db.Where(&core.Application{ClientID: clientID}).First(&application)

	if res.Error != nil {
		return 0, errors.New("Cannot find Application")
	}

	return application.ID, nil
}
