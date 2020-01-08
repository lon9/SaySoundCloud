package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/lon9/soundboard/backend/forms"
	mymiddleware "github.com/lon9/soundboard/backend/middleware"
	"github.com/lon9/soundboard/backend/models"
	"github.com/lon9/soundboard/backend/views"
)

// UserController is controller for users
type UserController struct{}

// NewUserController is constructor for UserController
func NewUserController() *UserController {
	return new(UserController)
}

// Show returns a user
func (uc *UserController) Show(c echo.Context) (err error) {
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

	user := new(models.User)
	if err := user.FindByID(uint(id)); err != nil {
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

// Update update User
func (uc *UserController) Update(c echo.Context) (err error) {
	userForm := new(forms.UserForm)
	if err = c.Bind(userForm); err != nil {
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
