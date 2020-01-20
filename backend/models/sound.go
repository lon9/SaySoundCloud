package models

import (
	"github.com/jinzhu/gorm"
)

// Sound is struct of Sound
type Sound struct {
	gorm.Model
	Name string `json:"name" gorm:"unique;not null;index"`
	Path string `json:"path" gorm:"unique;not null"`
}

// Sounds is slice of sounds
type Sounds []Sound
