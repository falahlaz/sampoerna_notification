package repository

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/sholludev/sampoerna_notification/models"
	"gorm.io/gorm"
)

type repositoryImpl struct {
}

func NewRepository() Repository {
	return &repositoryImpl{}
}

// Create implements Repository
func (*repositoryImpl) Create(c echo.Context, DB *gorm.DB, data *models.TNotifikasi) error {
	return DB.Create(data).Error
}

// GetByIDUser implements Repository
func (*repositoryImpl) GetByID(c echo.Context, DB *gorm.DB, ID uint) (models.TNotifikasi, error) {
	var data models.TNotifikasi
	err := DB.Preload("Kategori").Where("id = ?", ID).Where("is_active = true").First(&data).Error
	return data, err
}

// GetByID implements Repository
func (*repositoryImpl) GetByIDUser(c echo.Context, DB *gorm.DB, IDUser uint) ([]models.TNotifikasi, error) {
	var data []models.TNotifikasi
	err := DB.Preload("Kategori").Where("id_user = ?", IDUser).Where("is_active = true").Find(&data).Error
	return data, err
}

// UpdateToRead implements Repository
func (*repositoryImpl) UpdateToReadAll(c echo.Context, DB *gorm.DB, IDUser uint) ([]models.TNotifikasi, error) {
	var notifikasi []models.TNotifikasi
	err := DB.Where("id_user = ?", IDUser).Where("is_active = true AND is_read = false").Find(&notifikasi).Error
	if err != nil {
		return notifikasi, err
	}

	for _, v := range notifikasi {
		v.IsRead = true
		err = DB.Save(&v).Error
		if err != nil {
			return notifikasi, err
		}
	}

	return notifikasi, nil
}

// UpdateToReadSingle implements Repository
func (*repositoryImpl) UpdateToReadSingle(c echo.Context, DB *gorm.DB, data *models.TNotifikasi) error {
	data.IsRead = true
	return DB.Save(data).Error
}
