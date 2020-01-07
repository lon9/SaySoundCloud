package controllers

import (
	"strconv"

	"github.com/labstack/echo"
)

type response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

func newResponse(status int, message string, result interface{}) *response {
	return &response{status, message, result}
}

func parseLimitOffset(c echo.Context) (offset, limit int, err error) {
	offsetStr := c.QueryParam("offset")
	if offsetStr == "" {
		offsetStr = "0"
	}
	offset, err = strconv.Atoi(offsetStr)
	if err != nil {
		return
	}
	limitStr := c.QueryParam("limit")
	if limitStr == "" {
		limitStr = "5"
	}
	limit, err = strconv.Atoi(limitStr)
	if limit > 50 {
		limit = 50
	}
	return
}
