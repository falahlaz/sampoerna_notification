package service

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/sholludev/sampoerna_notification/app/notification"
)

type Service interface {
	SendSingleNotification(c echo.Context, request notification.SingleNotifRequestDTO) error
}
