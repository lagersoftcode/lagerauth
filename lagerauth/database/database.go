package database

import (
	"fmt"
	"log"

	"github.com/satori/go.uuid"

	"lagerauth/config"
	coreModels "lagerauth/models/core"
	oauthModels "lagerauth/models/oauth"
	passwordResetModels "lagerauth/models/passwordreset"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Database struct {
	*gorm.DB
}

func New(conf *config.Config) *Database {

	connStr := fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true&loc=UTC", conf.DBConfig.User, conf.DBConfig.Pass, conf.DBConfig.Host, conf.DBConfig.Port, conf.DBConfig.Database)

	//Gorm Configuration
	db, err := gorm.Open("mysql", connStr)
	if err != nil {
		log.Fatal(err)
	}

	migrate(db)
	seed(db)

	return &Database{DB: db}
}

func migrate(db *gorm.DB) {

	// Add gorm models
	db.AutoMigrate(
		&coreModels.User{},
		&coreModels.Application{},
		&coreModels.Role{},
		&coreModels.Permission{},
		&oauthModels.OAuthCode{},
		&oauthModels.OAuthToken{},
		&passwordResetModels.PasswordResetCode{},
	)

	// Create Log table if not exists:
	db.Exec(`create table if not exists logs (
		ID int not null auto_increment primary key,
		Level varchar(10) not null,
		Message varchar(250) not null,
		CreatedAt timestamp(4) default current_timestamp(6))
	`)
}

func seed(db *gorm.DB) {

	applicationClientId := uuid.FromStringOrNil("00000000-0000-0000-0000-000000000001") // using 1 so its not uuid.Nil
	mangerApplication := &coreModels.Application{
		Name:        "lagerauth",
		Description: "lagerauth manager application",
		ClientID:    applicationClientId,
		SecretKey:   "MZ4dCvVMCzTjgjcrAeeja336jqXtwNrSSX4mRHMwFSdrycF5",
		Enabled:     true,
	}

	// Create application if it doesnt exists
	db.Where(coreModels.Application{ClientID: applicationClientId}).FirstOrCreate(mangerApplication)
}
