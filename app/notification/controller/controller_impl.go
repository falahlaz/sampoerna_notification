package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gitlab.com/sholludev/sampoerna_notification/app/notification"
	"gitlab.com/sholludev/sampoerna_notification/app/notification/service"
	"gitlab.com/sholludev/sampoerna_notification/pkg/util/response"
)

type controllerImpl struct {
	Service service.Service
}

func NewController(service service.Service) Controller {
	return &controllerImpl{
		Service: service,
	}
}

// SendSingleNotification implements Controller
func (co *controllerImpl) SendSingleNotification(c echo.Context) error {
	var request notification.SingleNotifRequestDTO
	if err := c.Bind(&request); err != nil {
		return response.ErrorResponse(c, response.BuildError(response.ErrBadRequest, err))
	}

	if err := co.Service.SendSingleNotification(c, request); err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, http.StatusOK, "success send single notification", nil)
}
