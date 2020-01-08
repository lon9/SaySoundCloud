package forms

import (
	"errors"

	"firebase.google.com/go/auth"
	"github.com/jinzhu/gorm"
	"github.com/lon9/soundboard/backend/models"
)

// UserForm is form for user
type UserForm struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Update updates a user with UserForm
func (uf *UserForm) Update(id uint, idToken *auth.Token) (ret *models.User, err error) {
	user := new(models.User)
	err = user.FindByID(id)
	if err != nil {
		return nil, err
	}

	// authorization
	if idToken.UID != user.UID {
		return nil, errors.New("Unauthorized")
	}
	user.Name = uf.Name
	user.Description = uf.Description
	err = user.Update()
	return user, err
}
