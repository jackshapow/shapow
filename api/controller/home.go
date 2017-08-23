package controller

import (
	"github.com/labstack/echo"
	"net/http"
)

func (h *Handler) Home(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
