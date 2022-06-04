package service

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/sholludev/sampoerna_notification/app/notification"
)

type Service interface {
	SendSingleNotification(c echo.Context, request notification.SingleNotifRequestDTO) error
	GetByIDUser(c echo.Context, IDUser uint) ([]notification.NotifikasiResponseDTO, error)
	UpdateToReadAll(c echo.Context, IDUser uint) ([]notification.NotifikasiResponseDTO, error)
	UpdateToReadSingle(c echo.Context, ID uint) (notification.NotifikasiResponseDTO, error)
}
