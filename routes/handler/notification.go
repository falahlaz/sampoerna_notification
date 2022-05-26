package handler

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/sholludev/sampoerna_notification/app/notification/controller"
	"gitlab.com/sholludev/sampoerna_notification/app/notification/service"
)

type handlerNotification struct {
	Controller controller.Controller
}

func NewNotificationHandler() *handlerNotification {
	ns := service.NewService()

	return &handlerNotification{
		Controller: controller.NewController(ns),
	}
}

func (h *handlerNotification) Route(g *echo.Group) {
	g.POST("/single/send", h.Controller.SendSingleNotification)
}
