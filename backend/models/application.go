package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/lon9/soundboard/backend/database"
)

// Application is struct of application
type Application struct {
	gorm.Model
	Name             string `json:"name" gorm:"unique;not null"`
	Password         string `json:"-"`
	IsPassword       bool   `json:"isPassword"`
	User             *User  `json:"user"`
	UserID           uint   `json:"userId"`
	Description      string `json:"description"`
	AccessToken      string `json:"accessToken"`
	GuestAccessToken string `json:"guestAccessToken"`
}

// Create creates an application
func (a *Application) Create() (err error) {
	db := database.GetDB()
	return db.Create(a).Error
}

// FindByID finds an application by id
func (a *Application) FindByID(id uint) (err error) {
	db := database.GetDB()
	return db.Where("id = ?", id).First(a).Error
}

// Update updates an application
func (a *Application) Update() (err error) {
	db := database.GetDB()
	return db.Save(a).Error
}

// Delete deletes an application
func (a *Application) Delete() (err error) {
	db := database.GetDB()
	return db.Delete(a).Error
}

// Applications is slice of applications
type Applications []Application

// List lists applications
func (as *Applications) List(offset, limit int) (err error) {
	db := database.GetDB()
	return db.Order("name asc").Limit(limit).Offset(offset).Find(as).Error
}

// SearchByName searches applications by name
func (as *Applications) SearchByName(name string, limit, offset int) (err error) {
	db := database.GetDB()
	query := fmt.Sprintf("%%%s%%", name)
	return db.Where("name LIKE ?", query).Order("name asc").Limit(limit).Offset(offset).Find(as).Error
}
