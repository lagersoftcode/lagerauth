package logic

import (
	"errors"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/satori/go.uuid"
)

func GenerateCode(email string, clientID uuid.UUID) (uuid.UUID, error) {
	code := uuid.NewV4()

	userID, _ := db.User().GetIDFromEmail(email)
	err := db.Application().UserHasAccess(userID, clientID)
	if err != nil {
		return uuid.Nil, err
	}

	err = db.OAuth().AddOAuthCode(code, clientID, email)
	if err != nil {
		return uuid.Nil, err
	}

	return code, nil
}

func GenerateToken(code, clientID uuid.UUID, clientSecret string) (string, error) {
	applicationID, err := db.OAuth().GetApplicationIDFromCodeAndClientSecret(code, clientID, clientSecret)
	if err != nil {
		return "", errors.New("Cannot find application on code")
	}

	return generateToken(code, applicationID)
}

func GenerateTokenNoSecret(code, clientID uuid.UUID) (string, error) {

	applicationID, err := db.OAuth().GetApplicationIDFromCode(code, clientID)
	if err != nil {
		return "", errors.New("Cannot find application on code")
	}

	return generateToken(code, applicationID)
}

func ValidateAndGetTokenFromAuthorizationHeader(authorizationHeader string) (uuid.UUID, error) {
	/*
		1) Get jwt part from header
		2) Extract the claims from jwt
		3) Get the inner 'token', get client_secret trough the token and database.
		4) Validate jwt signature (see if it hasnt been tampered)
	*/
	if authorizationHeader == "" {
		return uuid.Nil, errors.New("No Authorization header presennt")
	}

	auth := strings.SplitN(authorizationHeader, " ", 2)
	if len(auth) < 2 || auth[0] != "Bearer" {
		return uuid.Nil, errors.New("Malformed Authorization header")
	}

	token, err := jwt.Parse(auth[1], func(token *jwt.Token) (interface{}, error) {
		innerToken := token.Claims.(jwt.MapClaims)["token"].(string)
		if innerToken == "" {
			return nil, errors.New("Cannot find inner token on Authorization header")
		}

		if token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, errors.New("Authorization Token alg has been tampered")
		}

		secret, err := db.OAuth().GetClientSecretFromToken(uuid.FromStringOrNil(innerToken))
		if err != nil {
			return nil, errors.New("Authorization token has been invalidated")
		}

		return []byte(secret), nil
	})

	if !token.Valid {
		return uuid.Nil, errors.New("Token signature fails")
	}

	return uuid.FromStringOrNil(token.Claims.(jwt.MapClaims)["token"].(string)), err
}

func generateToken(code uuid.UUID, applicationID uint) (string, error) {
	token := uuid.NewV4()

	userID, err := db.OAuth().GetUserIDFromCode(code)
	if err != nil {
		return "", errors.New("User not found on code")
	}

	err = db.OAuth().AddTokenAndBurnCode(code, token, applicationID, userID)
	if err != nil {
		return "", errors.New("Cannot burn code")
	}

	secret, err := db.Application().GetClientSecretFromApplicationID(applicationID)
	if err != nil {
		return "", errors.New("Cannot find secret for application on code")
	}

	email, err := db.User().GetEmailFromID(userID)
	if err != nil {
		return "", errors.New("Cant find email for given userid")
	}

	jwt := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"token": token,
		"email": email,
	})
	jwtString, err := jwt.SignedString([]byte(secret))
	return jwtString, err
}
