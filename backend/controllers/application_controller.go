package controllers

import (
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	mymiddleware "github.com/lon9/soundboard/backend/middleware"
	"github.com/lon9/soundboard/backend/models"
)

// ApplicationController is controller for applications
type ApplicationController struct{}

// NewApplicationController is constructor for ApplicationController
func NewApplicationController() *ApplicationController {
	return new(ApplicationController)
}

// Index index applications
func (ac *ApplicationController) Index(c echo.Context) (err error) {
	apps := new(models.Applications)

	offset, limit, err := parseLimitOffset(c)
	if err != nil {
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
		// Seach applications
		if err := apps.SearchByName(query, offset, limit); err != nil {
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
		return c.JSON(
			http.StatusOK,
			newResponse(
				http.StatusOK,
				http.StatusText(http.StatusOK),
				apps,
			),
		)
	}

	if err := apps.List(offset, limit); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			newResponse(
				http.StatusInternalServerError,
				http.StatusText(http.StatusInternalServerError),
				nil,
			),
		)
	}
	return c.JSON(
		http.StatusOK,
		newResponse(
			http.StatusOK,
			http.StatusText(http.StatusOK),
			apps,
		),
	)
}

// Show returns an application
func (ac *ApplicationController) Show(c echo.Context) (err error) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			newResponse(
				http.StatusBadRequest,
				http.StatusText(http.StatusBadRequest),
				nil,
			),
		)
	}

	app := new(models.Application)
	if err := app.FindByID(uint(id)); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			newResponse(
				http.StatusBadRequest,
				http.StatusText(http.StatusBadRequest),
				nil,
			),
		)
	}

	return c.JSON(
		http.StatusOK,
		newResponse(
			http.StatusOK,
			http.StatusText(http.StatusOK),
			app,
		),
	)
}

// Create creates an application
func (ac *ApplicationController) Create(c echo.Context) (err error) {
	app := new(models.Application)
	if err := c.Bind(app); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			newResponse(
				http.StatusBadRequest,
				http.StatusText(http.StatusBadRequest),
				nil,
			),
		)
	}

	idToken := mymiddleware.ExtractClaims(c)
	user := new(models.User)
	if err := user.FindByUID(idToken.UID); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			newResponse(
				http.StatusBadRequest,
				http.StatusText(http.StatusBadRequest),
				nil,
			),
		)
	}
	app.UserID = user.ID
	if err := app.Create(); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			newResponse(
				http.StatusInternalServerError,
				http.StatusText(http.StatusInternalServerError),
				nil,
			),
		)
	}

	return c.JSON(
		http.StatusCreated,
		newResponse(
			http.StatusCreated,
			http.StatusText(http.StatusCreated),
			app,
		),
	)
}

// Update updates a application
func (ac *ApplicationController) Update(c echo.Context) (err error) {
	app := new(models.Application)
	if err = c.Bind(app); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			newResponse(
				http.StatusBadRequest,
				http.StatusText(http.StatusBadRequest),
				nil,
			),
		)
	}
	idToken := mymiddleware.ExtractClaims(c)
	current := new(models.Application)
	if err := current.FindByID(app.ID); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			newResponse(
				http.StatusBadRequest,
				http.StatusText(http.StatusBadRequest),
				nil,
			),
		)
	}

	user := new(models.User)
	if err := user.FindByUID(idToken.UID); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			newResponse(
				http.StatusBadRequest,
				http.StatusText(http.StatusBadRequest),
				nil,
			),
		)
	}

	// Check the user is owner of the app
	if user.ID != current.UserID || user.ID != app.UserID {
		return c.JSON(
			http.StatusBadRequest,
			newResponse(
				http.StatusBadRequest,
				http.StatusText(http.StatusBadRequest),
				nil,
			),
		)
	}

	if err := app.Update(); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			newResponse(
				http.StatusInternalServerError,
				http.StatusText(http.StatusInternalServerError),
				nil,
			),
		)
	}

	return c.JSON(
		http.StatusOK,
		newResponse(
			http.StatusOK,
			http.StatusText(http.StatusOK),
			app,
		),
	)
}

// Destroy destroys an application
func (ac *ApplicationController) Destroy(c echo.Context) (err error) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			newResponse(
				http.StatusBadRequest,
				http.StatusText(http.StatusBadRequest),
				nil,
			),
		)
	}

	app := new(models.Application)
	app.ID = uint(id)
	if err := app.Delete(); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			newResponse(
				http.StatusBadRequest,
				http.StatusText(http.StatusBadRequest),
				nil,
			),
		)
	}

	return c.NoContent(http.StatusNoContent)
}
