package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gitlab.com/sholludev/sampoerna_notification/pkg/util/environment"
	"gitlab.com/sholludev/sampoerna_notification/routes/handler"
)

func Init(g *echo.Group) {
	g.GET("", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to "+environment.Get("APP_NAME")+"! version "+environment.Get("APP_VERSION"))
	})

	// Routes
	handler.NewNotificationHandler().Route(g.Group("/notifications"))
}
