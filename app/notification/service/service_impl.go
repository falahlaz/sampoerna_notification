package service

import (
	"time"

	"github.com/labstack/echo/v4"
	"gitlab.com/sholludev/sampoerna_notification/app/notification"
	"gitlab.com/sholludev/sampoerna_notification/app/notification/repository"
	"gitlab.com/sholludev/sampoerna_notification/pkg/firebase"
	"gitlab.com/sholludev/sampoerna_notification/pkg/log"
	"gitlab.com/sholludev/sampoerna_notification/pkg/util/response"
	"gorm.io/gorm"
)

type serviceImpl struct {
	DB         *gorm.DB
	Repository repository.Repository
}

func NewService(DB *gorm.DB, Repository repository.Repository) Service {
	return &serviceImpl{
		DB:         DB,
		Repository: Repository,
	}
}

// SendSingleNotification implements Service
func (s *serviceImpl) SendSingleNotification(c echo.Context, request notification.SingleNotifRequestDTO) error {
	// validate
	if err := c.Validate(request); err != nil {
		return response.BuildError(response.ErrUnprocessableEntity, err)
	}

	// send notification
	firebaseModel := firebase.SingleNotification{
		FCMToken: request.FCMToken,
		Message: firebase.MessageConstant{
			Action:  "FLUTTER_NOTIFICATION_CLICK",
			Message: request.Message,
			Type:    request.Type,
			Data:    request.Data,
		},
	}

	_, err := firebaseModel.Send(c.Request().Context())
	if err != nil {
		return response.BuildError(response.ErrServerError, err)
	}

	return nil
}

// GetByIDUser implements Service
func (s *serviceImpl) GetByIDUser(c echo.Context, IDUser uint) ([]notification.NotifikasiResponseDTO, error) {
	var notifikasi []notification.NotifikasiResponseDTO

	// get notifikasi
	result, err := s.Repository.GetByIDUser(c, s.DB, IDUser)
	if err != nil {
		return notifikasi, response.BuildError(response.ErrServerError, err)
	}

	for _, v := range result {
		notifikasi = append(notifikasi, v.ToResponse())
	}

	// log
	log.InsertLogActivity(c.Request().Context(), &log.LogActivity{
		TableName: "t_notifikasi",
		Row:       notifikasi,
		Action:    "get by id user",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	return notifikasi, nil
}

// UpdateToReadAll implements Service
func (s *serviceImpl) UpdateToReadAll(c echo.Context, IDUser uint) ([]notification.NotifikasiResponseDTO, error) {
	var notifikasi []notification.NotifikasiResponseDTO

	// update to read all
	result, err := s.Repository.UpdateToReadAll(c, s.DB, IDUser)
	if err != nil {
		return notifikasi, response.BuildError(response.ErrServerError, err)
	}

	// binds to response
	for _, v := range result {
		notifikasi = append(notifikasi, v.ToResponse())
	}

	// log
	log.InsertLogActivity(c.Request().Context(), &log.LogActivity{
		TableName: "t_notifikasi",
		Row:       notifikasi,
		Action:    "update to read all",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	return notifikasi, nil
}

// UpdateToReadSingle implements Service
func (s *serviceImpl) UpdateToReadSingle(c echo.Context, ID uint) (notification.NotifikasiResponseDTO, error) {
	var notifikasi notification.NotifikasiResponseDTO

	// get by id
	result, err := s.Repository.GetByID(c, s.DB, ID)
	if err != nil {
		return notifikasi, response.BuildError(response.ErrNotFound, err)
	}

	// update to read single
	err = s.Repository.UpdateToReadSingle(c, s.DB, &result)
	if err != nil {
		return notifikasi, response.BuildError(response.ErrServerError, err)
	}

	// binds to response
	notifikasi = result.ToResponse()

	// log
	log.InsertLogActivity(c.Request().Context(), &log.LogActivity{
		TableName: "t_notifikasi",
		Row:       notifikasi,
		Action:    "update to read single",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	return notifikasi, nil
}
