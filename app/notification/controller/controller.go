package controller

import "github.com/labstack/echo/v4"

type Controller interface {
	SendSingleNotification(c echo.Context) error
}
