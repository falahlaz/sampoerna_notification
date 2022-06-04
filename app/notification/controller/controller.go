package controller

import "github.com/labstack/echo/v4"

type Controller interface {
	SendSingleNotification(c echo.Context) error
	GetByIDUser(c echo.Context) error
	UpdateToReadAll(c echo.Context) error
	UpdateToReadSingle(c echo.Context) error
}
