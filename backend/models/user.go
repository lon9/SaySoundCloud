package models

import (
	"github.com/jinzhu/gorm"
)

// User is struct of user
type User struct {
	gorm.Model
	UID          string        `json:"uid" gorm:"unique"`
	Name         string        `json:"name"`
	Applications []Application `json:"applications"`
	Description  string        `json:"description"`
}

func (u *User) Create() (err error) {
	return
}
