package notification

type SingleNotifRequestDTO struct {
	FCMToken string      `json:"fcm_token" validate:"required"`
	Type     string      `json:"type" validate:"required"`
	Message  string      `json:"message"`
	Data     interface{} `json:"data"`
}
