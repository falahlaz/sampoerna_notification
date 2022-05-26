package service

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/sholludev/sampoerna_notification/app/notification"
	"gitlab.com/sholludev/sampoerna_notification/pkg/firebase"
	"gitlab.com/sholludev/sampoerna_notification/pkg/util/response"
)

type serviceImpl struct {
}

func NewService() Service {
	return &serviceImpl{}
}

// SendSingleNotification implements Service
func (s *serviceImpl) SendSingleNotification(c echo.Context, request notification.SingleNotifRequestDTO) error {
	// validate
	if err := c.Validate(request); err != nil {
		return response.BuildError(response.ErrUnprocessableEntity, err)
	}

	// send notification
	firebaseModel := firebase.SingleNotification{
		FCMToken: request.FCMToken,
		Message: firebase.MessageConstant{
			Action:  "FLUTTER_NOTIFICATION_CLICK",
			Message: request.Message,
			Type:    request.Type,
			Data:    request.Data,
		},
	}

	_, err := firebaseModel.Send(c.Request().Context())
	if err != nil {
		return response.BuildError(response.ErrServerError, err)
	}

	return nil
}
