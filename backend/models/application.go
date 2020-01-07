package models

import (
	"github.com/jinzhu/gorm"
)

// Application is struct of application
type Application struct {
	gorm.Model
	Name        string `json:"name"`
	Password    string `json:"-"`
	IsPassword  bool   `json:"isPassword"`
	User        *User  `json:"user"`
	UserID      uint   `json:"userId"`
	Description string `json:"description"`
	AccessToken string `json:"accessToken"`
}
