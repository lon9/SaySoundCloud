package controllers

import (
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/labstack/echo"
	"github.com/lon9/SaySoundCloud/backend/forms"
	mymiddleware "github.com/lon9/SaySoundCloud/backend/middleware"
	"github.com/lon9/SaySoundCloud/backend/models"
	"github.com/lon9/SaySoundCloud/backend/views"
)

// UserController is controller for users
type UserController struct{}

// NewUserController is constructor for UserController
func NewUserController() *UserController {
	return new(UserController)
}

// Me returns a user logged in
func (uc *UserController) Me(c echo.Context) (err error) {
	idToken := mymiddleware.ExtractClaims(c)
	user := new(models.User)
	if err = user.FindByUID(idToken.UID); err != nil {
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
	return c.JSON(
		http.StatusOK,
		newResponse(
			http.StatusOK,
			http.StatusText(http.StatusOK),
			user,
		),
	)
}

// Show returns a user
func (uc *UserController) Show(c echo.Context) (err error) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
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

	user := new(models.User)
	if err := user.FindByID(uint(id)); err != nil {
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
	return c.JSON(
		http.StatusOK,
		newResponse(
			http.StatusOK,
			http.StatusText(http.StatusOK),
			views.NewUserView(user),
		),
	)
}

// Create creates a user
func (uc *UserController) Create(c echo.Context) (err error) {
	idToken := mymiddleware.ExtractClaims(c)
	user := new(models.User)
	user.UID = idToken.UID
	user.Name = uuid.New().String()
	if err = user.Create(); err != nil {
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
	return c.JSON(
		http.StatusCreated,
		newResponse(
			http.StatusCreated,
			http.StatusText(http.StatusCreated),
			views.NewUserView(user),
		),
	)
}

// Update update User
func (uc *UserController) Update(c echo.Context) (err error) {
	userForm := new(forms.UserForm)
	if err = c.Bind(userForm); err != nil {
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

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
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

	idToken := mymiddleware.ExtractClaims(c)
	user, err := userForm.Update(uint(id), idToken)
	if err != nil {
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
	return c.JSON(
		http.StatusOK,
		newResponse(
			http.StatusOK,
			http.StatusText(http.StatusOK),
			views.NewUserView(user),
		),
	)
}

// Destroy destroys a user
func (uc *UserController) Destroy(c echo.Context) (err error) {

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
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

	user := new(models.User)
	if err := user.FindByID(uint(id)); err != nil {
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

	idToken := mymiddleware.ExtractClaims(c)
	// authorization
	if idToken.UID != user.UID {
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

	if err := user.Delete(); err != nil {
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
	return c.NoContent(http.StatusNoContent)
}
