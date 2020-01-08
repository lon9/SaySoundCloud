package controllers

import (
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/lon9/soundboard/backend/forms"
	mymiddleware "github.com/lon9/soundboard/backend/middleware"
	"github.com/lon9/soundboard/backend/models"
	"github.com/lon9/soundboard/backend/views"
	"github.com/lon9/wsrooms"
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
	} else {
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
	}

	ret := make([]*views.ApplicationView, len(*apps))
	for i, app := range *apps {
		ret[i] = views.NewApplicationView(&app)
	}

	return c.JSON(
		http.StatusOK,
		newResponse(
			http.StatusOK,
			http.StatusText(http.StatusOK),
			ret,
		),
	)
}

// Show returns an application
func (ac *ApplicationController) Show(c echo.Context) (err error) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
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
	idToken := mymiddleware.ExtractClaims(c)
	user := new(models.User)
	if user.FindByUID(idToken.UID); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			newResponse(
				http.StatusBadRequest,
				http.StatusText(http.StatusBadRequest),
				nil,
			),
		)
	}

	if user.ID == app.UserID {
		return c.JSON(
			http.StatusOK,
			newResponse(
				http.StatusOK,
				http.StatusText(http.StatusOK),
				views.NewOwnerApplicationView(app),
			),
		)
	}

	return c.JSON(
		http.StatusOK,
		newResponse(
			http.StatusOK,
			http.StatusText(http.StatusOK),
			views.NewApplicationView(app),
		),
	)
}

// Create creates an application
func (ac *ApplicationController) Create(c echo.Context) (err error) {
	appForm := new(forms.ApplicationForm)
	if err := c.Bind(appForm); err != nil {
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
	app, err := appForm.Create(idToken)
	if err != nil {
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
			views.NewOwnerApplicationView(app),
		),
	)
}

// Update updates an application
func (ac *ApplicationController) Update(c echo.Context) (err error) {
	appForm := new(forms.ApplicationForm)
	if err = c.Bind(appForm); err != nil {
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
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
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

	app, err := appForm.Update(uint(id), idToken)
	if err != nil {
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
			views.NewOwnerApplicationView(app),
		),
	)
}

// Destroy destroys an application
func (ac *ApplicationController) Destroy(c echo.Context) (err error) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
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

	// authorization
	if app.UserID != user.ID {
		return c.JSON(
			http.StatusBadRequest,
			newResponse(
				http.StatusBadRequest,
				http.StatusText(http.StatusBadRequest),
				nil,
			),
		)
	}

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

// WSAuth authenticate websocket connection
func (ac *ApplicationController) WSAuth(c echo.Context) (err error) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
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
	form := new(forms.WSAuthForm)
	if err := c.Bind(form); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			newResponse(
				http.StatusBadRequest,
				http.StatusText(http.StatusBadRequest),
				nil,
			),
		)
	}
	token, err := form.Auth(uint(id))
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
	return c.JSON(
		http.StatusOK,
		newResponse(
			http.StatusOK,
			http.StatusText(http.StatusOK),
			token,
		),
	)
}

//WS websocket controller
func (ac *ApplicationController) WS(c echo.Context) (err error) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
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

	// Authorizaion
	token := c.QueryParam("token")
	if app.GuestAccessToken != token {
		return c.JSON(
			http.StatusBadRequest,
			newResponse(
				http.StatusBadRequest,
				http.StatusText(http.StatusBadRequest),
				nil,
			),
		)
	}
	conn := wsrooms.NewConnection(c.Response(), c.Request(), nil)
	defer conn.Leave(string(app.ID))
	if c != nil {
		go conn.WritePump()
		conn.Join(string(app.ID))
		go conn.ReadPump()
	}
	return c.NoContent(http.StatusNoContent)
}
