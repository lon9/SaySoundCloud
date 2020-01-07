package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	mymiddleware "github.com/lon9/soundboard/backend/middleware"
	"github.com/lon9/soundboard/backend/models"
)

// UserController is controller for users
type UserController struct{}

// NewUserController is constructor for UserController
func NewUserController() *UserController {
	return new(UserController)
}

// Show returns a user
func (uc *UserController) Show(c echo.Context) (err error) {
	uid := c.Param("uid")
	user := new(models.User)
	if err := user.FindByUID(uid); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			newResponse(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), nil),
		)
	}
	return c.JSON(
		http.StatusOK,
		newResponse(http.StatusOK, http.StatusText(http.StatusOK), user),
	)
}

// Update update User
func (uc *UserController) Update(c echo.Context) (err error) {
	user := new(models.User)
	if err = c.Bind(user); err != nil {
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
	if idToken.UID != user.UID {
		return c.JSON(
			http.StatusBadRequest,
			newResponse(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), nil),
		)
	}
	if err := user.Update(); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			newResponse(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), nil),
		)
	}
	return c.JSON(
		http.StatusOK,
		newResponse(http.StatusOK, http.StatusText(http.StatusOK), user),
	)
}

// Destroy destroys a user
func (uc *UserController) Destroy(c echo.Context) (err error) {
	uid := c.Param("uid")

	idToken := mymiddleware.ExtractClaims(c)
	if idToken.UID != uid {
		return c.JSON(
			http.StatusBadRequest,
			newResponse(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), nil),
		)
	}

	user := new(models.User)
	user.UID = uid
	if err := user.Delete(); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			newResponse(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), nil),
		)
	}
	return c.NoContent(http.StatusNoContent)
}
