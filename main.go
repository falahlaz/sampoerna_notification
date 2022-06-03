package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	amiddleware "gitlab.com/sholludev/sampoerna_notification/middleware"
	"gitlab.com/sholludev/sampoerna_notification/pkg/database"
	"gitlab.com/sholludev/sampoerna_notification/pkg/firebase"
	"gitlab.com/sholludev/sampoerna_notification/pkg/log"
	"gitlab.com/sholludev/sampoerna_notification/pkg/util"
	"gitlab.com/sholludev/sampoerna_notification/pkg/util/environment"
	"gitlab.com/sholludev/sampoerna_notification/routes"
)

func main() {
	e := echo.New()

	// Init
	environment.Init()
	database.Init("mysql")
	log.Init()
	firebase.Init(e.NewContext(&http.Request{}, nil).Request().Context())

	// Middleware
	e.Use(
		middleware.Logger(),
		middleware.Recover(),
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		}),
		middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: fmt.Sprintf("\n%s | ${host} | ${time_custom} | ${status} | ${latency_human} | ${remote_ip} | ${method} | ${uri} ",
				environment.Get("APP_NAME"),
			),
			CustomTimeFormat: "2006/01/02 15:04:05",
			Output:           os.Stdout,
		}),
	)
	e.HTTPErrorHandler = amiddleware.NewErrorHandler
	e.Validator = &util.CustomValidation{Validator: validator.New()}

	// Migration
	database.Migrate()

	// Route
	routes.Init(e.Group("/api/v1"))

	e.Logger.Fatal(e.Start(":" + environment.Get("APP_PORT")))
}
