package models

import (
	"github.com/jinzhu/gorm"
)

// Sound is struct of Sound
type Sound struct {
	gorm.Model
	CmdName  string `json:"cmdName" gorm:"unique;not null"`
	FileName string `json:"fileName" gorm:"unique;not null"`
}

// Sounds is slice of sounds
type Sounds []Sound
