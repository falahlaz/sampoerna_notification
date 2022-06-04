package controller

import (
	"net/http"
	"strconv"

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

// GetByIDUser implements Controller
func (co *controllerImpl) GetByIDUser(c echo.Context) error {
	IDUserStr := c.Param("id_user")
	IDUser, err := strconv.ParseUint(IDUserStr, 10, 64)
	if err != nil {
		return response.ErrorResponse(c, response.BuildError(response.ErrBadRequest, err))
	}

	result, err := co.Service.GetByIDUser(c, uint(IDUser))
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, http.StatusOK, "success get by id user", result)
}

// UpdateToReadAll implements Controller
func (co *controllerImpl) UpdateToReadAll(c echo.Context) error {
	IDUserStr := c.Param("id_user")
	IDUser, err := strconv.ParseUint(IDUserStr, 10, 64)
	if err != nil {
		return response.ErrorResponse(c, response.BuildError(response.ErrBadRequest, err))
	}

	result, err := co.Service.UpdateToReadAll(c, uint(IDUser))
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, http.StatusOK, "success update to read all", result)
}

// UpdateToReadSingle implements Controller
func (co *controllerImpl) UpdateToReadSingle(c echo.Context) error {
	IDStr := c.Param("id")
	ID, err := strconv.ParseUint(IDStr, 10, 64)
	if err != nil {
		return response.ErrorResponse(c, response.BuildError(response.ErrBadRequest, err))
	}

	result, err := co.Service.UpdateToReadSingle(c, uint(ID))
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, http.StatusOK, "success update to read single", result)
}
