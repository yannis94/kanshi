package handlers

import (
	"github.com/labstack/echo/v4"
)

type NetworkHandler struct{}

func NewNetworkHandler() *NetworkHandler {
	return &NetworkHandler{}
}

func (h *NetworkHandler) GetInfo(c echo.Context) error {
	ipaddr := c.Request().RemoteAddr

	return c.JSON(200, map[string]string{"ip-address": ipaddr})
}
