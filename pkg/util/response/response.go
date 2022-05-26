package response

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gitlab.com/sholludev/sampoerna_notification/pkg/log"
	"gitlab.com/sholludev/sampoerna_notification/pkg/util/environment"
)

type Response struct {
	Success bool        `json:"success" default:"true"`
	Message string      `json:"message" default:"success"`
	Data    interface{} `json:"data"`
}

type ErrorConstant struct {
	Response     Response
	Code         int
	ErrorMessage error
}

func (r *ErrorConstant) Error() string {
	return fmt.Sprintf("error code %d", r.Code)
}

func (r *ErrorConstant) Builder() *ErrorConstant {
	return r
}

const (
	E_DUPLICATE            = "Duplicate data entry"
	E_NOT_FOUND            = "Data not found"
	E_UNPROCESSABLE_ENTITY = "Invalid parameters or payload"
	E_UNAUTHORIZED         = "Unauthorized access"
	E_BAD_REQUEST          = "Bad request"
	E_SERVER_ERROR         = "Internal server error"
)

var (
	ErrDuplicate = ErrorConstant{
		Response: Response{
			Success: false,
			Message: E_DUPLICATE,
		},
		Code: http.StatusConflict,
	}
	ErrNotFound = ErrorConstant{
		Response: Response{
			Success: false,
			Message: E_NOT_FOUND,
		},
		Code: http.StatusNotFound,
	}
	ErrUnprocessableEntity = ErrorConstant{
		Response: Response{
			Success: false,
			Message: E_UNPROCESSABLE_ENTITY,
		},
		Code: http.StatusUnprocessableEntity,
	}
	ErrUnauthorized = ErrorConstant{
		Response: Response{
			Success: false,
			Message: E_UNAUTHORIZED,
		},
		Code: http.StatusUnauthorized,
	}
	ErrBadRequest = ErrorConstant{
		Response: Response{
			Success: false,
			Message: E_BAD_REQUEST,
		},
		Code: http.StatusBadRequest,
	}
	ErrServerError = ErrorConstant{
		Response: Response{
			Success: false,
			Message: E_SERVER_ERROR,
		},
		Code: http.StatusInternalServerError,
	}
)

func BuildError(err ErrorConstant, msg error) error {
	err.ErrorMessage = msg
	return &err
}

func BuildCustomError(code int, message string) error {
	return &ErrorConstant{
		Response: Response{
			Success: false,
			Message: message,
		},
		Code:         code,
		ErrorMessage: nil,
	}
}

func SuccessResponse(c echo.Context, code int, msg string, data interface{}) error {
	response := Response{
		Success: true,
		Message: msg,
		Data:    data,
	}
	return c.JSON(code, response)
}

func ErrorResponse(c echo.Context, err error) error {
	body, e := ioutil.ReadAll(c.Request().Body)
	if e != nil {
		logrus.Warn("error read body, message : ", e.Error())
	}

	bHeader, e := json.Marshal(c.Request().Header)
	if e != nil {
		logrus.Warn("error read header, message : ", e.Error())
	}

	re, ok := err.(*ErrorConstant)
	if ok {
		log.InsertLogError(c, &log.LogError{
			Header:       string(bHeader),
			Body:         string(body),
			URL:          c.Request().URL.Path,
			HttpMethod:   c.Request().Method,
			ErrorMessage: re.Builder().ErrorMessage.Error(),
			Level:        "Error",
			AppName:      environment.Get("APP_NAME"),
			Version:      environment.Get("APP_VERSION"),
			Env:          environment.Get("ENV"),
			CreatedAt:    time.Now().Local().UTC(),
		})

		return c.JSON(re.Builder().Code, re.Builder().Response)
	} else {
		return c.JSON(re.Builder().Code, re.Builder().Response)
	}
}
