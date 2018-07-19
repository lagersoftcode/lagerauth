package database

import (
	"crypto/rand"
	"encoding/base64"
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
}

func seed(db *gorm.DB) {

	secretKey, _ := generateRandomString(64)
	applicationClientId := uuid.FromStringOrNil("00000000-0000-0000-0000-000000000001") // using 1 so its not uuid.Nil
	mangerApplication := &coreModels.Application{
		Name:        "lagerauth",
		Description: "lagerauth manager application",
		ClientID:    applicationClientId,
		SecretKey:   secretKey,
		Enabled:     true,
	}

	// Create application if it doesnt exists
	db.Where(coreModels.Application{ClientID: applicationClientId}).FirstOrCreate(mangerApplication)
}

func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

func generateRandomString(s int) (string, error) {
	b, err := generateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
}
