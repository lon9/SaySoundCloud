package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/lon9/soundboard/backend/database"
)

// Sound is struct of Sound
type Sound struct {
	gorm.Model
	Name string `json:"name" gorm:"unique;not null;index"`
	Path string `json:"path" gorm:"unique;not null"`
}

// FindByName finds a sound by name
func (s *Sound) FindByName(name string) (err error) {
	db := database.GetDB()
	return db.Where("name = ?").First(s).Error
}

// Sounds is slice of sounds
type Sounds []Sound

// List lists sounds
func (ss *Sounds) List(offset, limit int) (err error) {
	db := database.GetDB()
	return db.Order("name asc").Limit(limit).Offset(offset).Find(ss).Error
}

func (ss *Sounds) SearchByName(name string, offset, limit int) (err error) {
	db := database.GetDB()
	query := fmt.Sprintf("%%%s%%", name)
	return db.Where("name LIKE ?", query).Order("name asc").Limit(limit).Offset(offset).Find(ss).Error
}
