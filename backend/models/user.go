package models

import (
	"encoding/base64"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/lon9/SaySoundCloud/backend/database"
)

// User is struct of user
type User struct {
	gorm.Model
	UID          string        `json:"uid" gorm:"unique;not null"`
	Name         string        `json:"name" gorm:"unique;not null"`
	Applications []Application `json:"applications"`
	Description  string        `json:"description"`
}

// Create creates a user
func (u *User) Create() (err error) {
	db := database.GetDB()
	return db.Create(u).Error
}

// FindByID finds a user by id
func (u *User) FindByID(id uint) (err error) {
	db := database.GetDB()
	return db.Where("id = ?", id).First(u).Error
}

// FindByUID finds a user by uid
func (u *User) FindByUID(uid string) (err error) {
	db := database.GetDB()
	return db.Where("uid = ?", uid).First(u).Error
}

// FindByName finds a user by name
func (u *User) FindByName(name string) (err error) {
	db := database.GetDB()
	return db.Where("name = ?", name).First(u).Error
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

// GenerateAccessToken generates access token from user
func (u *User) GenerateAccessToken() (token string, err error) {
	id := uuid.New()
	idBytes, err := id.MarshalBinary()
	if err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(idBytes), nil
}

// Users is slice of users
type Users []User
