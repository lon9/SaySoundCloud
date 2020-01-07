package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

// HealthController controller for health request
type HealthController struct{}

// NewHealthController is constructer for HealthController
func NewHealthController() *HealthController {
	return new(HealthController)
}

// Index is index route for health
func (hc *HealthController) Index(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

// Secret is example secret route
func (hc *HealthController) Secret(c echo.Context) error {
	return c.String(http.StatusOK, "This is secret page")
}
