package forms

import (
	"errors"

	"firebase.google.com/go/auth"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/lon9/soundboard/backend/config"
	"github.com/lon9/soundboard/backend/models"
	"golang.org/x/crypto/bcrypt"
)

// ApplicationForm is a form for application
type ApplicationForm struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Password    string `json:"password"`
}

// Create creates an application
func (af *ApplicationForm) Create(idToken *auth.Token) (ret *models.Application, err error) {
	app := new(models.Application)
	user := new(models.User)
	if err := user.FindByUID(idToken.UID); err != nil {
		return nil, err
	}
	app.UserID = user.ID
	app.Name = af.Name
	app.Description = af.Description
	if af.Password != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(af.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		app.IsPassword = true
		app.Password = string(hash)
	}

	// generate access token
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.StandardClaims{
		Audience: string(user.ID),
		Subject:  user.UID,
		Id:       uuid.New().String(),
	})
	access, err := token.SignedString([]byte(config.GetConfig().GetString("server.access_token_secret")))
	if err != nil {
		return nil, err
	}

	hashedToken, err := bcrypt.GenerateFromPassword([]byte(access), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	app.AccessToken = string(hashedToken)

	token = jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.StandardClaims{
		Audience: string(user.ID),
		Subject:  user.UID,
		Id:       uuid.New().String(),
	})
	guest, err := token.SignedString([]byte(config.GetConfig().GetString("server.access_token_secret")))
	app.GuestAccessToken = guest

	err = app.Create()
	return app, err
}

// Update updates an application
func (af *ApplicationForm) Update(id uint, idToken *auth.Token) (ret *models.Application, err error) {
	user := new(models.User)
	if err := user.FindByUID(idToken.UID); err != nil {
		return nil, err
	}
	app := new(models.Application)
	if err := app.FindByID(id); err != nil {
		return nil, err
	}

	// owner authorization
	if app.UserID != user.ID {
		return nil, errors.New("Unauthorized")
	}
	app.Name = af.Name
	app.Description = af.Description
	if af.Password != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(af.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		app.IsPassword = true
		app.Password = string(hash)
	} else {
		app.IsPassword = false
	}
	err = app.Update()
	return app, err
}

// WSAuthForm is form for authentication of websocket
type WSAuthForm struct {
	Password string `json:"password"`
}

// Auth authenticate websocket connection
func (wsf *WSAuthForm) Auth(id uint) (token string, err error) {
	app := new(models.Application)
	if err = app.FindByID(id); err != nil {
		return
	}
	if err = bcrypt.CompareHashAndPassword([]byte(app.Password), []byte(wsf.Password)); err != nil {
		return
	}
	return app.GuestAccessToken, nil
}

// CmdForm is form for sound command
type CmdForm struct {
	Name        string `json:"name"`
	AccessToken string `json:"accessToken"`
}

// Auth authenticate cmd request
func (cf *CmdForm) Auth(id uint) (err error) {
	app := new(models.Application)
	if err = app.FindByID(uint(id)); err != nil {
		return
	}
	return bcrypt.CompareHashAndPassword([]byte(app.AccessToken), []byte(cf.AccessToken))
}
