package views

import (
	"github.com/jinzhu/gorm"
	"github.com/lon9/soundboard/backend/models"
)

// ApplicationView is view for an application
type ApplicationView struct {
	gorm.Model
	Name        string    `json:"name"`
	IsPassword  bool      `json:"password"`
	User        *UserView `json:"user"`
	UserID      uint      `json:"userId"`
	Description string    `json:"description"`
}

// NewApplicationView is constructor for view of an application
func NewApplicationView(app *models.Application) *ApplicationView {
	view := &ApplicationView{
		Name:        app.Name,
		IsPassword:  app.IsPassword,
		UserID:      app.UserID,
		Description: app.Description,
	}
	if app.User != nil {
		view.User = NewUserView(app.User)
	}
	view.Model = app.Model
	return view
}

// OwnerApplicationView is view for owner of the application
type OwnerApplicationView struct {
	gorm.Model
	Name        string    `json:"name"`
	IsPassword  bool      `json:"isPassword"`
	User        *UserView `json:"user"`
	UserID      uint      `json:"userId"`
	Description string    `json:"description"`
	AccessToken string    `json:"accessToken"`
}

// NewOwnerApplicationView is constructor for view of owner's application
func NewOwnerApplicationView(app *models.Application) *OwnerApplicationView {
	view := &OwnerApplicationView{
		Name:        app.Name,
		IsPassword:  app.IsPassword,
		UserID:      app.UserID,
		Description: app.Description,
		AccessToken: app.AccessToken,
	}
	if app.User != nil {
		view.User = NewUserView(app.User)
	}
	view.Model = app.Model
	return view
}
