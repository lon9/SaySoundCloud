package views

import (
	"github.com/lon9/soundboard/backend/models"
)

// SoundView is view for a sound
type SoundView struct {
	Name string `json:"name"`
}

// NewSoundView is constructor for SoundView
func NewSoundView(sound *models.Sound) *SoundView {
	return &SoundView{
		Name: sound.Name,
	}
}
