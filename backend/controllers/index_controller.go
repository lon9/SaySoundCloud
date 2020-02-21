package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/lon9/SaySoundCloud/backend/config"
)

// IndexController returns server info
type IndexController struct{}

// NewIndexController is constructor for IndexController
func NewIndexController() *IndexController {
	return new(IndexController)
}

func (ic *IndexController) Version(c echo.Context) error {
	conf := config.GetConfig()
	return c.String(http.StatusOK, conf.GetString("server.version"))
}
