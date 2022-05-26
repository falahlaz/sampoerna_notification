package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	res "gitlab.com/sholludev/sampoerna_notification/pkg/util/response"
)

func NewErrorHandler(err error, c echo.Context) {
	report, ok := err.(*echo.HTTPError)
	if !ok {
		report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	switch report.Code {
	case http.StatusNotFound:
		err = res.BuildCustomError(http.StatusNotFound, "Route not found")
	case http.StatusInternalServerError:
		err = res.BuildError(res.ErrServerError, err)
	default:
		err = res.BuildError(res.ErrServerError, err)
	}

	res.ErrorResponse(c, err)
}
