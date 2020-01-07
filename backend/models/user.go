package models

import (
	"github.com/jinzhu/gorm"
	"github.com/lon9/soundboard/backend/database"
)

// User is struct of user
type User struct {
	gorm.Model
	UID          string        `json:"uid" gorm:"unique"`
	Name         string        `json:"name"`
	Applications []Application `json:"applications"`
	Description  string        `json:"description"`
}

// Create creates a user
func (u *User) Create() (err error) {
	db := database.GetDB()
	return db.Create(u).Error
}

// FindByUID finds a user by uid
func (u *User) FindByUID(uid string) (err error) {
	db := database.GetDB()
	return db.Where("uid = ?", uid).First(u).Error
}

// Update updates user
func (u *User) Update() (err error) {
	db := database.GetDB()
	return db.Save(u).Error
}

// Delete deletes a user
func (u *User) Delete() (err error) {
	db := database.GetDB()
	return db.Delete(u).Error
}

// Users is slice of users
type Users []User
