package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gitlab.com/sholludev/sampoerna_notification/pkg/util/response"
)

func GetBearerToken(c echo.Context) (string, error) {
	// get auth token
	authToken := c.Request().Header.Get("Authorization")
	if authToken == "" {
		return "", response.BuildCustomError(http.StatusUnauthorized, "Unauthorized")
	}

	return authToken, nil
}
