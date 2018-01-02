package controller

import (
	"github.com/jackshapow/shapow/api/model"
	"github.com/labstack/echo"
	"net/http"
)

func (h *Handler) Home(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func (h *Handler) HomeReset(c echo.Context) error {
	// Probably need to do this a different way lulz
	model.ResetDB(*h.DB, h.NodeSettings)

	return c.Redirect(http.StatusFound, "/")
}
