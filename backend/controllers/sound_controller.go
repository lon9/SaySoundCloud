package controllers

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/lon9/SaySoundCloud/backend/models"
)

// SoundController is controller for sounds
type SoundController struct{}

// NewSoundController is constructor for SoundController
func NewSoundController() *SoundController {
	return new(SoundController)
}

// Index index sounds
func (sc *SoundController) Index(c echo.Context) (err error) {
	var sounds models.Sounds
	offset, limit, err := parseLimitOffset(c)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(
			http.StatusBadRequest,
			newResponse(
				http.StatusBadRequest,
				http.StatusText(http.StatusBadRequest),
				nil,
			),
		)
	}
	query := c.QueryParam("q")
	if query != "" {
		// Seach sounds
		if err := sounds.SearchByName(query, offset, limit); err != nil {
			c.Logger().Error(err)
			if gorm.IsRecordNotFoundError(err) {
				return c.JSON(
					http.StatusNotFound,
					newResponse(
						http.StatusNotFound,
						http.StatusText(http.StatusNotFound),
						nil,
					),
				)
			}
			return c.JSON(
				http.StatusInternalServerError,
				newResponse(
					http.StatusInternalServerError,
					http.StatusText(http.StatusInternalServerError),
					nil,
				),
			)
		}
	} else {
		if err := sounds.List(offset, limit); err != nil {
			c.Logger().Error(err)
			return c.JSON(
				http.StatusInternalServerError,
				newResponse(
					http.StatusInternalServerError,
					http.StatusText(http.StatusInternalServerError),
					nil,
				),
			)
		}
	}
}
