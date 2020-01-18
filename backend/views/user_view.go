package views

import (
	"github.com/jinzhu/gorm"
	"github.com/lon9/soundboard/backend/models"
)

// UserView is view for user
type UserView struct {
	gorm.Model
	UID          string             `json:"uid"`
	Name         string             `json:"name"`
	Applications []*ApplicationView `json:"applications"`
	Description  string             `json:"description"`
}

// NewUserView is constructor for view of users
func NewUserView(user *models.User) *UserView {
	appViews := make([]*ApplicationView, len(user.Applications))
	for i, app := range user.Applications {
		appViews[i] = NewApplicationView(&app)
	}
	view := &UserView{
		UID:          user.UID,
		Name:         user.Name,
		Applications: appViews,
		Description:  user.Description,
	}
	view.Model = user.Model
	return view
}
