package firebase

type MessageConstant struct {
	Action  string
	Message string
	Type    string
	Data    interface{}
}

type SingleNotification struct {
	FCMToken string
	Message  MessageConstant
}
