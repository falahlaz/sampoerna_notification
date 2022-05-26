package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gitlab.com/sholludev/sampoerna_notification/pkg/util/environment"
	"gitlab.com/sholludev/sampoerna_notification/pkg/util/response"
)

func Authorization(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authToken := c.Request().Header.Get("Authorization")
		if authToken == "" {
			return response.ErrorResponse(c, response.BuildError(response.ErrUnauthorized, nil))
		}

		// check token
		host := environment.Get("API_GATEWAY_HOST")
		client := &http.Client{}
		req, err := http.NewRequest("POST", host+"/auth/validate-token", nil)
		if err != nil {
			return response.ErrorResponse(c, response.BuildError(response.ErrServerError, err))
		}

		req.Header.Set("Authorization", authToken)
		resp, err := client.Do(req)
		if err != nil {
			return response.ErrorResponse(c, response.BuildError(response.ErrServerError, err))
		}

		if resp.StatusCode != http.StatusOK {
			return response.ErrorResponse(c, response.BuildError(response.ErrUnauthorized, err))
		}

		return next(c)
	}
}
