package repository

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/sholludev/sampoerna_notification/models"
	"gorm.io/gorm"
)

type Repository interface {
	GetByIDUser(c echo.Context, DB *gorm.DB, IDUser uint) ([]models.TNotifikasi, error)
	GetByID(c echo.Context, DB *gorm.DB, ID uint) (models.TNotifikasi, error)
	Create(c echo.Context, DB *gorm.DB, data *models.TNotifikasi) error
	UpdateToReadAll(c echo.Context, DB *gorm.DB, IDUser uint) ([]models.TNotifikasi, error)
	UpdateToReadSingle(c echo.Context, DB *gorm.DB, data *models.TNotifikasi) error
}
