package logic

import "lagerauth/models/core"

func GetApplications() ([]core.Application, error) {
	var applications []core.Application
	res := db.Find(&applications)
	return applications, res.Error
}

func GetApplication(id uint) (core.Application, error) {
	var application core.Application
	res := db.Find(&application, id)
	return application, res.Error
}

func DeleteApplication(id uint) error {
	res := db.Delete(core.Application{}, id)
	return res.Error
}

func CreateApplication(application core.Application) error {
	res := db.Create(&application)
	return res.Error
}

func UpdateApplication(application core.Application, id uint) error {
	var dbApp core.Application
	res := db.Find(&dbApp, id)
	if res.Error != nil {
		return res.Error
	}

	dbApp.Name = application.Name
	dbApp.ClientID = application.ClientID
	dbApp.SecretKey = application.SecretKey
	dbApp.Description = application.Description
	dbApp.Enabled = application.Enabled

	res = db.Save(&dbApp)
	return res.Error
}